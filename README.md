# stream-to-display

```
docker build -t stream-to-display .
```

```
docker run --rm \
  -p 8080:8080 \
  --device /dev/dri:/dev/dri \
  --device /dev/snd:/dev/snd \
  -e APP_HOST=0.0.0.0 \
  -e APP_PORT=8080 \
  -e VIDEO_OUTPUT=drm \
  -e DRM_DEVICE=/dev/dri/card0 \
  -e DRM_CONNECTOR=HDMI-A-1 \
  -e AUDIO_OUTPUT=alsa \
  -e AUDIO_DEVICE=alsa/hw:0,3 \
  -e HWDEC=vaapi \
  -e CACHE=true \
  -e DEMUXER_MAX_BYTES=102400 \
  -e DEMUXER_READAHEAD_SECS=60 \
  stream-to-display
```
