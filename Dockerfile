FROM golang:alpine3.23 AS builder

WORKDIR /aprendago

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -o /tmp/aprendago cmd/aprendago/main.go


FROM golang:alpine3.23 AS development

WORKDIR /aprendago

COPY . .

RUN go install github.com/air-verse/air@latest

COPY --from=builder /tmp/aprendago /usr/local/bin/aprendago

ENTRYPOINT [ "/go/bin/air" ]
CMD ["-c", ".air.toml"]

FROM gcr.io/distroless/cc:nonroot AS production

WORKDIR /aprendago

COPY --from=builder /tmp/aprendago /usr/local/bin/aprendago

USER nonroot:nonroot

ENTRYPOINT [ "/bin/sh" ]
CMD ["-c", "while true; do echo 'Im running' && sleep 300; done"]
