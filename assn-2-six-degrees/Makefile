CMAKE ?= cmake
BUILD_DIR ?= build

.PHONY: all build clean configure install

all: build

configure:
	python ../gen_cmake.py
	python ../reformat.py
	python gen_start.py
	cmake --preset windows2

build: configure
	cmake --build --preset windows2
	python gen_lsp_tdm.py

run_a: all
	./build/six-degrees.exe

run_b: all
	./build/imdb-test.exe

package:
	fyne package

go:
	go fmt .
	go build .

clean:
	pwsh -c "rm -r $(BUILD_DIR)"
