FROM golang:1 AS build-stage-01

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o eos-api cmd/server/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o eos-worker cmd/worker/worker.go

FROM alpine:latest AS production

COPY --from=build-stage-01 /app/eos-api .

COPY --from=build-stage-01 /app/eos-worker .

CMD ["./eos-api"]
