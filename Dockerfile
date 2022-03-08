FROM golang:latest

WORKDIR /go/src/todo_list

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

EXPOSE 9010
RUN go mod verify
# RUN echo 'Now Running server.go'
CMD ["go", "run", "server.go"]

