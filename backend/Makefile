build:
	go mod download && go build -o ./.bin/interferometer.exe ./cmd/app/main.go

run: build
	./.bin/interferometer.exe

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./.bin/interferometer.exe cmd/app/main.go

clean:
	rm -rf .bin

.DEFAULT_GOAL := run
