package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/webook-project-go/webook-user/domain"
	"time"
)

type Cache interface {
	Set(ctx context.Context, du domain.User) error
	Get(ctx context.Context, id int64) (domain.User, error)
}

func NewUserCache(client redis.Cmdable) Cache {
	return newRedisUserCache(client, time.Minute*15)
}

var (
	ErrKeyNotFound = redis.Nil
)

type cache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func newRedisUserCache(client redis.Cmdable, expiration time.Duration) Cache {
	return &cache{
		cmd:        client,
		expiration: expiration,
	}
}

func (cache *cache) Get(ctx context.Context, id int64) (domain.User, error) {
	key := cache.key(id)
	var user domain.User
	val, err := cache.cmd.Get(ctx, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	err = json.Unmarshal(val, &user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (cache *cache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}

func (cache *cache) Set(ctx context.Context, du domain.User) error {
	key := cache.key(du.Id)
	val, err := json.Marshal(&du)
	if err != nil {
		// log
		return errors.New("json marshal failed")
	}
	err = cache.cmd.Set(ctx, key, val, cache.expiration).Err()
	if err != nil {
		// log
		return err
	}
	return nil
}
