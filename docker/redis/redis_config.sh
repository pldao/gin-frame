#!/bin/bash

# Connect to Redis server
redis-cli -h localhost -p 6379 -a sky <<EOF
# Set the default expiration time (2 hours = 7200 seconds)
CONFIG SET default-ttl 7200

# Alternatively, you can set a default policy for keys that don't have an explicit expiration
# CONFIG SET expire-after-write 7200

# If you want to set an initial key-value pair with the default expiration
SET my_initial_key my_value EX 7200

# Exit the Redis client
quit
EOF

#这个脚本做了以下事情：
#使用 redis-cli 连接到 Redis 服务器，指定主机 localhost、端口 6379 和密码 sky。
#通过 CONFIG SET default-ttl 7200 设置默认的过期时间（2 小时）。
#（可选）CONFIG SET expire-after-write 7200 会设置一个策略，使得所有写入的键如果没有指定过期时间，都会自动带有 2 小时的过期时间。注意，这个特性在 Redis 6.2 版本及更高版本可用。
#使用 SET 命令设置一个初始的键值对 my_initial_key，并设置其过期时间为 2 小时。
#最后，quit 命令退出 redis-cli。
#为了运行这个脚本，确保你的 Docker 容器已经启动，并且 Redis 服务在运行。然后，在同一台机器上运行这个脚本：