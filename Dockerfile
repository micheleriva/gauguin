FROM golang:1.14.3-alpine

ENV DOCKERIZED=true
ENV GIN_MODE=release
ENV PORT=5491

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./gauguin

EXPOSE $PORT

CMD ["./gauguin"]