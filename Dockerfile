FROM golang:1.17-buster as builder

RUN apt-get update && \
    apt-get install -y git ca-certificates tzdata && \
    update-ca-certificates

WORKDIR /app

ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN export GOPROXY=https://proxy.golang.org && go mod download -x
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o go_playlist \
    ./app/server.go

FROM debian:buster-20200908-slim

ENV USER=appuser
ENV UID=10001
ENV TZ=Asia/Jakarta

COPY ./config/app.env app/app.env

RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

COPY --from=builder --chown=appuser:appuser /app/go_playlist .
COPY --from=builder --chown=appuser:appuser /app/config ./config
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
RUN apt-get update; \
    apt-get upgrade; \
    cp /usr/share/zoneinfo/${TZ} /etc/localtime; \
    date;

RUN echo Y || apt-get install curl

STOPSIGNAL SIGINT

EXPOSE 5011

ENTRYPOINT ["./go_playlist"]