package db

import (
	"github.com/go-redis/redis"
	"github.com/golang/glog"
)

type RedisDb struct {
	DB  *redis.Client
}

var redisDb *RedisDb

func init()  {
	redisDb = &RedisDb{}
}

func GetRedidDB() *RedisDb{
	return redisDb
}

func (this *RedisDb)Init(){
	redisdb := redis.NewClient(&redis.Options{
					Addr:     "localhost:6379",
					Password: "", // no password set
					DB:       0,  // use default DB
					})

	pong, err := redisdb.Ping().Result()
	if err!=nil{
		glog.Info("redis init error:",err.Error())
	}else{
		glog.Info("redis init success:",pong)
	}
	this.DB = redisdb
}

func (this *RedisDb)Del(key string){
	allkeys := this.DB.Keys(key)
	for _,key :=range(allkeys.Val()){
		ret:=this.DB.Del(key)
		glog.Info("redisdb ret:",ret)
	}
}