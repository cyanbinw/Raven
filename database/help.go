package database

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/swirling-melodies/Raven/common"
	"time"
	"xorm.io/xorm"
)

const (
	userName = ""
	password = ""
	ip       = ""
	port     = ""
	dbName   = ""
	month    = 12
	insert   = 1
	update   = 2
	delete   = 3
)

var db *sql.DB
var engine *xorm.Engine
var rdb *redis.Client

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间

func InitDB() {
	engine = common.InitDB()
}

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func V8Example() {
	ctx := context.Background()
	if err := InitRedis(); err != nil {
		return
	}

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
