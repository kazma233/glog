FROM alpine

COPY ./app /app

RUN ["chmod", "755", "app"]

EXPOSE 9600

ENTRYPOINT ["./app"]
