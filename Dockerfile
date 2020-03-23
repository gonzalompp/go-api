# build image
FROM golang:1.14 as builder

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 go build -o api cmd/api/main.go

# final image
FROM alpine:3.11

WORKDIR /app

COPY --from=builder /app/api  .

CMD ["./api"]
