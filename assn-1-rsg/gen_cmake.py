from pathlib import Path
import functools
import os
from sodatools import read_path, write_path, str_path

CURRENT = Path(__file__).resolve().parent

def which(name):
    path = os.environ["PATH"]
    for p in path.split(";"):
        exe = Path(p).joinpath(name)
        if exe.exists():
            return exe


@functools.cache
def get_gcc_path():
    out = which("gcc.exe")
    return out.resolve().parent.parent


if __name__ == "__main__":
    template = CURRENT.joinpath("CML.template")
    gcc_path = get_gcc_path()
    s = read_path(template)
    s = s.replace("{gcc}", str_path(gcc_path) + "/")
    write_path(CURRENT.joinpath("CMakeLists.txt"), s)
