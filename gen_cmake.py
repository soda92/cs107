from pathlib import Path
import os
from sodatools import read_path, write_path, str_path, get_glob_files
import argparse
import sys
import logging


CURRENT = Path(__file__).resolve().parent
sys.path.insert(0, str_path(CURRENT))


def get_cmake_templates() -> list[Path]:
    return get_glob_files("**/CML.template")


def which(name) -> Path:
    "find executable by name from PATH"
    path = os.environ["PATH"]
    for p in path.split(";"):
        exe = Path(p).joinpath(name)
        if exe.exists():
            return exe
    logging.fatal(f"{name} not found in PATH")
    exit(-1)


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("--gcc", action="store_true", help="use gcc in PATH")
    parser.add_argument("--tdm", action="store_true", help="use tdm gcc")
    args = parser.parse_args()
    for template in get_cmake_templates():
        target = template.parent.joinpath("CMakeLists.txt")
        s = read_path(template)
        if args.tdm:
            s = s.replace("g++", "C:/TDM-GCC-64/bin/g++.exe")
            s = s.replace("gcc", "C:/TDM-GCC-64/bin/gcc.exe")
        else:
            s = s.replace("g++", which("g++"))
            s = s.replace("gcc", which("gcc"))
        write_path(target, s)
