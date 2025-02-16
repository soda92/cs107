CMAKE ?= cmake
BUILD_DIR ?= build

.PHONY: all build clean configure install

all: build

configure:
	python ../gen_cmake.py
	python ../reformat.py
	cmake --preset windows2

build: configure
	cmake --build --preset windows2
	CC=C:/TDM-GCC-64/bin/gcc.exe CXX=C:/TDM-GCC-64/bin/g++.exe go build .

run: all
	./assn-1-rsg.exe

package:
	fyne package

clean:
	pwsh -c "rm -r $(BUILD_DIR)"
