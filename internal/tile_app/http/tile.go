package http

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	config "tile-map-server/configs"
	"tile-map-server/internal/tile_app/repository"
)

type TileRouter struct {
	*echo.Echo
	TileConfig      *config.TileConfig
	redisRepository *repository.RedisRepository
}

func NewTileRouter(router *echo.Echo, config *config.TileConfig) *TileRouter {
	return &TileRouter{
		Echo:            router,
		TileConfig:      config,
		redisRepository: repository.NewRedisRepository(config.RDM),
	}
}

func (t *TileRouter) Register() {
	group := t.Group("/tile-server/v1")
	group.GET("/tile/:z/:x/:y", t.GetTile)
}

func (t *TileRouter) GetTile(ctx echo.Context) error {
	childSpan, _ := opentracing.StartSpanFromContext(ctx.Request().Context(), "GetTile")
	defer func() {
		childSpan.Finish()
	}()

	time.Sleep(1 * time.Second)
	return nil
}
