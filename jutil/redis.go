package jutil

import (
	"github.com/go-redis/redis"
)

//redis客户端
type RedisClient struct {
	client *redis.Client
}

//连接
func (c *RedisClient) Connect() {
	config := NewConfig()
	address := config.GetString("redis.default.host")
	port := config.GetString("redis.default.port")
	password := config.GetString("redis.default.password")
	client := redis.NewClient(&redis.Options{
		Addr:        address + ":" + port,
		Password:    password,
		DB:          0,
		ReadTimeout: -1,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())
	}
	if pong != "PONG" {
		panic("pong error")
	}
	c.client = client

}

//获取所有的KEY
func (c *RedisClient) GetAllKeys() []interface{} {
	client := c.client
	command := client.Do("KEYS", "*")
	_ = client.Process(command)
	keys, err1 := command.Result()
	if err1 != nil {
		return nil
	}
	return keys.([]interface{})
}

//获取实例
func NewRedisClient() *RedisClient {
	redisClient := &RedisClient{}
	redisClient.Connect()
	return redisClient
}
