.PHONY: build run clean

BINARY_NAME=crossword
BUILD_DIR=bin

# Сборка проекта
build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/crossword/main.go

# Запуск
run: build
	$(BUILD_DIR)/$(BINARY_NAME) datasets/words35.txt

# Очистка
clean:
	rm -rf $(BUILD_DIR)
	rm -f cpu.prof solution.txt
