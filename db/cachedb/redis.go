package cachedb

import (
	"net/url"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRedis(redisUri string) (*redis.Client, error) {

	uri, err := url.Parse(redisUri)

	if err != nil {
		return nil, err
	}

	usr := uri.User.Username()
	pwd, _ := uri.User.Password()
	dbNum, _ := strconv.Atoi(uri.Path[1:])

	opt := redis.Options{
		Username: usr,
		Password: pwd,
		DB:       dbNum,
	}

	client := redis.NewClient(&opt)
	return client, nil
}

var instance *redis.Client

func Get() *redis.Client {
	return instance
}
