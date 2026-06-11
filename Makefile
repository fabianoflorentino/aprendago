.PHONY: default build run install fmt vet lint tidy clean \
        test test/race test/short test/chapter test/cover test/cover/html bench \
        build/linux build/darwin build/windows \
        docker/build docker/run \
        caps outline cap help

BINARY  = aprendago
CMD_DIR = ./cmd/aprendago
OUT_DIR = ./out

# ─── Default ──────────────────────────────────────────────────────────────────

default: help

# ─── Build ────────────────────────────────────────────────────────────────────

build: $(BINARY)

$(BINARY):
	go build -o $(BINARY) $(CMD_DIR)

build/linux: GOOS = linux
build/linux: GOARCH = amd64
build/linux: _build_cross

build/darwin: GOOS = darwin
build/darwin: GOARCH = amd64
build/darwin: _build_cross

build/windows: GOOS = windows
build/windows: GOARCH = amd64
build/windows: _build_cross

_build_cross:
	@mkdir -p $(OUT_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUT_DIR)/$(BINARY)-$(GOOS)-$(GOARCH) $(CMD_DIR)
	@echo "built: $(OUT_DIR)/$(BINARY)-$(GOOS)-$(GOARCH)"

run: build
	./$(BINARY) $(ARGS)

install:
	go install $(CMD_DIR)

# ─── Code Quality ─────────────────────────────────────────────────────────────

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	@which golangci-lint > /dev/null 2>&1 && golangci-lint run ./... || echo "golangci-lint not installed — skipping"

tidy:
	go mod tidy

clean:
	rm -rf $(BINARY) $(OUT_DIR)
	go clean -cache

# ─── Test ─────────────────────────────────────────────────────────────────────

test:
	go test ./... -count=1

test/race:
	go test ./... -race -count=1

test/short:
	go test ./... -short -count=1

test/chapter:
	go test ./internal/chapter/... -v -count=1

test/cover:
	go test ./... -coverprofile=coverage.out -count=1
	@go tool cover -func=coverage.out | grep total:

test/cover/html: test/cover
	@go tool cover -html=coverage.out -o coverage.html
	@echo "coverage report: coverage.html"

bench:
	go test ./... -bench=. -benchmem -count=1

# ─── Docker ───────────────────────────────────────────────────────────────────

docker/build:
	docker compose build --no-cache

docker/run:
	docker compose up -d

# ─── CLI Shortcuts ────────────────────────────────────────────────────────────

caps:
	go run $(CMD_DIR) caps

outline:
	go run $(CMD_DIR) outline

cap:
	go run $(CMD_DIR) cap $(or $(N),1) $(or $(M),topics)

# ─── Help ─────────────────────────────────────────────────────────────────────

help:
	@echo "╔══════════════════════════════════════════╗"
	@echo "║        aprendago — Makefile              ║"
	@echo "╚══════════════════════════════════════════╝"
	@echo ""
	@echo "  Build"
	@echo "    make build              compila binario"
	@echo "    make run ARGS=...       roda com argumentos"
	@echo "    make install            go install"
	@echo "    make build/linux        cross-compile linux/amd64"
	@echo "    make build/darwin       cross-compile darwin/amd64"
	@echo "    make build/windows      cross-compile windows/amd64"
	@echo ""
	@echo "  Test"
	@echo "    make test               roda todos os testes"
	@echo "    make test/race          com detector de condicao de corrida"
	@echo "    make test/short         modo rapido (-short)"
	@echo "    make test/chapter       tests do pacote chapter"
	@echo "    make test/cover         exibe cobertura"
	@echo "    make test/cover/html    abre relatorio HTML"
	@echo "    make bench              benchmarks"
	@echo ""
	@echo "  Code Quality"
	@echo "    make fmt                go fmt"
	@echo "    make vet                go vet"
	@echo "    make lint               golangci-lint (se instalado)"
	@echo "    make tidy               go mod tidy"
	@echo "    make clean              limpa artefatos"
	@echo ""
	@echo "  Docker"
	@echo "    make docker/build       docker compose build"
	@echo "    make docker/run         docker compose up -d"
	@echo ""
	@echo "  CLI Shortcuts"
	@echo "    make caps               lista capitulos"
	@echo "    make outline            outline completo"
	@echo "    make cap                capitulo 1 topics"
	@echo "    make cap N=3            capitulo N topics"
	@echo "    make cap N=3 M=overview overview do capitulo"
	@echo "    make cap N=3 M=\"topico\" topico especifico"
	@echo ""
	@echo "  Exemplos"
	@echo "    make run ARGS=\"cap 5 overview\""
	@echo "    make run ARGS=\"cap 8 \\\"Agrupamento de Dados\\\"\""
