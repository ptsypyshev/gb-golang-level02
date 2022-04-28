.SILENT: start

# For Windows only variables, because of some troubles to start compilation
# Should be set by user
GOOSVAR := windows
GOARCHVAR := amd64

start:
	echo "You have next options:"
	echo "build_win		(For Windows - To build application for Windows, Linux and FreeBSD)"
	echo "build_lin		(For *nix - To build application for Windows, Linux and FreeBSD)"
	echo "test		(To test application)"
	echo "run		(To run application)"

build_win:
	@echo "Compiling for some OS and Platform"
	go env -w GOARCH=386
	go env -w GOOS=windows
	go build -o bin/fibonacci-windows-386.exe main.go

	go env -w GOARCH=amd64
	go env -w GOOS=linux
	go build -o bin/fibonacci-linux-x86_64 main.go

	go env -w GOARCH=386
	go env -w GOOS=freebsd
	go build -o bin/fibonacci-freebsd-386 main.go

	go env -w GOARCH=$(GOARCHVAR)
	go env -w GOOS=$(GOOSVAR)

build_lin:
	@echo "Compiling for some OS and Platform"
	GOOS=windows GOARCH=386 go build -o bin/fibonacci-windows-386.exe main.go
	GOOS=linux GOARCH=amd64 go build -o bin/fibonacci-linux-x86_64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/fibonacci-freebsd-386 main.go

run:
	go run main.go

test:
	echo "Run tests for fibonacci package"
	go test ./fibonacci -v
