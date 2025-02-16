cmake_minimum_required(VERSION 3.22)

project(assn2)

set(CMAKE_CXX_COMPILER C:/TDM-GCC-64/bin/g++.exe)
set(CMAKE_C_COMPILER C:/TDM-GCC-64/bin/gcc.exe)

add_compile_options(-g)
set(CMAKE_EXPORT_COMPILE_COMMANDS true)

include_directories(.)

add_executable(six-degrees
"six-degrees.cc"
"path.cc"
"imdb.cc"
"mmap.cc"
)

add_executable(imdb-test
"imdb-test.cc"
"imdb.cc"
"mmap.cc"
)
