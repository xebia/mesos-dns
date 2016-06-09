include Makefile.mk

USERNAME=xebia
NAME=$(shell basename $(PWD))

pre-build: mesos-dns

post-build:
	rm -rf empty

mesos-dns: main.go 
		rm -f mesos-dns
		mkdir empty
	        docker run --rm -v $$(pwd):/src -v /var/run/docker.sock:/var/run/docker.sock centurylink/golang-builder


