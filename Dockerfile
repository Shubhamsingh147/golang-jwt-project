FROM golang:latest

COPY . /usr/local/go/src/golang-jwt-artica
WORKDIR /usr/local/go/src/golang-jwt-artica

RUN go mod tidy
RUN go mod vendor
RUN go build -o main main.go
EXPOSE 8000
CMD [ "/app/main" ]