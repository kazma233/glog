FROM ubuntu:18.04

COPY ./app /app

RUN ["chmod", "755", "/app"] \
    echo "Asia/Shanghai" > /etc/timezone

EXPOSE 9600

ENTRYPOINT ["/app"]
