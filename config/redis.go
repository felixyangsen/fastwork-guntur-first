package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func ConnectRDB() {
	var (
		rdsn        = fmt.Sprintf("%s:%s", os.Getenv("RDB_HOST"), os.Getenv("RDB_PORT"))
		username    = os.Getenv("RDB_USERNAME")
		password    = os.Getenv("RDB_PASSWORD")
		database, _ = strconv.Atoi(os.Getenv("RDB_DATABASE"))
	)

	rdb = redis.NewClient(&redis.Options{
		Addr:     rdsn,
		Username: username,
		Password: password, // no password set
		DB:       database, // use default DB
	})

	log.Println("Connected to redis database")
}

func GetRDB() *redis.Client {
	return rdb
}
