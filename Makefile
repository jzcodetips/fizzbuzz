.PHONY: all

all: run docker_build_run docker_rm

run:
	go mod vendor
	go run ./cmd/fizzbuzz/main.go -p 8888

docker_build_run:
	docker build -t fizzbuzz .
	docker run -d -p 8888:8888 --name fizzbuzz fizzbuzz

docker_rm:
	docker stop fizzbuzz
	docker rm fizzbuzz
