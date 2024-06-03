#!/bin/bash
rm -rf data
# 启动 Redis 服务
echo "Starting Redis..."
cd ./redis
docker-compose up -d
echo "Redis started successfully!"

# 启动 MongoDB 服务
echo "Starting MongoDB..."
cd ../mongodb
docker-compose up -d
echo "MongoDB started successfully!"

# 启动 MySQL 服务
echo "Starting MySQL..."
cd ../mysql
docker-compose up -d
echo "MySQL started successfully!"

# 启动 Minio 服务
echo "Starting Minio..."
cd ../minio
docker-compose up -d
echo "Minio started successfully!"

echo "All services started successfully!"
