package cache

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type CacheInterface interface {
	Ping() (err error)
	Get(key string) (val []byte, err error)
	GetAllKeys(buyerId string) (keys []string, err error)
	Write(key string, data interface{}, ttl time.Duration) (err error)
	DeleteRecord(key string) (err error)
	FlushByBuyerId(buyerId string) (err error)
}

type Cache struct {
	pool *redis.Pool
}

var Redis CacheInterface
var Rpool *redis.Pool

func NewCache() CacheInterface {
	dialConnectTimeoutOption := redis.DialConnectTimeout(5 * time.Second)
	readTimeoutOption := redis.DialReadTimeout(5 * time.Second)
	writeTimeoutOption := redis.DialWriteTimeout(5 * time.Second)
	password := ""

	Rpool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(fmt.Sprintf("redis://%s@%s:%d", password, "0.0.0.0", 6379),
				dialConnectTimeoutOption, readTimeoutOption, writeTimeoutOption)
			if err != nil {
				log.Panic("RedisDialURLError ", err)
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				log.Panic("RedisPINGError ", err)
				return err
			}
			return nil
		},
		MaxIdle:         10,
		MaxActive:       20,
		IdleTimeout:     10 * time.Second,
		Wait:            true,
		MaxConnLifetime: 0,
	}

	return &Cache{
		pool: Rpool,
	}
}

func (c *Cache) Ping() (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	_, err = conn.Do("PING")

	return
}

func (c *Cache) Get(key string) (val []byte, err error) {
	conn := c.pool.Get()
	defer conn.Close()

	val, err = redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return val, err
	}
	return val, nil
}

func (c *Cache) GetAllKeys(buyerId string) (keys []string, err error) {
	conn := c.pool.Get()
	defer conn.Close()

	keys, err = redis.Strings(conn.Do("KEYS", fmt.Sprintf("buyer_%s*", buyerId)))
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (c *Cache) Write(key string, data interface{}, ttl time.Duration) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, int64(ttl.Seconds()), data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) DeleteRecord(key string) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	_, err = conn.Do("DEL", key)
	return err
}

func (c *Cache) FlushByBuyerId(buyerId string) (err error) {
	conn := c.pool.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", fmt.Sprintf("buyer_%s*", buyerId)))
	if err != nil {
		return err
	}

	for _, k := range keys {
		_, err = conn.Do("DEL", k)
		if err != nil {
			return err
		}
	}

	return
}
