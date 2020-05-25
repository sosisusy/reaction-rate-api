FROM golang:latest

MAINTAINER sosisusy <qwexodn@gmail.com>


WORKDIR /source
COPY . .

RUN go get github.com/canthefason/go-watcher \
    && go install github.com/canthefason/go-watcher/cmd/watcher

EXPOSE 80

CMD [ "watcher"]