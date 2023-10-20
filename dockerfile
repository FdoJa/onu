FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin

RUN apt-get update && apt-get install -y protobuf-compiler

RUN go get google.golang.org/grpc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


RUN protoc --go_out=. --go-grpc_out=. proto/*.proto

EXPOSE 80
ENTRYPOINT ["/app/bin"]
