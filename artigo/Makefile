.PHONY: dist test deploy_rpi dist_arm dist_native


# by default, my raspberry pi instance is called rpi
# replace with yours
REMOTE_HOST ?= rpi

test:
	go test ./...

dist_native:
	go build -o ./dist/deviced ./cmd/deviced
	go build -o ./dist/merkled ./cmd/merkled
	go build -o ./dist/acunitd ./cmd/acunitd
	go build -o ./dist/extractor ./cmd/extractor

dist_arm:
	GOARCH=arm GOARM=6 go build -o ./dist/rpi/deviced.arm ./cmd/deviced
	GOARCH=arm GOARM=6 go build -o ./dist/rpi/merkled.arm ./cmd/merkled
	GOARCH=arm GOARM=6 go build -o ./dist/rpi/acunitd.arm ./cmd/acunitd

dist: test dist_arm dist_native
	mkdir -p dist/ && mkdir -p dist/rpi

	cp scripts/* ./dist/rpi/
	cp scripts/* ./dist


deploy_rpi: dist
	scp -r ./dist/rpi $(REMOTE_HOST):