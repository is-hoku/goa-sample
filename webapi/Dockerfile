FROM golang:latest

RUN apt update && apt upgrade -y

RUN mkdir /go/src/webapi
WORKDIR /go/src/webapi

ADD . /go/src/webapi

CMD [ "./scripts/startup.sh" ]
