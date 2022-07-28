package db

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

type User struct {
	Id         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Created_at int    `json:"created_at"`
}

func (db *Database) SaveUser(user *User) error {
	member := &redis.Z{
		Member: user.Id,
	}
	pipe := db.Client.TxPipeline()
	pipe.ZAdd(Ctx, "user", member)
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return err
	}
	fmt.Println()
	return nil
}
