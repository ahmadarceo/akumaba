FROM alpine

RUN \
  apk add --no-cache rsync
COPY rsyncd.conf /etc/rsyncd.conf

ENTRYPOINT ["rsync", "--daemon", "--no-detach", "--config=/etc/rsyncd.conf"]
