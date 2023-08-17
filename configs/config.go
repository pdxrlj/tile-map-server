package config

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"tile-map-server/pkg/redis"
)

var C *Config

type TileConfig struct {
	Ctx    context.Context
	Config *Config
	RDM    *redis.Redis
}

func NewTileConfig(ctx context.Context, config *Config, rdm *redis.Redis) *TileConfig {
	return &TileConfig{
		Ctx:    ctx,
		Config: config,
		RDM:    rdm,
	}
}

func UnmarshalToConfig(dst interface{}) error {
	err := viper.Unmarshal(dst)
	if err != nil {
		return err
	}
	return nil
}

func ViperBindFlagsAlias(command cobra.Command) {
	// redis alias
	_ = viper.BindPFlag("redis.host", command.PersistentFlags().Lookup("redis_host"))
	_ = viper.BindPFlag("redis.port", command.PersistentFlags().Lookup("redis_port"))
	_ = viper.BindPFlag("redis.password", command.PersistentFlags().Lookup("redis_password"))
	_ = viper.BindPFlag("redis.username", command.PersistentFlags().Lookup("redis_username"))
	_ = viper.BindPFlag("redis.db", command.PersistentFlags().Lookup("redis_db"))
	_ = viper.BindPFlag("redis.redis_prefix_key", command.PersistentFlags().Lookup("redis_prefix_key"))

	// 水印
	_ = viper.BindPFlag("watermark.start_water", command.PersistentFlags().Lookup("with-start-water"))
	_ = viper.BindPFlag("watermark.water_text", command.PersistentFlags().Lookup("with-water-text"))
	_ = viper.BindPFlag("watermark.water_density", command.PersistentFlags().Lookup("with-water-density"))
	_ = viper.BindPFlag("watermark.water_font_size", command.PersistentFlags().Lookup("with-water-font-size"))
	_ = viper.BindPFlag("watermark.water_level", command.PersistentFlags().Lookup("with-start-water-level"))
	_ = viper.BindPFlag("watermark.water_font_color", command.PersistentFlags().Lookup("with-water-font-color"))
	_ = viper.BindPFlag("watermark.water_font_rotate", command.PersistentFlags().Lookup("with-water-font-rotate"))
	_ = viper.BindPFlag("watermark.water_font_merge_percent", command.PersistentFlags().Lookup("water-font-merge-percent"))
	_ = viper.BindPFlag("watermark.switch_img", command.PersistentFlags().Lookup("with-water-switch-img"))

	// 应用配置
	_ = viper.BindPFlag("app_config.config_path", command.PersistentFlags().Lookup("config_path"))
	_ = viper.BindPFlag("app_config.log_level", command.PersistentFlags().Lookup("log_level"))
	_ = viper.BindPFlag("app_config.cache_size", command.PersistentFlags().Lookup("cache_size"))
	_ = viper.BindPFlag("app_config.port", command.PersistentFlags().Lookup("port"))
	_ = viper.BindPFlag("app_config.access-ctrl", command.PersistentFlags().Lookup("access-ctrl"))
	_ = viper.BindPFlag("app_config.access-ctrl-url", command.PersistentFlags().Lookup("access-ctrl-url"))
	_ = viper.BindPFlag("app_config.access-ctrl-effective-time", command.PersistentFlags().Lookup("access-ctrl-effective-time"))
	_ = viper.BindPFlag("app_config.show-spend-time", command.PersistentFlags().Lookup("show-spend-time"))
	_ = viper.BindPFlag("app_config.opentracing_collection_endpoint", command.PersistentFlags().Lookup("opentracing_collection_endpoint"))

	// basemap

	_ = viper.BindPFlag("basemap_config.with-basemaps", command.PersistentFlags().Lookup("with-basemaps"))
	_ = viper.BindPFlag("basemap_config.with-basemap", command.PersistentFlags().Lookup("with-basemap"))
	_ = viper.BindPFlag("basemap_config.internet_access_address", command.PersistentFlags().Lookup("internet_access_address"))
}
