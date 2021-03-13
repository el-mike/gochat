# build stage
FROM golang:1.14-alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="Michal Huras <michal.huras5@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# final stage
FROM alpine

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8080
RUN chmod a+x ./main

CMD ["./main"]
