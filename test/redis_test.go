package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
Addr:     "localhost:6379",
Password: "", // no password set
DB:       0,  // use default DB
})

func TestSetRedis(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", time.Second*20).Err()
	if err != nil {
		log.Println(err)
	}
}

func TestGetRedis(t *testing.T) {
	result, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("key: ",result)
}