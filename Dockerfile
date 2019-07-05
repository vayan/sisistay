FROM golang:1.12

WORKDIR /go/src/github.com/vayan/sisistay

COPY . .

ENV PORT=8080

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN ["go", "get", "github.com/golang/dep/cmd/dep"]

RUN ["dep", "ensure"]
RUN ["go", "build", "./src/cmd/server/"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./src/cmd/server/" -command="./server"
