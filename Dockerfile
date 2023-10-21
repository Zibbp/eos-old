FROM golang:1 AS build-stage-01

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o eos-api cmd/server/main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o eos-worker cmd/worker/worker.go

FROM debian:12-slim AS production

ENV FONTCONFIG_CACHE=/tmp

COPY --from=build-stage-01 /app/eos-api .

COPY --from=build-stage-01 /app/eos-worker .

RUN apt update && apt install curl imagemagick -y

COPY .docker/policy.xml /etc/ImageMagick-6/policy.xml

# jellyfin's ffmpeg has better support for hardware acceleration
RUN curl -L -o ffmpeg.deb https://github.com/jellyfin/jellyfin-ffmpeg/releases/download/v6.0-7/jellyfin-ffmpeg6_6.0-7-bookworm_amd64.deb

RUN apt install --fix-broken ./ffmpeg.deb -y

RUN rm ffmpeg.deb

RUN ln -s /usr/lib/jellyfin-ffmpeg/ffmpeg /usr/local/bin/ffmpeg && ln -s /usr/lib/jellyfin-ffmpeg/ffprobe /usr/local/bin/ffprobe

CMD ["./eos-api"]
