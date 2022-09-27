FROM golang:latest

RUN apt update && apt upgrade -y
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz && mv ./migrate /usr/bin/migrate
RUN curl -LO https://release.ariga.io/atlas/atlas-linux-amd64-latest && mv ./atlas-linux-amd64-latest /usr/bin/atlas && chmod 755 /usr/bin/atlas

WORKDIR /go/src/webapi
ENV PATH $PATH:/go/src/webapi/script

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest

CMD [ "./script/startup.sh" ]
