FROM docker.io/centos:latest

LABEL maintainer="Agateram"

ADD redis-4.0.10.tar.gz /redis/
ENV PATH $PATH:/redis/redis-4.0.10/src/

EXPOSE 6379
CMD [ "/redis/redis-4.0.10/src/redis-server","/redis/redis-4.0.10/redis.conf" ]