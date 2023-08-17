package cmd

import (
	"context"
	"log/slog"
	"math"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"tile-map-server/configs"
	"tile-map-server/internal/tile_app"
	"tile-map-server/pkg/log"
	"tile-map-server/pkg/opentracing"
	"tile-map-server/pkg/redis"
)

var root = cobra.Command{
	Use:   "tile",
	Long:  "tile is a command line tool for generating map tiles from a variety of sources. ",
	Short: "tile is a command line tool for generating map tiles from a variety of sources. ",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := config.UnmarshalToConfig(&config.C)
		if err != nil {
			return err
		}

		// 初始化日志
		log.Init(config.GetAppConfigLogLevel())

		// redis 初始化
		rdm, err := redis.NewRedis(
			redis.WithRedisAddr(config.GetRedisAddr()),
			redis.WithRedisDB(config.GetRedisDB()),
			redis.WithRedisPassword(config.GetRedisPassword()),
			redis.WithRedisUsername(config.GetRedisUsername()),
		)
		if err != nil {
			return err
		}

		// 初始化opentracing
		slog.Info("opentracing_collection_endpoint", "endpoint", config.GetAppConfigOpentracingCollectionEndpoint())
		tracing := opentracing.Init("tile-map-server-v2", config.GetAppConfigOpentracingCollectionEndpoint()).
			Start(context.Background())
		defer tracing.Close()
		defer tracing.CloseRootSpan()

		// 初始化瓦片服务
		tileConfig := config.NewTileConfig(tracing.GetSpanCtx(), config.C, rdm)

		tile_app.NewTileAppServer(tileConfig).Run()

		return nil
	},
}

func Exec() error {
	err := root.Execute()
	if err != nil {
		return err
	}
	return nil
}

func init() {
	configPath, err := config.GetConfigPath()
	if err != nil {
		panic(err)
	}

	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err = viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// app command line flags
	CommandLine()

	// app command bind flags alias
	config.ViperBindFlagsAlias(root)
}

func CommandLine() {
	root.PersistentFlags().StringSlice("with-basemaps", []string{}, "注册一个地图服务使用数组的形式[layer,tile_format,storage_type,storage_way,path]")

	root.PersistentFlags().Bool("with-basemap", false, "是否默认开启底图服务")

	root.PersistentFlags().StringP("internet_access_address", "d", "http://127.0.0.1:8080", "注册瓦片可以访问的外网地址")

	// 水印
	root.PersistentFlags().String("with-water-text", "航天宏图", "设置水印文字")
	root.PersistentFlags().Int("with-water-density", 4, "设置水印密度")
	root.PersistentFlags().Bool("with-start-water", false, "是否开启水印")
	root.PersistentFlags().Int("with-start-water-level", 10, "设置水印开始级别")
	root.PersistentFlags().Int("with-water-font-size", 30, "设置水印文字大小")
	root.PersistentFlags().String("with-water-font-color", "#757A81", "设置水印文字的颜色")
	root.PersistentFlags().String("with-water-font-rotate", "0", "设置水印文字旋转的角度")
	root.PersistentFlags().String("water-font-merge-percent", "70", "设置水印与图片的融合度")
	root.PersistentFlags().Bool("with-water-switch-img", true, "设置水印使用图片")

	// App配置
	root.PersistentFlags().StringP("config_path", "c", "", "瓦片下载读取配置文件路径")
	root.PersistentFlags().StringP("log_level", "t", "info", "瓦片下载读取配置文件类型")
	root.PersistentFlags().Int("cache_size", math.MaxInt, "设置缓存大小")
	root.PersistentFlags().IntP("port", "p", 8080, "设置服务端口")
	root.PersistentFlags().Bool("access-ctrl", false, "是否开启用户使用认证")
	root.PersistentFlags().String("access-ctrl-url", "", "验证用户是否以及授权使用")
	root.PersistentFlags().Duration("access-ctrl-effective-time", time.Hour*13, "使用软件认证的有效时间")
	root.PersistentFlags().Bool("show-spend-time", false, "是否显示请求耗时")
	root.PersistentFlags().String("opentracing_collection_endpoint", "http://192.168.1.234:14268/api/traces", "opentracing 采集地址")

	// redis
	root.PersistentFlags().String("redis_host", "127.0.0.1", "redis的Host")
	root.PersistentFlags().String("redis_port", "6379", "redis的端口")
	root.PersistentFlags().String("redis_password", "", "redis的密码")
	root.PersistentFlags().String("redis_username", "", "redis的账号密码")
	root.PersistentFlags().Int("redis_db", 0, "redis的db")
	root.PersistentFlags().String("redis_prefix_key", "engine:tile-map:", "redis 存储的前缀key值")
}
