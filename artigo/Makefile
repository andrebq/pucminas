.PHONY: dist test deploy_rpi

dist: test
	mkdir -p dist/ && mkdir -p dist/rpi

	go build -o ./dist/deviced ./cmd/deviced
	GOARCH=arm GOARM=6 go build -o ./dist/rpi/deviced.arm ./cmd/deviced

deploy_rpi: dist
	scp -r ./dist/rpi $(REMOTE_HOST):

test:
	go test ./...