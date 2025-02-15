from pathlib import Path
import json
from sodatools import write_path, str_path

CURRENT = Path(__file__).resolve().parent

tdm_dir = r"c:\TDM-GCC-64"
db = Path(tdm_dir).joinpath("compile_commands.json")

dirs = [
    r"{gcc}lib\gcc\x86_64-w64-mingw32\10.3.0\include\c++",
    r"{gcc}x86_64-w64-mingw32\include",
    r"{gcc}lib\gcc\x86_64-w64-mingw32\10.3.0\include\c++\x86_64-w64-mingw32",
]

includes_list = []

for i in dirs:
    i = i.replace("\\", "/")
    i = i.replace("{gcc}", tdm_dir + "/")
    includes_list.append("-isystem " + i)
includes = " ".join(includes_list)

objs = []

dummy_cc = str_path(Path(tdm_dir).joinpath("demo.cc"))
write_path(Path(dummy_cc), "")

for i in dirs:
    i = i.replace("\\", "/")
    i = i.replace("{gcc}", tdm_dir + "/")
    files = Path(i).glob("*")

    for f in files:
        if f.name == "iostream":
            print(f)
        obj = {
            "directory": tdm_dir,
            "command": f"C:/TDM-GCC-64/bin/g++.exe {includes} {dummy_cc}",
            "file": str_path(f),
            "output": str_path(Path(tdm_dir).joinpath("a.obj")),
        }
        objs.append(obj)


s = json.dumps(objs, indent=2)
write_path(db, s)
