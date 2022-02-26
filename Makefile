build:
	go build -o server main.go

run: build
	./server

watch: 
	export PATH=$$PATH:$$GOPATH/bin
	reflex -s -r '\.go$$' make run
