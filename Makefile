all: tpp
	go build

tpp-files: k8s-tpp.tp
	lttng-gen-tp k8s-tpp.tp

tpp: tpp-files
	gcc -I. -c k8s-tpp.c -o k8s-tpp.o
	ar -rc k8s-tpp.a k8s-tpp.o

clean:
	rm -f *.o *.a *.c *.h
