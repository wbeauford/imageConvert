.DEFAULT_GOAL := default
default:
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo
	@docker build -t imageconvert .