.PHONY: all build container


build: 
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bgp-crd controller/bgp-controller.go

container: build
	docker build -t bgp-crd:latest .
