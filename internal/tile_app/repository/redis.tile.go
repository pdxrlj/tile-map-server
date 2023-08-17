package repository

type TileInfo struct {
	Layer string `json:"layer,omitempty" redis:"Layer"`
	//注册的标识ID
	Identify string `json:"identify,omitempty" redis:"Identify"`

	//是否是时序瓦片
	TimeSequence string `json:"time_sequence" redis:"TimeSequence"`

	//瓦片的格式必填 png/jpeg/webp/pbf
	TileFormat string `json:"tile_format,omitempty" redis:"TileFormat"`

	// 创建瓦片的方式 web墨卡托，经纬度，自定义
	CreateTileType string `json:"create_tile_type,omitempty" redis:"CreateTileType"`
	// 存储的方式 s3 minio obs nfs file,arcgis-10.2 等
	StorageType string `json:"storage_type,omitempty" redis:"StorageType"`

	Crs string `json:"crs,omitempty" redis:"crs"`

	// 瓦片的存储格式 zxy zyx
	//StorageWay string `json:"storage_way,omitempty" redis:"StorageWay"`

	// nfs file 文件存储方式 /xxx/xxx/path/{z}/{x}/{y}.png

	// s3 minio obs 文件存储方式 object://ak:sk@endpoint/region/bucket/path/{z}/{x}/{y}.png

	// mbtiles 文件存储方式 db://user:password@host:port/database/collection/{z}/{x}/{y}

	Path string `json:"path,omitempty" redis:"Path"`

	//SecretKey string `json:"secret_key,omitempty" redis:"SecretKey"`
	//AccessKey string `json:"access_key,omitempty" redis:"AccessKey"`
	//Region    string `json:"region,omitempty" redis:"Region"`
	//Endpoint  string `json:"endpoint,omitempty" redis:"Endpoint"`
	//Bucket    string `json:"bucket,omitempty" redis:"Bucket"`
	//Prefix    string `json:"prefix,omitempty" redis:"Prefix"`
	//
	//Host       string `json:"host,omitempty" redis:"Host"`
	//Port       string `json:"port,omitempty" redis:"Port"`
	//User       string `json:"user,omitempty" redis:"User"`
	//Password   string `json:"password,omitempty" redis:"Password"`
	//Database   string `json:"database,omitempty" redis:"Database"`
	//Collection string `json:"collection,omitempty" redis:"Collection"`

	WaterLevel        string `json:"water_level,omitempty" redis:"WaterLevel"`
	WaterDensity      string `json:"water_density,omitempty" redis:"WaterDensity"`
	WaterText         string `json:"water_text,omitempty" redis:"WaterText"`
	WaterFontSize     string `json:"water_font_size,omitempty" redis:"WaterFontSize"`
	WaterFontColor    string `json:"water_font_color,omitempty" redis:"WaterFontColor"`
	WaterFontRotate   string `json:"water_font_rotate,omitempty" redis:"WaterFontRotate"`
	WaterMergePercent string `json:"water_merge_percent,omitempty" redis:"WaterMergePercent"`
	WaterSwitchImg    string `json:"water_switch_img,omitempty" redis:"WaterSwitchImg"`
}
