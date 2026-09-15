# stream-to-display

```
docker build -t stream-to-display .
```

```
docker run --rm \
  --device=/dev/dri/card1 \
  --device=/dev/dri/renderD128 \
  --device=/dev/snd:/dev/snd \
  -e DRM_DEVICE=/dev/dri/card1 \
  -e DRM_CONNECTOR=HDMI-A-1 \
  -e AUDIO_OUTPUT=alsa \
  -e AUDIO_DEVICE=alsa/hw:0,3 \
  -e CACHE=1 \
  -e DEMUXER_MAX_BYTES=100M \
  -e DEMUXER_READAHEAD_SECS=60 \
  -e LOOP=inf \
  -e STREAM_RESOLUTION=1080 \
  -e STREAM_URL="https://www.youtube.com/watch?v=dQw4w9WgXcQ" \
  stream-to-display
```

Не поддерживается апаратное ускорение nvidia
