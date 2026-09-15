# stream-to-display

```
docker build \
  --build-arg APP_HOST=0.0.0.0 \
  --build-arg APP_PORT=8080 \
  --build-arg VIDEO_OUTPUT=drm \
  --build-arg DRM_DEVICE=/dev/dri/card1 \
  --build-arg DRM_CONNECTOR=eDP-1 \
  --build-arg AUDIO_OUTPUT=alsa \
  --build-arg AUDIO_DEVICE=alsa/hw:0,3 \
  --build-arg HWDEC=vaapi \
  --build-arg CACHE=true \
  --build-arg DEMUXER_MAX_BYTES=102400 \
  --build-arg DEMUXER_READAHEAD_SECS=60 \
  -t stream-to-display .
```

```
docker run --rm \
  -p 8080:8080 \
  --device /dev/dri:/dev/dri \
  --device=/dev/dri/renderD128 \
  --device /dev/snd:/dev/snd \
  stream-to-display
```
