package request

import (
	"tile-map-server/internal/tile_app/repository"
	"tile-map-server/tools"
)

type Tile struct {
	Layer      string
	TileMatrix string
	TileCol    string
	TileRow    string

	Timestamp string
	Path      string
	Filename  string

	storageInfo *repository.TileInfo
}

func (t *Tile) GetIntTileCol() int {
	return tools.StringToInt(t.TileCol)
}

func (t *Tile) GetIntTileRow() int {
	return tools.StringToInt(t.TileRow)
}

func (t *Tile) GetIntTileMatrix() int {
	return tools.StringToInt(t.TileMatrix)
}

func (t *Tile) SetStorageInfo(info *repository.TileInfo) {
	t.storageInfo = info
}
