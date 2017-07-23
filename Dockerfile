FROM alpine:3.6
WORKDIR /app
COPY app /app/
ENTRYPOINT ["./app"]

