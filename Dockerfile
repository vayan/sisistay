FROM golang:1.12

WORKDIR /go/src/github.com/vayan/sisistay

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon", "github.com/golang/dep/cmd/dep"]

RUN ["dep", "ensure"]

ENTRYPOINT CompileDaemon -build="go build ./src/cmd/server/" -command="./server"
