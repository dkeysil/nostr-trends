# Build

FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /nostr-trends ./cmd/trends

# Deploy

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /nostr-trends /nostr-trends

USER nonroot:nonroot

ENTRYPOINT ["/nostr-trends"]