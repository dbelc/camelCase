FROM golang:1.8-alpine

WORKDIR /go/src/github.com/dbelc/camelCase
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD ["camelCase"]