CMAKE ?= cmake
BUILD_DIR ?= build

.PHONY: all build clean configure install

all: build

configure:
	python gen_cmake.py
	$(CMAKE) -B $(BUILD_DIR) -GNinja

build: configure
	cmake --build build
	# pwsh -c "cp build/compile_commands.json .."
	python gen_lsp.py
	python reformat.py

clean:
	pwsh -c "rm -r $(BUILD_DIR)"
