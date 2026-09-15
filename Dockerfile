FROM python:3.14-alpine

ENV PYTHONDONTWRITEBYTECODE=1 \
	PYTHONUNBUFFERED=1

WORKDIR /app

RUN apk add --no-cache mpv

COPY . /app

RUN pip install --no-cache-dir --disable-pip-version-check .

CMD ["python", "main.py"]
