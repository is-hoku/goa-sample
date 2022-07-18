FROM golang:latest

RUN apt update && apt upgrade -y
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate /usr/bin/migrate

RUN mkdir /go/src/webapi
WORKDIR /go/src/webapi

ADD . /go/src/webapi
ENV PATH $PATH:/go/src/webapi/script

RUN go install github.com/cosmtrek/air@latest

CMD [ "./script/startup.sh" ]
