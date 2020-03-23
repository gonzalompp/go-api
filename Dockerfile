# build image
FROM golang:1.14              
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN go build -o api cmd/api/main.go
CMD ["/app/api"]
