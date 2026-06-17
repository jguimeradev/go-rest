APP_NAME := go-rest
CMD_DIR  := ./cmd
BIN_DIR  := ./bin

.PHONY: build run clean vet

build:
	go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

run: build
	$(BIN_DIR)/$(APP_NAME)

clean:
	rm -f $(BIN_DIR)/$(APP_NAME)

vet:
	go vet ./...
