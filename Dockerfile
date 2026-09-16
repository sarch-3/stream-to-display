FROM golang:1.27-bookworm AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/stream-to-display ./cmd/app

FROM debian:bookworm-slim

RUN apt-get update \
	&& apt-get install -y --no-install-recommends \
		mpv \
		wget \
		i965-va-driver \
		intel-media-va-driver \
		mesa-va-drivers \
		libgl1-mesa-dri \
		alsa-utils \
		ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

RUN wget https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O /usr/local/bin/yt-dlp \
    && chmod a+rx /usr/local/bin/yt-dlp


COPY --from=builder /out/stream-to-display /usr/local/bin/stream-to-display

USER root

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
