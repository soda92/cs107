CMAKE ?= cmake
BUILD_DIR ?= build

.PHONY: all build clean configure install

all: build

configure:
	$(CMAKE) -B $(BUILD_DIR) -GNinja

build: configure
	cmake --build build
	pwsh -c "cp build/compile_commands.json .."

clean:
	pwsh -c "rm -r $(BUILD_DIR)"
