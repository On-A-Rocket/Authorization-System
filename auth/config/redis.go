package config

import (
	"os"

	"github.com/go-redis/redis/v8"
)

type RedisInterface interface {
	Address() string
	Password() string
	Client() *redis.Client
}

type Redis struct {
	address  string
	password string
}

func newRedisConfig() *Redis {
	redis := &Redis{
		address:  "localhost:6379",
		password: "1234",
	}

	if address := os.Getenv("REDIS_ADDRESS"); address != "" {
		redis.address = address
	}
	if password := os.Getenv("REDIS_PASSWORD"); password != "" {
		redis.password = password
	}

	return redis
}

func (r *Redis) Client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     r.address,
		Password: r.password,
		DB:       0,
	})
	if _, err := client.Ping().Result(); err != nil {
		panic(err)
	}
	return client
}

func (redis *Redis) Address() string {
	return redis.address
}

func (redis *Redis) Password() string {
	return redis.password
}
