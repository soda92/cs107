CMAKE ?= cmake
BUILD_DIR ?= build

.PHONY: all build clean configure install

all: main

main: main.cc
	g++ $@.cc -o $@.exe
	./$@.exe

string: string.cc
	g++ $@.cc -o $@.exe
	./$@.exe

str2: str2.cc
	g++ $@.cc -o $@.exe
	./$@.exe
