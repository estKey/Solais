FROM redis:latest

LABEL maintainer="Agateram"

EXPOSE 6379

COPY redis.conf /etc/redis/redis.conf
RUN redis-server /etc/redis/redis.conf
