FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY go_server .

EXPOSE 8080
CMD ["./go_server"]
