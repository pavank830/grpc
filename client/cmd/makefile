.PHONY: all clean

all:
	go mod tidy
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=0 go build  -o userClient.exe
else
	CGO_ENABLED=0 go build  -o userClient
endif

clean:
ifeq ($(OS),Windows_NT)
	-powershell.exe -ExecutionPolicy Bypass -Command 'Remove-Item -Path ./userClient.exe -ErrorAction Continue'
else
	-rm -f ./userClient
endif