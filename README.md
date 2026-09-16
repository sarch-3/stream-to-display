# stream-to-display

Собрать образ:

```bash
docker build -t stream-to-display .
```

Настройки находятся в `.env`. Создайте его из `.env.example` и измените нужные значения.

```bash
docker run --rm \
	--env-file .env \
  -p 8080:8080 \
  --device=/dev/dri
  --device=dev/snd \
  stream-to-display
```
