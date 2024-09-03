package app

import (
	"github.com/https_whoyan/Lets_Go_Book_Course/cmd/flag"
	myRedis "github.com/https_whoyan/Lets_Go_Book_Course/pkg/redis"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
)

const (
	flashKey = "flash"
	authKey  = "auth"
)

func (app *Application) standSessionManager(flagCfg *flag.FlagConfig) {
	const network = "tcp"
	pool := &redis.Pool{
		MaxIdle: 3,
		Dial: func() (redis.Conn, error) {
			return myRedis.Dial(network, flagCfg.RedisNetwork, flagCfg.RedisAddr)
		},
	}

	sessionManager := scs.New()
	sessionManager.Store = redisstore.New(pool)
	sessionManager.Cookie.Secure = true

	app.sessionManager = sessionManager
}
