FROM golang:latest
# RUN rm /bin/sh && ln -s /bin/bash /bin/sh

WORKDIR /go/src/todo_list
# RUN go mod init github.com/jiwanjeon/go-todolist
# RUN go get github.com/jinzhu/gorm
# RUN go get github.com/lib/pq
# RUN go get github.com/gorilla/mux
# RUN go get github.com/stretchr/testify
# RUN go get github.com/jinzhu/gorm/dialects/postgres
# source .env --> . setting.env & . setting & . ./setting.env & . ./setting & ["/bin/bash", "-c", "source ~/.setting.env"]

# SHELL ["/bin/bash", "-c"]
# RUN . .env
# RUN --env-file=setting
# RUN /bin/bash -c ". setting.env"

# Source 명령어가 실행 되지 않아 아래와 같이 ENV command으로 실행
# RUN source .env

ENV DIALECT "postgres"
ENV HOST "host.docker.internal"
ENV DBPORT "5432"
ENV USER "postgres"
ENV NAME "todo_list_second"
ENV PASSWORD "wjswldhks"

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

WORKDIR /go/src/todo_list/cmd/main
EXPOSE 9010
RUN go mod verify
RUN go build main.go
CMD ["./main", "-migrate"]
# RUN ./main -migrate