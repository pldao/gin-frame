# Makefile

# 设置变量
GO_CMD := go
MAIN_PKG := main.go
SERVER_CMD := $(GO_CMD) run $(MAIN_PKG)
DB_MIGRATE_CMD := migrate -database 'mysql://root:sky@tcp(127.0.0.1:3307)/ginframe?charset=utf8mb4&parseTime=True&loc=Local'
MIGRATION_PATH := data/migrations

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

# 清理（可根据需要定义）
clean:
	# 这里可以添加清理临时文件或构建产物的命令

# 使用说明
help:
	@echo "Makefile Usage:"
	@echo "  make run - 启动后端程序"
	@echo "  make store-api - 将API存入数据库"
	@echo "  make get-migrate - 下载migrate工具"
	@echo "  make migrate-up - 数据库迁移 up"
	@echo "  make migrate-down - 数据库迁移 down"
	@echo "  make clean - 清理"

.PHONY: all run store-api get-migrate migrate-up migrate-down clean help
