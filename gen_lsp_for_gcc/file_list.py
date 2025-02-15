import glob
from pathlib import Path


def get_file_list(i: str):
    lis = list(glob.glob("**/*", recursive=True, root_dir=i))
    ret = []
    for p in lis:
        ret.append(Path(i).joinpath(p))
    # print(i)
    return ret
