package http

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"

	"tile-map-server/internal/pkg/request"
	"tile-map-server/internal/tile_app/repository"
)

type InterceptorFunc[T any] func(next InterceptorNextFunc[T]) InterceptorNextFunc[T]
type InterceptorNextFunc[T any] func(t T) error

func runInterceptors[T any](
	data T,
	next InterceptorNextFunc[T],
	interceptors ...InterceptorFunc[T],
) error {
	for i := len(interceptors) - 1; i >= 0; i-- {
		next = interceptors[i](next)
	}

	return next(data)
}

func ParseRequest(ctx context.Context, r echo.Context, repo *repository.RedisRepository) InterceptorFunc[request.Tile] {
	return func(next InterceptorNextFunc[request.Tile]) InterceptorNextFunc[request.Tile] {
		return func(t request.Tile) error {
			t = request.ParseRequest(r, repo)
			return next(t)
		}
	}
}

func TileStorageInfo(ctx context.Context, repo *repository.RedisRepository) InterceptorFunc[request.Tile] {
	return func(next InterceptorNextFunc[request.Tile]) InterceptorNextFunc[request.Tile] {
		return func(t request.Tile) error {
			key := t.Layer
			// 如果是时序瓦片，key=layer+"_"+identify
			if t.Timestamp != "" {
				key = t.Layer + "_" + t.Timestamp
			}

			// 获取存储的瓦片的信息
			info, err := repo.GetTileInfo(ctx, key)
			if err != nil {
				return err
			}

			if info == nil {
				return errors.New("not found tile info")
			}
			// TODO 如果是时序瓦片存的时候 path=filepath.Join(layerInfo.Path, request.Timestamp) 但是这里没有处理，需要处理
			t.SetStorageInfo(info)

			return next(t)
		}
	}
}

func BuildTileRoutingRules(ctx context.Context) InterceptorFunc[request.Tile] {
	return func(next InterceptorNextFunc[request.Tile]) InterceptorNextFunc[request.Tile] {
		return func(t request.Tile) error {

			return next(t)
		}
	}
}
