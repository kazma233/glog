FROM golang:1.14

WORKDIR /app
COPY . /app

RUN go build -o app

RUN ["chmod", "755", "app"]

EXPOSE 9600
ENTRYPOINT ["./app"]