from pathlib import Path
import json
from sodatools import read_path, write_path

CURRENT = Path(__file__).resolve().parent

commands = CURRENT.joinpath("build").joinpath("compile_commands.json")


def fix_command(c) -> str:
    print(c)


def fix_commands():
    c = read_path(commands)
    obj = json.loads(c)
    for item in obj:
        item["command"] = fix_command(item["command"])
    s = json.dumps(obj, indent=2)
    write_path(CURRENT.parent.joinpath("compile_commands.json"), s)


if __name__ == "__main__":
    fix_commands()
