FROM golang:1.14.3-alpine

ENV GIN_MODE=release
ENV PORT=5491

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
RUN mkdir -p /gauguin
WORKDIR /gauguin

RUN apk update && apk upgrade && apk add --no-cache bash git && apk add --no-cache chromium

RUN echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories \
  && echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories \
  && apk add --no-cache \
  harfbuzz@edge \
  nss@edge \
  freetype@edge \
  ttf-freefont@edge \
  && rm -rf /var/cache/* \
  && mkdir /var/cache/apk

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./gauguin

EXPOSE $PORT

RUN chromium-browser --headless --disable-gpu --remote-debugging-port=9222 --disable-web-security --safebrowsing-disable-auto-update --disable-sync --disable-default-apps --hide-scrollbars --metrics-recording-only --mute-audio --no-first-run --no-sandbox

ENTRYPOINT ["./gauguin"]