cmake_minimum_required(VERSION 3.22)

project(assn1)

set(CMAKE_EXPORT_COMPILE_COMMANDS true)

set(SOURCES
    definition.cc
    production.cc
    random.cc
    main.cc
)

add_library(rsg ${SOURCES})
target_include_directories(rsg PUBLIC .)
