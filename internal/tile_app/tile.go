package tile_app

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"tile-map-server/configs"
	"tile-map-server/internal/tile_app/http"
	"tile-map-server/pkg/logger"
)

type TileAppServer struct {
	TileConfig *config.TileConfig
	echo       *echo.Echo
}

func NewTileAppServer(config *config.TileConfig) *TileAppServer {
	return &TileAppServer{
		TileConfig: config,
		echo:       echo.New(),
	}
}

func (t *TileAppServer) Run() {
	t.echo.Logger = logger.NewLogger()
	http.NewTileRouter(t.echo, t.TileConfig).Register()

	t.echo.Logger.Fatal(t.echo.Start(":" + strconv.Itoa(config.GetAppConfigPort())))
}
