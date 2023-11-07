package session

import (
	"encoding/gob"
	"net/http"
	"os"
	"subscription-service/data"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
)

// Create sets up a session, using Redis for session store
func Create() *scs.SessionManager {
	gob.Register(data.User{})

	// set up session
	session := scs.New()
	session.Store = redisstore.New(connect())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

// connect returns a pool of connections to Redis using the
// environment variable REDIS
func connect() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}
