#!/bin/bash

cd ../
rm build/ -rf && cmake -G "Unix Makefiles" -B build . && cd build && make && ./ThulacgoTest
