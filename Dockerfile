FROM alpine:3.6
WORKDIR /app
COPY ./server /app/
ENTRYPOINT ["./server"]

