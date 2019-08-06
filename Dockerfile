FROM golang:latest

LABEL maintainer="Agateram"

WORKDIR /solais
COPY /solais .

# RUN go clean -modcache
RUN go build -o solais .

EXPOSE 8080

CMD ["./solais"]

