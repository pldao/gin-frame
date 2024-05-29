# Makefile

# 设置变量
GO_CMD := go
MAIN_PKG := main.go
GO_BUILD := debug
SERVER_CMD := $(GO_CMD) run $(MAIN_PKG)
DB_MIGRATE_CMD := migrate -database 'mysql://root:sky@tcp(127.0.0.1:3307)/ginframe?charset=utf8mb4&parseTime=True&loc=Local'
MIGRATION_PATH := data/migrations
GOLINT=golangci-lint

# 目标规则
all: run

# 启动后端程序
run:
	$(SERVER_CMD) server
# 将API全部存入数据库
store-api:
	$(SERVER_CMD) server -R true

# 下载migrate工具（假设使用go get方式）
get-migrate:
	$(GO_CMD) install -u github.com/golang-migrate/migrate/cli/migrate

# 数据库迁移 up
migrate-up:
	$(DB_MIGRATE_CMD) -path $(MIGRATION_PATH) up

# 数据库迁移 down
migrate-down:
	$(DB_MIGRATE_CMD) -path $(MIGRATION_PATH) down

install:
	go mod tidy
	go mod verify
	go mod download

swag:
	swag init -g main.go

clean:

###############################################################################
###                                Linting                                  ###
###############################################################################

lint-install:
	@echo "--> Checking if golangci-lint is installed or needs updating"
	@if ! golangci-lint version --short 2>&1 | grep -q 'golangci-lint'; then \
		echo "golangci-lint not found. Installing the latest version"; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		current_version=$(golangci-lint version --short); \
		latest_version=$(curl -s https://api.github.com/repos/golangci/golangci-lint/releases/latest | jq -r '.tag_name'); \
		if [[ "$${current_version}" != "$${latest_version}" ]]; then \
			echo "Updating golangci-lint from $${current_version} to $${latest_version}"; \
			go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		else \
			echo "golangci-lint is already up-to-date at version $${current_version}"; \
		fi \
	fi

lint: lint-install
	@echo "--> Running linter"
	$(GOLINT) run --build-tags=$(GO_BUILD) --out-format=tab

format: lint-install
	@golangci-lint run --build-tags=$(GO_BUILD) --out-format=tab --fix

# 使用说明
help:
	@echo "Makefile Usage:"
	@echo "  make run - 启动后端程序"
	@echo "  make store-api - 将API存入数据库"
	@echo "  make get-migrate - 下载migrate工具"
	@echo "  make migrate-up - 数据库迁移 up"
	@echo "  make migrate-down - 数据库迁移 down"
	@echo "  make clean - 清理"
	@echo "  make help - 显示帮助信息"
	@echo "  make lint-install - 安装lint工具"
	@echo "  make lint - 运行lint工具"
	@echo "  make format - 格式化代码"


.PHONY: all run store-api get-migrate migrate-up migrate-down clean help lint-install lint format
