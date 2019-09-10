FROM golang:latest
RUN go build .
ENTRYPOINT ["./main"]