from sodatools import read_path, write_path
from pathlib import Path

# file = r"C:\TDM-GCC-64\x86_64-w64-mingw32\include\stdio.h"

# c = read_path(Path(file))

# c = "#define __MINGW_NOTHROW // soda modify\n" + c

# write_path(Path(file), c)
cnt = 0


def fix_content(c):
    c = c.replace("__MINGW_ATTRIB_NORETURN", "")
    c = c.replace("__MINGW_NOTHROW", "")
    return c


def write_virtual(f):
    if f.is_dir():
        return
    global cnt
    base = Path(r"c:\TDM-GCC-64")
    virtual = Path(r"c:\TDM-GCC-64\virtual")
    rel = Path(f).relative_to(base)
    vfile = virtual.joinpath(rel)

    vfile.parent.mkdir(parents=True, exist_ok=True)

    try:
        c = fix_content(read_path(f))
        write_path(vfile, c)
    except UnicodeDecodeError:
        return

