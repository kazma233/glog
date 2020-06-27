FROM golang:1.14.4-alpine3.12

COPY ./app /app

RUN ["chmod", "755", "app"]

EXPOSE 9600

ENTRYPOINT ["./app"]
