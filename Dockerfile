FROM golang:1.18 AS builder

ADD ./ /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /bin/app ./...

FROM golang:1.18 AS runner

WORKDIR /app
COPY --from=builder /bin/app /app

CMD ["/app/app"]