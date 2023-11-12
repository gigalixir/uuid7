FROM golang:1.21.4

WORKDIR /app

VOLUME /app

ENV GOOS=linux GOARCH=amd64

CMD [ "go", "build", "-buildvcs=false", "-o", "uuid7" ]
