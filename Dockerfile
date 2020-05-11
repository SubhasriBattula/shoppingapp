FROM golang:latest
# create a working directory
WORKDIR /go/src/app
# add source code
ADD main.go main.go
# run main.go
CMD ["go", "run", "main.go"]
