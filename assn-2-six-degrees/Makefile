CMAKE ?= cmake
BUILD_DIR ?= build

.PHONY: all build clean configure install

all: build

configure:
	$(CMAKE) -B $(BUILD_DIR) -GNinja

build: configure
	cmake --build build

run_a: all
	./build/six-degrees.exe

run_b: all
	./build/imdb-test.exe

package:
	fyne package

clean:
	pwsh -c "rm -r $(BUILD_DIR)"
