FROM golang:latest

WORKDIR /go/src/todo_list

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN /bin/bash -c ". setting.env"

WORKDIR /go/src/todo_list/cmd/main
EXPOSE 9010
RUN go mod verify
RUN go build main.go
CMD ["./main", "-migrate"]