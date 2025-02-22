.PHONY: gen_cmake reformat pre

gen_cmake:
	py gen_cmake.py --tdm

reformat:
	py reformat.py

pre: gen_cmake reformat


all:
	cd assn-1-rsg; make

run1:
	cd assn-1-rsg; make run

assn2:
	cd assn-2-six-degrees; make

assn2g:
	cd assn-2-six-degrees; make go

2m: pre
	cd assn-2-six-degrees; make run_a

2t: pre
	cd assn-2-six-degrees; make run_b
