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

run_a: all
	./build/six-degrees.exe

run_b: all
	./build/imdb-test.exe

package:
	fyne package

clean:
	pwsh -c "rm -r $(BUILD_DIR)"
