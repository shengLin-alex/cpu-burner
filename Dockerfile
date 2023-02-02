FROM golang:1.17.5-alpine3.15 as builder

WORKDIR /app

COPY main.go main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
COPY --from=builder /app/main /main
ENV GIN_MODE=release
ENTRYPOINT ["/main"]
