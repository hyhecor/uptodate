FROM golang:1.26 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o uptodate .

FROM alpine:3.23

RUN apk --no-cache add ca-certificates

COPY --from=builder /src/uptodate /usr/local/bin

ENTRYPOINT ["uptodate"]