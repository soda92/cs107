from pathlib import Path
import os
from sodatools import read_path, write_path, str_path, get_glob_files
import argparse
import sys
import logging


CURRENT = Path(__file__).resolve().parent
sys.path.insert(0, str_path(CURRENT))


if __name__ == "__main__":
    for template in get_glob_files("**/CMakePresets.json"):
        target = template.parent.joinpath("CMakeUserPresets.json")
        s = """
{
  "version": 10,
  "cmakeMinimumRequired": {
    "major": 3,
    "minor": 23,
    "patch": 0
  },
  "configurePresets": [
    {
      "name": "windows2",
      "inherits": "windows-only",
      "displayName": "Windows-only configuration",
      "cacheVariables": {
        "CMAKE_CXX_COMPILER": {
          "type": "STRING",
          "value": "C:/TDM-GCC-64/bin/g++.exe"
        },
        "CMAKE_C_COMPILER": "C:/TDM-GCC-64/bin/gcc.exe",
        "SECOND_CACHE_VARIABLE": "ON"
      }
    }
  ],
  "buildPresets": [
    {
      "name": "windows2",
      "configurePreset": "windows2"
    }
  ]
}
        """
        write_path(target, s)
