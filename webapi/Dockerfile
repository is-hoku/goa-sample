FROM golang:latest

RUN apt update && apt upgrade -y
RUN curl -LO https://release.ariga.io/atlas/atlas-linux-amd64-latest && mv ./atlas-linux-amd64-latest /usr/local/bin/atlas && chmod 755 /usr/local/bin/atlas

WORKDIR /go/src/webapi
ENV PATH $PATH:/go/src/webapi/script

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
RUN go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
RUN go install goa.design/goa/v3/cmd/goa@v3
RUN go install github.com/google/wire/cmd/wire@latest
RUN go install github.com/rakyll/gotest@latest

CMD [ "./script/startup.sh" ]
