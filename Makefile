all: tpp
	go build

k8s-tpp.a: k8s-tpp.c k8s-tp.h
	gcc -I. -c k8s-tpp.c -o k8s-tpp.o
	ar -rc k8s-tpp.a k8s-tpp.o

tpp: k8s-tpp.a

clean:
	rm -f *.o *.a *.out

benchmark: tpp
	go test -bench=.

profile: tpp
	go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out