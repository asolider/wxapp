package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/parnurzeal/gorequest"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               "127.0.0.1:6379",
		Password:           "",
		DB:                 0,
		MaxRetries:         2,
		PoolSize:           5,
		IdleTimeout:        60 * time.Second,
		IdleCheckFrequency: 5 * time.Second,
	})
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

const (
	ACCESS_TOKEN_CACHE_KEY string = "access_token_cache_key"
)

func GetAccessToken() string {
	//redisClient.Set(ACCESS_TOKEN_CACHE_KEY, "test", 10000*time.Second)
	cache, _ := redisClient.Get(ACCESS_TOKEN_CACHE_KEY).Result()
	if cache != "" {
		return cache
	}

	token := getTokenFromRemote()
	redisClient.Set(ACCESS_TOKEN_CACHE_KEY, token.AccessToken, time.Duration(token.ExpiresIn-3600)*time.Second)
	return token.AccessToken
}

func getTokenFromRemote() AccessToken {
	// https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
	remoteUrl := "https://api.weixin.qq.com/cgi-bin/token"
	queryParms := map[string]string{
		"grant_type": "client_credential",
		"appid":      APPID,
		"secret":     APPSECRET,
	}
	_, body, err := gorequest.New().Get(remoteUrl).Query(queryParms).EndBytes()
	res := AccessToken{}
	if err != nil {
		log.Printf("获取access_token 出错: %s", err)
		return res
	}
	log.Printf("从服务器获取token")
	json.Unmarshal(body, &res)
	return res
}
