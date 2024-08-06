package flag

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

type FlagConfig struct {
	NetAddr string
	DbAddr  string

	RedisNetwork  string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

var (
	defaultNewAddr string
	defaultDbAddr  string

	defaultRedisNetwork  string
	defaultRedisAddr     string
	defaultRedisPassword string
)

func NewFlagConfig() (*FlagConfig, error) {
	if err := StandEnv(); err != nil {
		return nil, err
	}

	netAddr := flag.String("addr", defaultNewAddr, "The network address of the database server")
	dbAddr := flag.String("dsn", defaultDbAddr, "The database server address")

	redisNetwork := flag.String("network", defaultRedisNetwork, "The network address of the redis server")
	redisAddr := flag.String("redis", defaultRedisAddr, "The address of the redis server")
	redisPassword := flag.String("password", defaultRedisPassword, "The password of the redis server")

	flag.Parse()
	return &FlagConfig{
		NetAddr:       *netAddr,
		DbAddr:        *dbAddr,
		RedisNetwork:  *redisNetwork,
		RedisAddr:     *redisAddr,
		RedisPassword: *redisPassword,
	}, nil
}

func StandEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	defaultNewAddr = os.Getenv("DEFAULT_NET_ADDR")
	defaultDbAddr = os.Getenv("DEFAULT_DB_ADDR")

	defaultRedisNetwork = os.Getenv("DEFAULT_REDIS_HOST")
	defaultRedisAddr = os.Getenv("DEFAULT_REDIS_PORT")
	defaultRedisPassword = os.Getenv("DEFAULT_REDIS_PASSWORD")
	return nil
}
