cmake_minimum_required(VERSION 3.22)

project(assn2)

add_compile_options(-g)
set(CMAKE_EXPORT_COMPILE_COMMANDS true)

include_directories(.)

add_executable(six-degrees
"six-degrees.cc"
"path.cc"
"imdb.cc"
"mmap.cc"
"build/six_degree_main.cc"
)

add_executable(imdb-test
"imdb-test.cc"
"imdb.cc"
"mmap.cc"
"build/imdb_test_main.cc"
)
