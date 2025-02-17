from pathlib import Path

CURRENT = Path(__file__).resolve().parent

CURRENT.joinpath("build/six_degree_main.cc").write_text(encoding='utf8',
data = """

extern "C"{
#include "header.h"
}

int main(int argc, const char *argv[]){
return six_dg_main(argc, argv);
}

""")

CURRENT.joinpath("build/imdb_test_main.cc").write_text(encoding='utf8',
data = """

extern "C"{
#include "header.h"
}

int main(int argc, char** argv){
return imdb_test_main(argc, argv);
}

""")