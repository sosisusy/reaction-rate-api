FROM golang:latest

MAINTAINER sosisusy <qwexodn@gmail.com>


COPY ./source /source
WORKDIR /source



RUN go mod download \
    && go build -o main.exe

EXPOSE 80

CMD [ "./main.exe" ]