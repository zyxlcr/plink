package app

import (
	"chatcser/config"

	red "github.com/redis/go-redis/v9"
)

func Reids() *red.Client {
	r := config.GVA_CONFIG.Redis
	return red.NewClient(&red.Options{
		Addr:     r.Addr, //"localhost:6379",
		Password: r.Pass, // no password set
		DB:       0,      // use default DB
	})
}
