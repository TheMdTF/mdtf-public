CC=gcc

sample_matching_algorithm.o: sample_matching_algorithm.h sample_matching_algorithm.c
	@mkdir -p ../lib
	gcc -c sample_matching_algorithm.c -o ../lib/libsample_matching_algorithm.so
	
clean:
	@rm -rf ../lib
