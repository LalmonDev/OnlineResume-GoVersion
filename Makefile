build_main:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o OnlineResume_Version_Mac_M1
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o OnlineResume_Version_Linux

clean:
	@go clean & rm -rf OnlineResume_Version_*