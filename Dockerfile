FROM golang:1.22.1-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN mkdir -p /tmp

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o quake-parser .

FROM scratch

COPY --from=builder /app/quake-parser /quake-parser

COPY --from=builder /app/report /report

COPY --from=builder /tmp /tmp

WORKDIR /

ENTRYPOINT ["/quake"]