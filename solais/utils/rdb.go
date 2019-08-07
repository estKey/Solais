package rdb

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func Connect(addr string) {
	if addr == "" {
		name = "127.0.0.1:6379"
	}
	c, err := redis.Dial("tcp", addr)
	if err != nil {
			fmt.Println("Connect to redis error", err)
			return
	}
	defer c.Close()
}

func Set(key string, value string) {
	_, err = c.Do("SET", key, value)
	if err != nil {
			fmt.Println("redis set failed:", err)
	}

}

func Get(key string) {
	username, err := redis.String(c.Do("GET", key))
	if err != nil {
			fmt.Println("redis get failed:", err)
	} else {
			fmt.Printf("Get %v: %v \n", key, username)
	}
}