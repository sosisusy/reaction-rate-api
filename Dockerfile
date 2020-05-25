FROM golang:latest

MAINTAINER sosisusy <qwexodn@gmail.com>


COPY ./source /source
WORKDIR /source



RUN go mod download
# RUN go mod download \
#     && go build -o main.exe

EXPOSE 80

VOLUME [ "/data" ]

# CMD [ "./main.exe" ]
CMD [ "/bin/bash", "-c", "go run ./server.go" ]