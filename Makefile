V = "1"


help:
	@echo "Argument :"
	@echo "docker - docker build and run the image"
	@echo "build - just build and start"
	@echo "V is the variable for the image name"

build :
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

docker :
	@sudo docker build -t runserver:$(V) .
	@sudo docker images
