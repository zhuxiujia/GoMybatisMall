package redis

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

type RedisTemplete struct {
	ConnectionFactory *RedisConnectionFactory
}

func (it RedisTemplete) New(url string, password string) RedisTemplete {
	var factory = RedisConnectionFactory{
		Redis_url:      url,
		Redis_password: password,
	}
	var templete = RedisTemplete{}
	templete.ConnectionFactory = &factory
	return templete
}

func (templete RedisTemplete) ListLPush(key string, value string) error {
	if templete.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var conn, err = templete.ConnectionFactory.GetConnection()
	if err != nil {
		log.Println("redis set failed:", err)
		return err
	}
	defer templete.ConnectionFactory.Close(conn)
	_, doErr := conn.Do("LPUSH", key, value)
	if doErr != nil {
		log.Println("redis set failed:", doErr)
		return doErr
	}
	return nil
}

func (templete RedisTemplete) ListRPush(key string, value string) error {
	if templete.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var conn, err = templete.ConnectionFactory.GetConnection()
	if err != nil {
		log.Println("redis set failed:", err)
		return err
	}
	defer templete.ConnectionFactory.Close(conn)
	_, doErr := conn.Do("RPUSH", key, value)
	if doErr != nil {
		log.Println("redis set failed:", doErr)
		return doErr
	}
	return nil
}

func (templete RedisTemplete) ListLPop(key string) (string, error) {
	if templete.ConnectionFactory == nil {
		return "", errors.New("redis ConnectionFactory == nil")
	}
	var conn, err = templete.ConnectionFactory.GetConnection()
	defer templete.ConnectionFactory.Close(conn)
	if err != nil {
		return "", err
	}
	username, err := redis.String(conn.Do("LPOP", key))
	if err != nil {
		return "", err
	} else {
		return username, nil
	}
}

func (templete RedisTemplete) ListRPop(key string) (string, error) {
	if templete.ConnectionFactory == nil {
		return "", errors.New("redis ConnectionFactory == nil")
	}
	var conn, err = templete.ConnectionFactory.GetConnection()
	defer templete.ConnectionFactory.Close(conn)
	if err != nil {
		return "", err
	}
	username, err := redis.String(conn.Do("RPOP", key))
	if err != nil {
		return "", err
	} else {
		return username, nil
	}
}

func (templete RedisTemplete) Set(key string, value string, expireSeconds int64) error {
	if templete.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var conn, err = templete.ConnectionFactory.GetConnection()
	if err != nil {
		log.Println("redis set failed:", err)
		return err
	}
	defer templete.ConnectionFactory.Close(conn)
	_, doErr := conn.Do("SET", key, value, "EX", strconv.FormatInt(expireSeconds, 10))
	if doErr != nil {
		log.Println("redis set failed:", doErr)
		return doErr
	}
	return nil
}

func (templete RedisTemplete) Delete(key string) error {
	if templete.ConnectionFactory == nil {
		return errors.New("redis ConnectionFactory == nil")
	}
	var conn, err = templete.ConnectionFactory.GetConnection()
	if err != nil {
		fmt.Println("redis set failed:", err)
		return err
	}
	defer templete.ConnectionFactory.Close(conn)
	_, doErr := conn.Do("DEL", key)
	if doErr != nil {
		fmt.Println("redis set failed:", doErr)
		return doErr
	}
	return nil
}

func (templete RedisTemplete) Get(key string) (string, error) {
	if templete.ConnectionFactory == nil {
		return "", errors.New("redis ConnectionFactory == nil")
	}
	var conn, err = templete.ConnectionFactory.GetConnection()
	defer templete.ConnectionFactory.Close(conn)
	if err != nil {
		return "", err
	}
	username, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	} else {
		return username, nil
	}
}
