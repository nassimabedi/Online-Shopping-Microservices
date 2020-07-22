package repository

import (
	"arvan.ir/app-services/wallet-service/config"
	"fmt"
	"github.com/go-redis/redis/v7"
	mgo "gopkg.in/mgo.v2"
)

// dbs struct for managing database connections
type dbs struct {
	Redis   redis.UniversalClient
	MongoDB *mgo.Session
}

var DBS dbs

// Init function for init databases
func Init() {
	redisConnection()
	MongoConnection()
}

// redisConnection function for connecting to redis server
func redisConnection() {

	opt := redis.UniversalOptions{
		Addrs: config.Configs.Redis.Addresses,
	}
	DBS.Redis = redis.NewUniversalClient(&opt)

	result, err := DBS.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("redis connected with result :%s \n", result)
}

func MongoConnection() {
	session, err := mgo.Dial(config.Configs.MongoDB.Addresses)
	if err != nil {
		fmt.Println("session err:", err)
		panic(err)
	}
	DBS.MongoDB = session

}
