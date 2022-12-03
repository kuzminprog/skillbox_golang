package repository

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type DataBase struct {
	rdb    *redis.Client
	lastId int
}

func NewDataBase(address string) (*DataBase, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	lastId := 0

	iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		id, err := strconv.Atoi(iter.Val())
		lastId = max(lastId, id)

		if err != nil {
			return nil, err
		}
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return &DataBase{rdb: rdb, lastId: lastId}, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
