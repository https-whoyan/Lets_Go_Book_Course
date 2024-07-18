package flag

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

type FlagConfig struct {
	NetAddr string
	DbAddr  string
}

var (
	defaultNewAddr string
	defaultDbAddr  string
)

func NewFlagConfig() (*FlagConfig, error) {
	if err := StandEnv(); err != nil {
		return nil, err
	}

	netAddr := flag.String("addr", defaultNewAddr, "The network address of the database server")
	dbAddr := flag.String("dsn", defaultDbAddr, "The database server address")
	flag.Parse()
	return &FlagConfig{
		NetAddr: *netAddr,
		DbAddr:  *dbAddr,
	}, nil
}

func StandEnv() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	defaultNewAddr = os.Getenv("DEFAULT_NET_ADDR")
	defaultDbAddr = os.Getenv("DEFAULT_DB_ADDR")
	return nil
}
