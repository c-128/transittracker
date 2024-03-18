FROM docker.io/golang:1.22.1 AS builder

WORKDIR /app

COPY . .
RUN go build -o main main.go

FROM docker.io/debian:12.5
COPY --from=builder /app/main /main

RUN apt-get update
RUN apt-get install ca-certificates -y

EXPOSE 3000

CMD ["/main"]
