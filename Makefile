build:
	go build -o ./build/glimr

execute: 
	./build/glimr

run: build execute