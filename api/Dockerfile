FROM golang:1.20.13-alpine AS builder

WORKDIR /app
ADD . .

ARG env
ARG job
RUN echo "Build $env"

ENV GO111MODULE=on

ADD ".env.$env" ".env"
RUN go get -v
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o api .

FROM alpine

RUN apk update && apk add ca-certificates && apk add --update coreutils && rm -rf /var/cache/apk/*
RUN apk add -U tzdata

COPY --from=builder /app/api .
COPY --from=builder /app/data.json data.json
COPY --from=builder /app/iller_su.json iller_su.json
COPY --from=builder /app/docs /docs
COPY --from=builder /app/public /public
COPY --from=builder /app/.env .env
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip

CMD ["./api"]
EXPOSE 3000
