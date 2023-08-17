package request

import (
	"strings"

	"github.com/labstack/echo/v4"

	"tile-map-server/internal/tile_app/repository"
)

func ParseRequest(r echo.Context, redisRepo *repository.RedisRepository) Tile {
	query := r.Request().URL.Query().Encode()
	queryArray := strings.Split(query, "&")
	queryMap := make(map[string]string, len(queryArray))
	for i := range queryArray {
		key, val, found := strings.Cut(queryArray[i], "=")
		if found {
			queryMap[strings.ToLower(key)] = val
		}
	}
	tile := Tile{
		Layer:      queryMap["layer"],
		TileMatrix: queryMap["tilematrix"],
		TileCol:    queryMap["tilecol"],
		TileRow:    queryMap["tilerow"],
		Timestamp:  queryMap["timestamp"],
	}

	if tile.Layer == "" {
		tile.Layer = r.Param("key")
	}

	if queryMap["objectid"] != "" {
		tile.Layer = queryMap["objectid"]
	}

	if queryMap["objectid"] != "" && tile.Timestamp == "" {
		tile.Timestamp, _ = redisRepo.GetMultipleTensesTileMaxOrMinTime(tile.Layer)
	}

	if tile.TileMatrix == "" {
		tile.TileMatrix = queryMap["z"]
	}

	if tile.TileCol == "" {
		tile.TileCol = queryMap["x"]
	}

	if tile.TileRow == "" {
		tile.TileRow = queryMap["y"]
	}

	if tile.TileMatrix == "" {
		tile.TileMatrix = r.Param("z")
	}

	if tile.TileCol == "" {
		tile.TileCol = r.Param("x")
	}

	if tile.TileRow == "" {
		tile.TileRow = r.Param("y")
	}
	if tile.Path == "" {
		tile.Path = r.Param("path")
		pathSplit := strings.Split(strings.TrimPrefix(tile.Path, "/"), "/")
		if len(pathSplit) == 3 {
			if tile.TileMatrix == "" {
				tile.TileMatrix = pathSplit[0]
			}

			if tile.TileCol == "" {
				tile.TileCol = pathSplit[1]
			}

			if tile.TileRow == "" {
				tile.TileRow = pathSplit[2]
			}
		}
	}

	if strings.Contains(tile.TileRow, ".") {
		tile.TileRow = tile.TileRow[:strings.Index(tile.TileRow, ".")]
	}

	return tile
}
