package repository

import (
	"Online-Shopping-Microservices/microservices/wallet-service/constant"
	"context"
	"github.com/RezaOptic/go-utils/logger"
	"github.com/go-redis/redis/v7"
)

// QueueRepoInterface
type QueueRepoInterface interface {
	PopUserFromQueue() (string, error)
	InsertForRealTime(user string) (int64, error)
	GetWinningUsers() ([]string, error)
}

// QueueRepo
type QueueRepo struct {
	Context context.Context
	Self    QueueRepoInterface
}

func NewQueueRepo(ctx context.Context) QueueRepoInterface {
	x := &QueueRepo{Context: ctx}
	x.Self = x
	return x

}

// PopUserFromQueue method for pop users from redis queue
func (q QueueRepo) PopUserFromQueue() (string, error) {

	result, err := DBS.Redis.LPop(constant.UsersKey).Result()

	if err != nil {
		if err == redis.Nil {
			logger.ZSLogger.Errorf("error on pop from user queue with error :%s", err)
			return "", nil

		}
		return "", err
	}
	return result, nil
}

//insert into hash for get real time user
func (q QueueRepo) InsertForRealTime(user string) (int64, error) {

	result, err := DBS.Redis.RPush(constant.RealTimeKey, user).Result()

	if err != nil {
		if err == redis.Nil {
			logger.ZSLogger.Errorf("error on push  to hash for monitoring with error :%s", err)
			return 0, nil
		}
		return 0, err
	}
	return result, nil
}

//insert into hash for get real time user
func (q QueueRepo) GetWinningUsers() ([]string, error) {
	var result []string

	result, err := DBS.Redis.LRange(constant.RealTimeKey, constant.WinningUserStart, constant.WinningUserEnd).Result()

	if err != nil {
		if err == redis.Nil {
			logger.ZSLogger.Errorf("error on push  to hash for winngin users with error :%s", err)
			return result, nil
		}
		return result, err
	}
	return result, nil
}
