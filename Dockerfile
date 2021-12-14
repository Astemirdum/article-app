FROM golang:1.16

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

EXPOSE 8080

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o article-app ./cmd/main.go

CMD ["./article-app"]