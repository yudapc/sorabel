sorabel-osx: app.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

sorabel-linux: app.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@