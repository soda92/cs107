from pathlib import Path

CURRENT = Path(__file__).resolve().parent
f1 = CURRENT.joinpath("build/six_degree_main.cc")

if not f1.exists():
    f1.write_text(
        encoding="utf8",
        data="""

extern "C"{
#include "header.h"
}

int main(int argc, const char *argv[]){
return six_dg_main(argc, argv);
}

""",
    )

f2 = CURRENT.joinpath("build/imdb_test_main.cc")
if not f2.exists():
    f2.write_text(
        encoding="utf8",
        data="""

extern "C"{
#include "header.h"
}

int main(int argc, char** argv){
return imdb_test_main(argc, argv);
}

""",
    )
