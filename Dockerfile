FROM golang:1.20rc3-alpine3.17 as builder

WORKDIR /app

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY . .
RUN go mod download
RUN go build -ldflags="-w -s" -o app ./cmd/api

FROM alpine:3.17.4

WORKDIR /app

ENV TZ="UTC"

#COPY --from=builder /app/config.yaml /app/config.yaml
COPY --from=builder /app/app /app/app
CMD [ "./app" ]
