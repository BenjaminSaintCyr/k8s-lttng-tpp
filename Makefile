all: tpp
	go build

tpp:
	gcc -I. -c k8s-tpp.c -o k8s-tpp.o
	ar -rc k8s-tpp.a k8s-tpp.o

clean:
	rm -f *.o *.a
