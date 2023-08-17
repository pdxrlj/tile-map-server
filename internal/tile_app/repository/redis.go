package repository

import (
	"context"

	"github.com/golang-module/carbon/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"

	"tile-map-server/pkg/redis"
)

type RedisRepository struct {
	rdm *redis.Redis
}

func NewRedisRepository(rdm *redis.Redis) *RedisRepository {
	return &RedisRepository{
		rdm: rdm,
	}
}

// GetMultipleTensesTileMaxOrMinTime
// @Summary 获取多个时态的最大或最小时间
func (r *RedisRepository) GetMultipleTensesTileMaxOrMinTime(layer string, operate ...string) (string, error) {
	defaultOperate := ">"
	if len(operate) != 0 {
		defaultOperate = operate[0]
	}
	result, err := r.rdm.LRange(context.Background(), redis.KeyGlue("multiple", layer), 0, -1).
		Result()
	if err != nil {
		return "", err
	}

	if len(result) == 0 {
		return "", errors.New("compare time data is empty")
	}

	var defaultTime = carbon.Parse(result[0], carbon.PRC)
	for _, item := range result {
		parse := carbon.Parse(item, carbon.PRC)
		if compare := parse.Compare(defaultOperate, defaultTime); compare {
			defaultTime = parse
		}
	}

	return defaultTime.Format("YmdHis", carbon.PRC), nil
}

// GetTileInfo
// @Summary 获取瓦片信息
func (r *RedisRepository) GetTileInfo(ctx context.Context, key string) (*TileInfo, error) {
	tileInfo := (*TileInfo)(nil)
	if err := r.rdm.HGetAll(ctx, key).Scan(&tileInfo); err != nil {
		span, _ := opentracing.StartSpanFromContext(ctx, "GetTileInfo")
		defer span.Finish()
		span.LogFields(log.String("error", err.Error()))
		return nil, err
	}
	return tileInfo, nil
}
