FROM golang:1.14.3-alpine as build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./gauguin

FROM golang:1.14.3-alpine

ENV DOCKERIZED=true
ENV GIN_MODE=release
ENV PORT=5491

WORKDIR /app

COPY --from=build /app/gauguin gauguin

EXPOSE $PORT

CMD ["./gauguin"]