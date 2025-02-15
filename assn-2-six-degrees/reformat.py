from sodatools import CD
import glob
from pathlib import Path

files = list(glob.glob("*.h"))
files.extend(list(glob.glob("*.cc")))

CURRENT = Path(__file__).resolve().parent
for file in files:
    with CD(CURRENT.parent):
        import subprocess

        subprocess.run(["C:/src/llvm-project-install/bin/clang-format", CURRENT.name + "/" + file, "-i"])
