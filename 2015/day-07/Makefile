.DEFAULT_GOAL := all
NAME = day-07-some-assembly-required

all: time-node time-lisp
time-c: build
	time ./a.out

build:
	clang ./$(NAME).c

run-c: build
	./a.out

run-node:
	node ./$(NAME).js

time-node:
	time node ./$(NAME).js

run-lisp:
	sbcl --script ./$(NAME).lisp

time-lisp:
	time sbcl --script ./$(NAME).lisp

clean:
	rm -f ./*.out
	rm -f ./*test.txt
	rm -f ./*.fasl

leaks: build
	leaks --atExit -- ./a.out
