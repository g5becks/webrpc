all:
	@echo "please read Makefile source or README to see available commands"

tools:
	GO111MODULE=off go get -u github.com/goware/webify

generate: generate-server generate-client

generate-server:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=go -pkg=main -server -out=./server/hello_api.gen.go

generate-client:
	../../bin/webrpc-gen -schema=hello-api.ridl -target=ts -client -out=./webapp/src/client.gen.ts

bootstrap:
	rm -rf webapp/node_modules
	cd webapp && yarn

run-server:
	go run ./server

run-client:
	cd webapp && npm start
