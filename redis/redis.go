package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.99.100:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
	} else {
		fmt.Println("Connect to redis success")
	}

	defer c.Close()
	// set key value
	n, err := c.Do("APPEND", "key", "value")
	_ = n

	key, err := redis.String(c.Do("GET", "key"))
	fmt.Println("return", key)
}
