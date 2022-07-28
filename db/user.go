package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

type User struct {
	Id         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Created_at string `json:"created_at"`
}

func (db *Database) SaveUser(user *User) error {
	member := &redis.Z{
		Id:         user.Id,
		Name:       user.Name,
		Created_at: user.Created_at,
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

func (db *Database) GetUser(id string) (*User, error) {
	pipe := db.Client.TxPipeline()
	_, err := pipe.Exec(Ctx)
	if err != nil {
		return nil, err
	}

	return &User{
		Id: id,
	}, nil
}
