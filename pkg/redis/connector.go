package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func Dial(network string, addr string, port string) (redis.Conn, error) {
	dialAddress := fmt.Sprintf("%s:%s", addr, port)
	return redis.Dial(network, dialAddress)
}
