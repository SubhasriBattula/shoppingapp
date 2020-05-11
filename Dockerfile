FROM golang:latest
# create a working directory
WORKDIR /go/src/app
# copy
COPY ./supermarket/supermarket.go /supermarket/supermarket.go
# run main.go
CMD ["go", "run", "main.go"]
