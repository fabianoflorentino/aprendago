.PHONY: default build run test test/chapter test/short vet clean tidy caps outline cap help

BINARY  = aprendago
CMD_DIR = ./cmd/aprendago

# ─── Default ──────────────────────────────────────────────────────────────────

default: help

# ─── Build ────────────────────────────────────────────────────────────────────

build: $(BINARY)

$(BINARY):
	go build -o $(BINARY) $(CMD_DIR)

run: $(BINARY)
	./$(BINARY) $(ARGS)

# ─── Test ──────────────────────────────────────────────────────────────────────

test:
	go test ./... -count=1

test/chapter:
	go test ./internal/chapter/... -v -count=1

test/short:
	go test ./... -short -count=1

# ─── Code Quality ──────────────────────────────────────────────────────────────

vet:
	go vet ./...

tidy:
	go mod tidy

clean:
	rm -f $(BINARY)
	go clean -cache

# ─── Convenience Shortcuts ────────────────────────────────────────────────────

caps: ARGS = --caps
caps: run

outline: ARGS = --outline
outline: run

cap: ARGS = --cap=$(or $(N),1) --topics
cap: run

# ─── Help ─────────────────────────────────────────────────────────────────────

help:
	@echo "aprendago — Makefile"
	@echo ""
	@echo "  Build"
	@echo "    make build          compila o binario"
	@echo "    make run ARGS=..    roda com argumentos"
	@echo ""
	@echo "  Test"
	@echo "    make test           roda todos testes"
	@echo "    make test/chapter   testes do chapter"
	@echo "    make test/short     testes sem lentidao"
	@echo ""
	@echo "  Code Quality"
	@echo "    make vet            go vet"
	@echo "    make tidy           go mod tidy"
	@echo "    make clean          limpa artefatos"
	@echo ""
	@echo "  Shortcuts"
	@echo "    make caps           lista capitulos"
	@echo "    make outline        outline completo"
	@echo "    make cap N=3        capitulo N topics"
	@echo ""
	@echo "  Exemplo personalizado:"
	@echo "    make run ARGS=\"--cap=5 --overview\""
	@echo "    make run ARGS=\"cap 3 overview\""
