.PHONY: all docker clean

all:
	go mod tidy
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=0 go build  -o userServer.exe
else
	CGO_ENABLED=0 go build  -o userServer
endif

docker:
	go mod tidy
	GOOS=linux CGO_ENABLED=0 go build -o userServer
	docker build -f Dockerfile . -t pavank830/user:v1

clean:
ifeq ($(OS),Windows_NT)
	-powershell.exe -ExecutionPolicy Bypass -Command 'Remove-Item -Path ./userServer.exe -ErrorAction Continue'
else
	-rm -f ./userServer
endif