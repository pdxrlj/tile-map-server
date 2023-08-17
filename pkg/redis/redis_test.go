package redis

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var redisClient *Redis

func TestMain(m *testing.M) {
	fmt.Printf("TestMain\n")
	redis, err := NewRedis(
		WithRedisPrefix("test"),
		WithRedisAddr("192.168.1.234:6379"),
		WithRedisPassword("123"),
		WithRedisUsername(""),
	)
	if err != nil {
		panic(err)
	}
	redisClient = redis
	exitCode := m.Run()

	err = redisClient.Close()
	if err != nil {
		panic(err)
	}

	// Exit
	os.Exit(exitCode)
}

func TestRedis_Set(t *testing.T) {
	fmt.Printf("TestRedis_Set\n")
	err := redisClient.Set(context.Background(), "test", "test", 0).Err()
	if err != nil {
		assert.NoError(t, err)
	}
}

func TestRedis_Get(t *testing.T) {
	fmt.Printf("TestRedis_Get\n")
	s, err := redisClient.Get(context.Background(), "test").Result()
	if err != nil {
		assert.NoError(t, err)
	}
	fmt.Printf("s:%+v\n", s)
}

func TestRedis_Keys(t *testing.T) {
	fmt.Printf("TestRedis_Keys\n")
	s, err := redisClient.Keys(context.Background(), "*").Result()
	if err != nil {
		assert.NoError(t, err)
	}
	fmt.Printf("s:%+v\n", s)
}

func TestRedis_Exists(t *testing.T) {
	fmt.Printf("TestRedis_Exists\n")
	s, err := redisClient.Exists(context.Background(), "test").Result()
	assert.NoError(t, err)

	fmt.Printf("exists:%+v\n", s)
}

func TestRedis_Del(t *testing.T) {
	fmt.Printf("TestRedis_Del\n")
	s, err := redisClient.Del(context.Background(), "test").Result()
	assert.NoError(t, err)

	fmt.Printf("del:%+v\n", s)
}

func TestRedis_HSet(t *testing.T) {
	fmt.Printf("TestRedis_HSet\n")
	err := redisClient.HSet(context.Background(), "test", "test", "test").Err()
	assert.NoError(t, err)
}

func TestRedis_HGet(t *testing.T) {
	fmt.Printf("TestRedis_HGet\n")
	s, err := redisClient.HGet(context.Background(), "test", "test").Result()
	assert.NoError(t, err)

	fmt.Printf("s:%+v\n", s)
}

func TestRedis_HKeys(t *testing.T) {
	fmt.Printf("TestRedis_HKeys\n")
	s, err := redisClient.HKeys(context.Background(), "test").Result()
	assert.NoError(t, err)

	fmt.Printf("HKeys:%+v\n", s)
}

func TestRedis_HGetAll(t *testing.T) {
	fmt.Printf("TestRedis_HGetAll\n")
	s, err := redisClient.HGetAll(context.Background(), "test").Result()
	assert.NoError(t, err)

	fmt.Printf("GetAll:%+v\n", s)
}

func TestRedis_HDel(t *testing.T) {
	fmt.Printf("TestRedis_HDel\n")
	s, err := redisClient.HDel(context.Background(), "test", "test").Result()
	assert.NoError(t, err)

	fmt.Printf("Del:%+v\n", s)
}

func TestRedis_HExists(t *testing.T) {
	fmt.Printf("TestRedis_HExists\n")
	s, err := redisClient.HExists(context.Background(), "test", "test").Result()
	assert.NoError(t, err)

	fmt.Printf("Exists:%+v\n", s)
}

func TestRedis_HLen(t *testing.T) {
	fmt.Printf("TestRedis_HLen\n")
	s, err := redisClient.HLen(context.Background(), "test").Result()
	assert.NoError(t, err)

	fmt.Printf("Len:%+v\n", s)
}

func TestRedis_HMSet(t *testing.T) {
	fmt.Printf("TestRedis_HMSet\n")
	err := redisClient.HMSet(context.Background(), "test", map[string]interface{}{"test": "test"}).Err()
	assert.NoError(t, err)
}

func TestRedis_HMGet(t *testing.T) {
	fmt.Printf("TestRedis_HMGet\n")
	s, err := redisClient.HMGet(context.Background(), "test", "test").Result()
	assert.NoError(t, err)

	fmt.Printf("HMGet:%+v\n", s)
}
