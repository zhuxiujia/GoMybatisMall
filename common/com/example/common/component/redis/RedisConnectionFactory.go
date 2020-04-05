package redis

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

//简单工厂
type RedisConnectionFactory struct {
	Redis_url      string
	Redis_password string
}

//新开一个连接
func (factory RedisConnectionFactory) GetConnection() (redis.Conn, error) {
	c, err := redis.Dial("tcp", factory.Redis_url)
	if err != nil {
		log.Println("Connect to redis error", err)
		return nil, err
	}
	if factory.Redis_password != "" {
		_, error := c.Do("AUTH", factory.Redis_password)
		if error != nil {
			return nil, error
		}
	}
	return c, nil
}

//关闭连接
func (factory RedisConnectionFactory) Close(conn redis.Conn) {
	if conn != nil {
		conn.Close()
	}
}
