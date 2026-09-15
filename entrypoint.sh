#!/bin/sh

# Запускаем mpv в фоне (убираем --no-terminal, чтобы в логах Docker было видно, если он упадет!)
mpv --no-config \
    --idle=yes \
    --input-ipc-server="$SOCKET_PATH" \
    --vo="$VIDEO_OUTPUT" \
    --drm-device="$DRM_DEVICE" \
    --drm-connector="$DRM_CONNECTOR" \
    --ao="$AUDIO_OUTPUT" \
    --audio-device="$AUDIO_DEVICE" \
    --hwdec="$HWDEC" \
    --cache="$CACHE" \
    --demuxer-max-bytes="$DEMUXER_MAX_BYTES" \
    --demuxer-readahead-secs="$DEMUXER_READAHEAD_SECS" &

MPV_PID=$!

echo "Ждем появления сокета $SOCKET_PATH..."
# Ждем максимум 5 секунд, проверяя, жив ли процесс mpv
for i in $(seq 1 50); do
    if [ -S "$SOCKET_PATH" ]; then
        echo "mpv успешно стартовал, сокет готов!"
        break
    fi
    if ! kill -0 $MPV_PID 2>/dev/null; then
        echo "ОШИБКА: mpv упал на старте! Проверь параметры железа."
        exit 1
    fi
    sleep 0.1
done

# Передаем управление твоему приложению
exec /usr/local/bin/stream-to-display
