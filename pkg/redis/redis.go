package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"chatcser/pkg/utils/syncx"

	red "github.com/redis/go-redis/v9"
	// "github.com/zeromicro/go-zero/core/errorx"
	// "github.com/zeromicro/go-zero/core/logx"
	// "github.com/zeromicro/go-zero/core/mapping"
	// "github.com/zeromicro/go-zero/core/syncx"
)

const (
	// ClusterType means redis cluster.
	ClusterType = "cluster"
	// NodeType means redis node.
	NodeType = "node"
	// Nil is an alias of redis.Nil.
	Nil = red.Nil

	blockingQueryTimeout = 5 * time.Second
	readWriteTimeout     = 2 * time.Second
	defaultSlowThreshold = time.Millisecond * 100
	defaultPingTimeout   = time.Second
)

var (
	// ErrNilNode is an error that indicates a nil redis node.
	ErrNilNode    = errors.New("nil redis node")
	slowThreshold = syncx.ForAtomicDuration(defaultSlowThreshold)
)

type (
	// Option defines the method to customize a Redis.
	Option func(r *Redis)

	// A Pair is a key/pair set used in redis zset.
	Pair struct {
		Key   string
		Score int64
	}

	// A FloatPair is a key/pair for float set used in redis zet.
	FloatPair struct {
		Key   string
		Score float64
	}

	// Redis defines a redis node/cluster. It is thread-safe.
	Redis struct {
		Addr string
		Type string
		Pass string
		tls  bool
		//brk   breaker.Breaker
		hooks []red.Hook
	}

	// RedisNode interface represents a redis node.
	RedisNode interface {
		red.Cmdable
	}

	// GeoLocation is used with GeoAdd to add geospatial location.
	GeoLocation = red.GeoLocation
	// GeoRadiusQuery is used with GeoRadius to query geospatial index.
	GeoRadiusQuery = red.GeoRadiusQuery
	// GeoPos is used to represent a geo position.
	GeoPos = red.GeoPos

	// Pipeliner is an alias of redis.Pipeliner.
	Pipeliner = red.Pipeliner

	// Z represents sorted set member.
	Z = red.Z
	// ZStore is an alias of redis.ZStore.
	ZStore = red.ZStore

	// IntCmd is an alias of redis.IntCmd.
	IntCmd = red.IntCmd
	// FloatCmd is an alias of redis.FloatCmd.
	FloatCmd = red.FloatCmd
	// StringCmd is an alias of redis.StringCmd.
	StringCmd = red.StringCmd
	// Script is an alias of redis.Script.
	Script = red.Script
)

// MustNewRedis returns a Redis with given options.
func MustNewRedis(conf RedisConf, opts ...Option) *Redis {
	rds, err := NewRedis(conf, opts...)
	//logx.Must(err)
	print(err)
	return rds
}

// New returns a Redis with given options.
// Deprecated: use MustNewRedis or NewRedis instead.
func New(addr string, opts ...Option) *Redis {
	return newRedis(addr, opts...)
}

// NewRedis returns a Redis with given options.
func NewRedis(conf RedisConf, opts ...Option) (*Redis, error) {
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	rds := newRedis(conf.Host, opts...)
	// if !conf.NonBlock {
	// 	if err := rds.checkConnection(conf.PingTimeout); err != nil {
	// 		return nil, errors.New("redis connect error") //errorx.Wrap(err, fmt.Sprintf("redis connect error, addr: %s", conf.Host))
	// 	}
	// }

	return rds, nil
}

func newRedis(addr string, opts ...Option) *Redis {
	r := &Redis{
		Addr: addr,
		Type: NodeType,
		//brk:  breaker.NewBreaker(),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

// NewScript returns a new Script instance.
func NewScript(script string) *Script {
	return red.NewScript(script)
}

// BitCountCtx is redis bitcount command implementation.
// func (s *Redis) BitCountCtx(ctx context.Context, key string, start, end int64) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.BitCount(ctx, key, &red.BitCount{
// 			Start: start,
// 			End:   end,
// 		}).Result()
// 		return err
// 	}, acceptable)

// 	return
// }

// BitOpAndCtx is redis bit operation (and) command implementation.
// func (s *Redis) BitOpAndCtx(ctx context.Context, destKey string, keys ...string) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.BitOpAnd(ctx, destKey, keys...).Result()
// 		return err
// 	}, acceptable)

// 	return
// }

// BitOpNotCtx is redis bit operation (not) command implementation.
// func (s *Redis) BitOpNotCtx(ctx context.Context, destKey, key string) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.BitOpNot(ctx, destKey, key).Result()
// 		return err
// 	}, acceptable)

// 	return
// }

// BitOpOr is redis bit operation (or) command implementation.
// func (s *Redis) BitOpOr(destKey string, keys ...string) (int64, error) {
// 	return s.BitOpOrCtx(context.Background(), destKey, keys...)
// }

// BitOpOrCtx is redis bit operation (or) command implementation.
// func (s *Redis) BitOpOrCtx(ctx context.Context, destKey string, keys ...string) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.BitOpOr(ctx, destKey, keys...).Result()
// 		return err
// 	}, acceptable)

// 	return
// }

// BitOpXorCtx is redis bit operation (xor) command implementation.
// func (s *Redis) BitOpXorCtx(ctx context.Context, destKey string, keys ...string) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.BitOpXor(ctx, destKey, keys...).Result()
// 		return err
// 	}, acceptable)

// 	return
// }

// BitPos is redis bitpos command implementation.
// func (s *Redis) BitPos(key string, bit, start, end int64) (int64, error) {
// 	return s.BitPosCtx(context.Background(), key, bit, start, end)
// }

// BitPosCtx is redis bitpos command implementation.
// func (s *Redis) BitPosCtx(ctx context.Context, key string, bit, start, end int64) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.BitPos(ctx, key, bit, start, end).Result()
// 		return err
// 	}, acceptable)

// 	return
// }

// Blpop uses passed in redis connection to execute blocking queries.
// Doesn't benefit from pooling redis connections of blocking queries
func (s *Redis) Blpop(node RedisNode, key string) (string, error) {
	return s.BlpopCtx(context.Background(), node, key)
}

// BlpopCtx uses passed in redis connection to execute blocking queries.
// Doesn't benefit from pooling redis connections of blocking queries
func (s *Redis) BlpopCtx(ctx context.Context, node RedisNode, key string) (string, error) {
	return s.BlpopWithTimeoutCtx(ctx, node, blockingQueryTimeout, key)
}

// BlpopEx uses passed in redis connection to execute blpop command.
// The difference against Blpop is that this method returns a bool to indicate success.
func (s *Redis) BlpopEx(node RedisNode, key string) (string, bool, error) {
	return s.BlpopExCtx(context.Background(), node, key)
}

// BlpopExCtx uses passed in redis connection to execute blpop command.
// The difference against Blpop is that this method returns a bool to indicate success.
func (s *Redis) BlpopExCtx(ctx context.Context, node RedisNode, key string) (string, bool, error) {
	if node == nil {
		return "", false, ErrNilNode
	}

	vals, err := node.BLPop(ctx, blockingQueryTimeout, key).Result()
	if err != nil {
		return "", false, err
	}

	if len(vals) < 2 {
		return "", false, fmt.Errorf("no value on key: %s", key)
	}

	return vals[1], true, nil
}

// BlpopWithTimeout uses passed in redis connection to execute blpop command.
// Control blocking query timeout
func (s *Redis) BlpopWithTimeout(node RedisNode, timeout time.Duration, key string) (string, error) {
	return s.BlpopWithTimeoutCtx(context.Background(), node, timeout, key)
}

// BlpopWithTimeoutCtx uses passed in redis connection to execute blpop command.
// Control blocking query timeout
func (s *Redis) BlpopWithTimeoutCtx(ctx context.Context, node RedisNode, timeout time.Duration,
	key string) (string, error) {
	if node == nil {
		return "", ErrNilNode
	}

	vals, err := node.BLPop(ctx, timeout, key).Result()
	if err != nil {
		return "", err
	}

	if len(vals) < 2 {
		return "", fmt.Errorf("no value on key: %s", key)
	}

	return vals[1], nil
}

// Decr is the implementation of redis decr command.
// func (s *Redis) Decr(key string) (int64, error) {
// 	return s.DecrCtx(context.Background(), key)
// }

// DecrCtx is the implementation of redis decr command.
// func (s *Redis) DecrCtx(ctx context.Context, key string) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.Decr(ctx, key).Result()
// 		return err
// 	}, acceptable)

// 	return
// }

// Decrby is the implementation of redis decrby command.
// func (s *Redis) Decrby(key string, decrement int64) (int64, error) {
// 	return s.DecrbyCtx(context.Background(), key, decrement)
// }

// DecrbyCtx is the implementation of redis decrby command.
// func (s *Redis) DecrbyCtx(ctx context.Context, key string, decrement int64) (val int64, err error) {
// 	err = s.brk.DoWithAcceptable(func() error {
// 		conn, err := getRedis(s)
// 		if err != nil {
// 			return err
// 		}

// 		val, err = conn.DecrBy(ctx, key, decrement).Result()
// 		return err
// 	}, acceptable)

// 	return
// }
