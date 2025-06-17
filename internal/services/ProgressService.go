package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"progress-tracker/internal/queries"
	"time"
)

type ProgressService struct {
	rdb *redis.Client
}

func NewProgressService() *ProgressService {
	serivce := &ProgressService{redis.NewClient(&redis.Options{
		Addr: viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		DB:   viper.GetInt("redis.db"),
	})}

	if _, err := serivce.rdb.Ping(context.TODO()).Result(); err != nil {
		log.Fatal("connection error redis:", err)
	}

	return serivce
}

func (p *ProgressService) SetProgress(query queries.SetProgressQuery) error {
	return p.rdb.Set(context.TODO(), query.JobID.String(), query.Progress, 0).Err()
}

func (p *ProgressService) GetProgress(jobID uuid.UUID) (float32, error) {
	value := p.rdb.Get(context.TODO(), jobID.String())

	return value.Float32()
}

func (p *ProgressService) StartQueueWorker() {
	go readFromQueue(p)
}

func readFromQueue(p *ProgressService) {
	for {
		select {
		case <-context.TODO().Done():

		default:
			result, err := p.rdb.BRPop(context.TODO(), 5*time.Second, "progress").Result()
			if err != nil {
				if !errors.Is(err, redis.Nil) {
					log.Printf("read from queue error: %v", err)
				}
				continue
			}

			message := result[1]
			var query queries.SetProgressQuery
			if err := json.Unmarshal([]byte(message), &query); err == nil {
				_ = p.SetProgress(query)
			}
		}
	}
}
