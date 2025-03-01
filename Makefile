#!/bin/bash
APP_NAME = scheduleassist
VERSION = 1.0.0
BUILD_DIR = ./build
CMD_DIR = ./cmd

TEST_COVERAGE_FILE=coverage.out

LDFLAGS = "-s -w"

PLATFORMS := \
  linux/amd64 \
  linux/arm64 \
  darwin/amd64 \
  darwin/arm64

.PHONY: all build run clean test help

help:
	@echo "Доступные команды:"
	@echo "  make build     - Скомпилировать проект"
	@echo "  make run       - Скомпилировать и запустить проект"
	@echo "  make clean     - Удалить скомпилированные файлы"
	@echo "  make test      - Запустить тесты"
	@echo "  make lint      - Проверить код с помощью linters"
	@echo "  make help      - Показать это сообщение"

all: clean build

build:
	@echo "Project building..."
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		echo "Building for the platform '$$platform'"; \
		GOOS=$${platform%/*}; \
		GOARCH=$${platform##*/}; \
		go build -tags=postgres -ldflags=$(LDFLAGS) -o $(BUILD_DIR)/$$platform/$(APP_NAME) $(CMD_DIR); \
		tar -cf ${BUILD_DIR}/${APP_NAME}_$${GOOS}_$${GOARCH}.tar.gz -C $(BUILD_DIR)/$$platform .; \
	done
	@echo "Done."

clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(TEST_COVERAGE_FILE) coverage.html
	@echo "Done."

test:
	@echo "Запуск тестов..."
	go test  -cover ./...

coverage:
	go test -coverprofile=$(TEST_COVERAGE_FILE) ./...
	go tool cover -func=$(TEST_COVERAGE_FILE)

# Открыть отчет в HTML
coverage-html: coverage
	go tool cover -html=$(TEST_COVERAGE_FILE) -o coverage.html
	open coverage.html