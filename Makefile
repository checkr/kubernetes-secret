.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build -o bin/linux/kubernetes-secret
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/kubernetes-secret
	GOOS=windows GOARCH=amd64 go build -o bin/windows/kubernetes-secret.exe
	chmod +x bin/linux/kubernetes-secret bin/darwin/kubernetes-secret bin/windows/kubernetes-secret.exe

release:
	tar -zcvf kubernetes-secret.tar.gz bin/
