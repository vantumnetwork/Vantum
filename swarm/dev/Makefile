.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/vantum/vantum/cmd/swarm
	go build -o bin/geth     github.com/vantum/vantum/cmd/geth
	go build -o bin/bootnode github.com/vantum/vantum/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/vantum/vantum/swarm/...
