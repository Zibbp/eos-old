FROM golang:1.18 AS build-stage-01

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o eos-api cmd/server/main.go

FROM alpine:latest AS production

COPY --from=build-stage-01 /app/eos-api .

CMD ["./eos-api"]
