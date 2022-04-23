FROM golang:1.18-alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
RUN apk update --no-cache && apk add --no-cache tzdata
WORKDIR /build
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/article ./main.go

FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app
EXPOSE 8080
COPY --from=builder /app/article /app/article
CMD ["./hello"]
