FROM golang:latest

MAINTAINER sosisusy <qwexodn@gmail.com>


COPY . /source
WORKDIR /source

RUN go mod download

EXPOSE 80

CMD [ "/bin/bash", "-c", "go run ./server.go" ]