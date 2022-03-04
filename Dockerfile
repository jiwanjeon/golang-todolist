FROM golang:latest

# WORKDIR /go/src/todo_list

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

WORKDIR /go/src/todo_list/cmd/main
EXPOSE 9010
RUN go mod verify
RUN go build main.go
CMD ["./main", "-migrate"]

# CMD ["go", "run", "main.go", "-migrate"]
