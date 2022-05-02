build: 
	go build -o ./glimr

server: build
	./glimr server

client: build
	./glimr client