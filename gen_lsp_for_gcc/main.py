import sys
from pathlib import Path
import json
from sodatools import write_path, str_path

CURRENT = Path(__file__).resolve().parent
sys.path.insert(0, str_path(CURRENT))

tdm_dir = r"c:/TDM-GCC-64"
tdm_dir_v = r"c:/TDM-GCC-64/virtual"
db = Path(tdm_dir_v).joinpath("compile_commands.json")

dirs = [
    r"{gcc}lib\gcc\x86_64-w64-mingw32\10.3.0\include\c++",
    r"{gcc}x86_64-w64-mingw32\include",
    r"{gcc}lib\gcc\x86_64-w64-mingw32\10.3.0\include\c++\x86_64-w64-mingw32",
]

includes_list = []

for i in dirs:
    i = i.replace("\\", "/")
    source = i.replace("{gcc}", tdm_dir + "/")
    virtual = i.replace("{gcc}", tdm_dir_v + "/")
    includes_list.append("-isystem " + virtual)
includes = " ".join(includes_list)

objs = []

dummy_cc = str_path(Path(tdm_dir_v).joinpath("demo.cc"))
write_path(Path(dummy_cc), "")

for i in dirs:
    i = i.replace("\\", "/")
    i = i.replace("{gcc}", tdm_dir + "/")
    from file_list import get_file_list
    files = get_file_list(i)

    for f in files:
        # if f.name == "iostream":
        #     print(f)
        from fix_mingw import write_virtual
        write_virtual(f)
        obj = {
            "directory": tdm_dir_v,
            "command": f"C:/TDM-GCC-64/bin/g++.exe {includes} {dummy_cc}",
            "file": dummy_cc,
            "output": str_path(Path(tdm_dir_v).joinpath("a.obj")),
        }
        objs.append(obj)


s = json.dumps(objs, indent=2)
write_path(db, s)
