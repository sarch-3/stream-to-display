FROM python:3.14-alpine

ENV PYTHONDONTWRITEBYTECODE=1 \
	PYTHONUNBUFFERED=1

WORKDIR /app

RUN apk add --no-cache mpv

COPY ./requirements.txt /app

RUN pip install --no-cache-dir --disable-pip-version-check .

COPY . /app

CMD ["python", "main.py"]
