FROM golang:latest

LABEL maintainer="Agateram"

WORKDIR /solais
COPY /solais .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

