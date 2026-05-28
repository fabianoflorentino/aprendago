FROM golang:alpine3.23 AS development

WORKDIR /aprendago

COPY . .

RUN go install github.com/air-verse/air@latest \
  && go clean -modcache \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -o /tmp/aprendago cmd/aprendago/main.go

ENTRYPOINT [ "/go/bin/air" ]
CMD ["-c", ".air.toml"]

FROM alpine:3.23 AS production

WORKDIR /aprendago

COPY --from=development /tmp/aprendago /usr/local/bin/aprendago

ENTRYPOINT [ "/bin/sh" ]
CMD ["-c", "while true; do echo 'Im running' && sleep 300; done"]
