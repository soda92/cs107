// clang-format off
//go:build ignore
// clang-format on
#include "mmap.h"

#include <Windows.h>
#include <cerrno>
#include <io.h>
#include <map>

using namespace std;

static map<void *, void *> _mmap_map; // map addressses to cater for granularity
long getpagesize(void)
{
  static long g_pagesize = 0;
  if (!g_pagesize) {
    SYSTEM_INFO system_info;
    GetSystemInfo(&system_info);
    g_pagesize = system_info.dwPageSize;
  }
  return g_pagesize;
}
long getregionsize(void)
{
  static long g_regionsize = 0;
  if (!g_regionsize) {
    SYSTEM_INFO system_info;
    GetSystemInfo(&system_info);
    g_regionsize = system_info.dwAllocationGranularity;
  }
  return g_regionsize;
}

void *mmap(void *start, size_t len, int prot, int flags, int fd, off_t offset)
{

  /*
  This is a minimal implementation supporting:
  PROT_NONE, PROT_READ, PROT_WRITE, PROT_EXEC
  It handles the following mapping flags:
  MAP_NOCACHE, MAP_FIXED, MAP_SHARED, MAP_PRIVATE
  */

  void *ret = MAP_FAILED;
  HANDLE hmap = NULL;
  long wprot = 0, wflags = 0;
  HANDLE hfile = (HANDLE)_get_osfhandle(fd);
  long file_len = _filelength(fd);

  /* map *NIX protections and flags to their WIN32 equivalents */
  if ((prot & PROT_READ) && (~prot & PROT_WRITE)) {
    /* read only, maybe exec */
    wprot = (prot & PROT_EXEC) ? (PAGE_EXECUTE_READ) : (PAGE_READONLY);
    wflags = (prot & PROT_EXEC) ? (FILE_MAP_EXECUTE) : (FILE_MAP_READ);
    if (flags & MAP_NOCACHE)
      wprot = wprot | SEC_NOCACHE;
  } else if (prot & PROT_WRITE) {
    /* read/write, maybe exec */
    if ((flags & MAP_SHARED) && (~flags & MAP_PRIVATE)) {
      /* changes are committed to the file */
      wprot = (prot & PROT_EXEC) ? (PAGE_EXECUTE_READWRITE) : (PAGE_READWRITE);
      wflags = (prot & PROT_EXEC) ? (FILE_MAP_EXECUTE) : (FILE_MAP_WRITE);
    } else if ((flags & MAP_PRIVATE) && (~flags & MAP_SHARED)) {
      /* does not affect the original file */
      wprot = PAGE_WRITECOPY;
      wflags = FILE_MAP_COPY;
    } else {
      /* MAP_PRIVATE + MAP_SHARED is not allowed, abort */
      errno = EINVAL;
      return MAP_FAILED;
    }
  }

  int granularityOffset = (int)(offset % getregionsize());
  offset -= granularityOffset; // align offset
  len = (offset + len + granularityOffset > file_len) ? file_len - offset - granularityOffset : len; // check to see if offset + len in bounds

  /* create the windows map object */
  hmap = CreateFileMapping(hfile, NULL, wprot, 0, offset + len + granularityOffset, NULL);

  if (hmap == NULL) {
    /* the fd was checked before, so it must have bad access rights */
    errno = EPERM;
    return MAP_FAILED;
  }

  ret = MapViewOfFileEx(hmap, wflags, 0, offset, len + granularityOffset,
      (flags & MAP_FIXED) ? (start) : (NULL));

  int err = GetLastError();

  /* Drop the map, it will not be deleted until last 'view' is
  closed */
  CloseHandle(hmap);

  if (ret == NULL || err != 0) {
    /* if MAP_FIXED was set, the address was probably wrong */
    errno = (flags & MAP_FIXED) ? (EINVAL) : (ENOMEM);
    return MAP_FAILED;
  }

  if (granularityOffset) // map the return base address to viewoffile base address
    _mmap_map[(void *)((char *)ret + granularityOffset)] = ret;

  return ((void *)((char *)ret + granularityOffset));
}

template <typename K, typename V>
bool MAP_HAS1(const std::map<K, V>& map, const K& key)
{
  return map.find(key) != map.end();
}

int munmap(void *start, size_t len)
{
  if (start == NULL) {
    errno = EINVAL;
    return -1;
  }

  if (MAP_HAS1(_mmap_map, start)) // does it have a adjusted base for granularity?
    return UnmapViewOfFile(_mmap_map[start]) ? 0 : -1;
  else
    return UnmapViewOfFile(start) ? 0 : -1;
}
