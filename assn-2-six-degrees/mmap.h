// Simonr's profile photo
// Simonr
// unread,
// Aug 18, 2010, 7:18:44 PM
// to C++ RTMP Server
// Hi,

// If anybodies interested... I have implemented MMAP for the windows
// code:

// Add the following to Win32platform.h:
#ifndef __mmap_h__
#define __mmap_h__
#include <cstddef>

#include <_mingw_off_t.h>
#define MAP_NOCACHE (0)
#define MAP_NOEXTEND 0x0100 /* for MAP_FILE, don't change file size \
                             */
#define MAP_FAILED ((void *)-1)
// #define O_RDONLY 0x0000
// #define O_RDWR 0x0002

#define PROT_READ 0x1 /* Page can be read. */
#define PROT_WRITE 0x2 /* Page can be written. */
#define PROT_EXEC 0x4 /* Page can be executed. */
#define PROT_NONE 0x0 /* Page can not be accessed. \
                       */

#define MAP_SHARED 0x01 /* Share changes. */
#define MAP_PRIVATE 0x02 /* Changes are private. */
#define MAP_FIXED 0x10 /* Interpret addr exactly. \
                        */

#define DLLEXP

DLLEXP long getpagesize(void);
DLLEXP long getregionsize(void);
DLLEXP void *mmap(void *start, size_t len, int prot, int flags, int fd, off_t offset);
DLLEXP int munmap(void *start, size_t len);

// Add the following to Win32platform.cpp
#endif