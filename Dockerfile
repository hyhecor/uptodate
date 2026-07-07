FROM golang:1.26 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o uptodate .

FROM alpine:3.23

RUN apk --no-cache add ca-certificates

COPY --from=builder /src/uptodate /bin/uptodate

ENTRYPOINT ["uptodate"]
