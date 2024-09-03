package app

import (
	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/https_whoyan/Lets_Go_Book_Course/cmd/flag"

	"github.com/gomodule/redigo/redis"
)

func (app *Application) standSessionManager(flagCfg *flag.FlagConfig) {
	const network = "tcp"
	address := flagCfg.RedisNetwork + ":" + flagCfg.RedisAddr
	pool := &redis.Pool{
		MaxIdle: 3,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(network, address)
		},
	}

	sessionManager := scs.New()
	sessionManager.Store = redisstore.New(pool)

	app.sessionManager = sessionManager
	app.infoLogger.Printf("ok")
}
