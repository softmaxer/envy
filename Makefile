BINARY=envy

build:
	@GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY}-darwin-amd64 .
	@GOARCH=arm64 GOOS=darwin go build -o ./target/${BINARY}-darwin-arm64 .
	@GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY}-linux-amd64 .
	@GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY}-windows-amd64 .

run: build
	@./target/${BINARY}-linux-amd64

clean:
	rm -r ./target/*

