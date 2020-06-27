FROM ubuntu:18.04

COPY ./app /app

RUN ["chmod", "755", "/app"]

EXPOSE 9600

ENTRYPOINT ["/app"]
