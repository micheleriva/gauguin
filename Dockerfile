FROM golang:1.14.3-alpine

ENV DOCKERIZED=true
ENV GIN_MODE=release
ENV PORT=5491

RUN mkdir -p /gauguin
WORKDIR /gauguin

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./gauguin

EXPOSE $PORT

ENTRYPOINT ["./gauguin"]