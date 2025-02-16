from sodatools import CD, str_path, get_glob_files
import subprocess
from pathlib import Path


CURRENT = Path(__file__).resolve().parent

if __name__ == "__main__":
    files = get_glob_files("**/*.h", "**/*.cc")
    for file in files:
        file = str_path(file)
        with CD(CURRENT.parent):
            subprocess.run(["clang-format.exe", file, "-i"])
