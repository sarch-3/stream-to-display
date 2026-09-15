# stream-to-display

```
docker build -t stream-to-display .
```

```
docker run --rm \
  --device=/dev/dri/card1 \
  --device=/dev/snd:/dev/snd \
  -e DRM_DEVICE=/dev/dri/card1 \
  -e DRM_CONNECTOR=HDMI-A-1 \
  -e AUDIO_OUTPUT=alsa \
  -e AUDIO_DEVICE=alsa/hw:0,3 \
  -e CACHE=1 \
  -e LOOP=inf \
  -e STREAM_RESOLUTION=1080 \
  -e STREAM_URL="https://www.youtube.com/watch?v=dQw4w9WgXcQ" \
  stream-to-display
```

```
docker run --rm \
  --device=/dev/dri/card1 \
  --device=/dev/dri/renderD128 \
  --group-add "$(stat -c '%g' /dev/dri/card1)" \
  -v "$PWD/test.mp4:/app/test.mp4:ro" \
  -e STREAM_URL=/app/test.mp4 \
  -e DRM_DEVICE=/dev/dri/card1 \
  -e DISPLAY_NAME=HDMI-A-1 \
  stream-to-display
```
