FROM docker.m.daocloud.io/alpine:3.18.6

LABEL maintainer = "18395806407@163.com"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
      apk add --no-cache zip alpine-conf wget

COPY  htpasswd /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/htpasswd"]
