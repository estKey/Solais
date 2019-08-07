package rdb

import (
	"Solais/utils/log"

	"github.com/gomodule/redigo/redis"
)

func Set(addr string, key string, value string) {
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		log.Error.Println("Connect to redis error", err)
		return
	}
	_, err = c.Do("SET", key, value, "NX")
	if err != nil {
		log.Error.Println("redis set failed:", err)
	}
	defer c.Close()
}

func Get(addr string, key string) string {
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		log.Error.Println("Connect to redis error:", err)
		return "err"
	}
	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		log.Error.Println("redis get failed:", err)
		return "err"
	} else {
		log.Error.Println("Get %v: %v \n", key, value)
	}
	defer c.Close()
	return value
}

func Update(addr string, key string, value string) {
	if addr == "" {
		addr = "127.0.0.1:6379"
	}
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		log.Error.Println("Connect to redis error:", err)
		return
	}
	_, err = c.Do("SET", key, value, "XX")
	if err != nil {
		log.Error.Println("redis update failed:", err)
	}
	defer c.Close()
}

