.PHONY: dist test deploy_rpi

dist: test
	mkdir -p dist/ && mkdir -p dist/rpi

	go build -o ./dist/deviced ./cmd/deviced
	go build -o ./dist/merkled ./cmd/merkled
	GOARCH=arm GOARM=6 go build -o ./dist/rpi/deviced.arm ./cmd/deviced
	GOARCH=arm GOARM=6 go build -o ./dist/rpi/merkled.arm ./cmd/merkled


deploy_rpi: dist
	scp -r ./dist/rpi $(REMOTE_HOST):

test:
	go test ./...