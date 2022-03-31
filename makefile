.PHONY: all dockerize-server build-server build-client clean

ifeq ($(OS),Windows_NT)
    MAKE = "C:/MinGW/msys/1.0/bin/make.exe"
endif

all: 
	go mod tidy
	"${MAKE}" build-server build-client

dockerize-server:
	"${MAKE}" -C server docker

build-server:
	"${MAKE}" -C server all

build-client:
	"${MAKE}" -C client all

clean:
	"${MAKE}" -C client clean
	"${MAKE}" -C server clean