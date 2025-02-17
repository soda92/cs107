CMAKE ?= cmake
BUILD_DIR ?= build
CXX := C:/TDM-GCC-64/bin/g++.exe

.PHONY: all build clean configure install

all: main

main: main.cc
	$(CXX) $@.cc -o $@.exe
	./$@.exe

string: string.cc
	$(CXX) $@.cc -o $@.exe
	./$@.exe

str2: str2.cc
	$(CXX) $@.cc -o $@.exe
	./$@.exe

fs: fs.cc
	$(CXX) $@.cc -o $@.exe
	./$@.exe
