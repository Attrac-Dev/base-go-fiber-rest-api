build:
	go build -o server main.go

run: build
	./server

watch: 
	nodemon --exec "go run" main.go