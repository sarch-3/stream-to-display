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
		yt-dlp \
		i965-va-driver \
		intel-media-va-driver \
		mesa-va-drivers \
		libgl1-mesa-dri \
		alsa-utils \
		ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

COPY --from=builder /out/stream-to-display /usr/local/bin/stream-to-display

ENV APP_HOST=0.0.0.0 \
	APP_PORT=8080 \
	SOCKET_PATH=/tmp/mpvsocket \
	VIDEO_OUTPUT=drm \
	DRM_DEVICE=/dev/dri/card0 \
	DRM_CONNECTOR=HDMI-A-1 \
	AUDIO_OUTPUT=alsa \
	AUDIO_DEVICE=default \
	HWDEC=vaapi \
	CACHE=true \
	DEMUXER_MAX_BYTES=1048576 \
	DEMUXER_READAHEAD_SECS=10

EXPOSE 8080

USER root

CMD ["sh", "-c", "mpv --no-config --idle=yes --no-terminal --input-ipc-server=\"$SOCKET_PATH\" --vo=\"$VIDEO_OUTPUT\" --drm-device=\"$DRM_DEVICE\" --drm-connector=\"$DRM_CONNECTOR\" --ao=\"$AUDIO_OUTPUT\" --audio-device=\"$AUDIO_DEVICE\" --hwdec=\"$HWDEC\" --cache=\"$CACHE\" --demuxer-max-bytes=\"$DEMUXER_MAX_BYTES\" --demuxer-readahead-secs=\"$DEMUXER_READAHEAD_SECS\" & exec /usr/local/bin/stream-to-display"]
