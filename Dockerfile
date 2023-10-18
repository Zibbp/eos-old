FROM golang:1 AS build-stage-01

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o eos-api cmd/server/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o eos-worker cmd/worker/worker.go

FROM debian:12-slim AS production

COPY --from=build-stage-01 /app/eos-api .

COPY --from=build-stage-01 /app/eos-worker .

RUN apt update && apt install ffmpeg imagemagick -y

CMD ["./eos-api"]
