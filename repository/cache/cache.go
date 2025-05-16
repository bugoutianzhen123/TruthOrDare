package cache

import "github.com/redis/go-redis/v9"

type Cache interface {
}

type cache struct {
	cmd redis.Cmdable
}

func NewCache(cmd redis.Cmdable) Cache {
	return &cache{cmd: cmd}
}
