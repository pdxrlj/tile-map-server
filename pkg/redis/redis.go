package redis

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Options struct {
	*redis.Options
	prefix string
}

const (
	pingTimeout = time.Second * 5
)

type Option func(config *Options) error

type Redis struct {
	*redis.Client
}

func WithRedisUsername(username string) Option {
	return func(config *Options) error {
		config.Username = username
		return nil
	}
}

func WithRedisPassword(password string) Option {
	return func(config *Options) error {
		config.Password = password
		return nil
	}
}

func WithRedisDB(db int) Option {
	return func(config *Options) error {
		config.DB = db
		return nil
	}
}

func WithRedisAddr(addr string) Option {
	return func(config *Options) error {
		config.Addr = addr
		return nil
	}
}

func WithRedisOptions(options *redis.Options) Option {
	return func(config *Options) error {
		config.Options = options
		return nil
	}
}

func WithRedisPrefix(prefix string) Option {
	return func(config *Options) error {
		config.prefix = prefix
		return nil
	}
}

func NewRedis(options ...Option) (*Redis, error) {
	baseConfig := &Options{
		Options: &redis.Options{},
	}
	for _, opt := range options {
		if err := opt(baseConfig); err != nil {
			return nil, err
		}
	}

	client := redis.NewClient(baseConfig.Options)
	ctx, _ := context.WithTimeoutCause(context.Background(), pingTimeout, errors.New("redis ping timeout"))

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errors.Wrap(err, "redis ping timeout")
	}
	client.AddHook(prefixKeyHook{
		prefix: baseConfig.prefix,
	})

	return &Redis{
		Client: client,
	}, nil
}

func (r *Redis) GetRedisClient() (*redis.Client, error) {
	return r.Client, nil
}

var (
	_ redis.Hook = (*prefixKeyHook)(nil)
)

type prefixKeyHook struct {
	prefix string
}

func (p prefixKeyHook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

func (p prefixKeyHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		cmd.Args()[1] = p.prefix + cmd.Args()[1].(string)

		return next(ctx, cmd)
	}
}

func (p prefixKeyHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		for _, cmd := range cmds {
			cmd.Args()[1] = p.prefix + cmd.Args()[1].(string)
		}

		return next(ctx, cmds)
	}
}

func KeyGlue(strs ...string) string {
	s := strings.Builder{}
	l := len(strs)
	s.Grow(l)

	for i := range strs {
		s.WriteString(strs[i])
		if i != l-1 {
			s.WriteString(":")
		}
	}

	return s.String()
}
