all:
	g++ -std=c++11 -I./cppthulac -DLOGGING_LEVEL=LL_WARNING -O3 -Wall -fprofile-arcs -ftest-coverage thulac.cc -shared -fPIC -o thulac.so

test:
	go test -count=1 -v ./...
	gcov --branch-probabilities --branch-counts thulac.cc -o .

clean:
	rm -f *.o
	rm -f *.so
	rm -f *.gcda
	rm -f *.gcno
	rm -f *.gcov
	rm -rf bw-output
