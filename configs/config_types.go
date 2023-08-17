package config

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

type RedisConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	DB              int           `mapstructure:"db"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Timeout         time.Duration `mapstructure:"timeout"` // 单位被转化纳秒
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	ConnMinIdle     int           `mapstructure:"conn_min_idle"`
	ConnMaxOpen     int           `mapstructure:"conn_max_open"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idletime"`
	RedisPrefixKey  string        `mapstructure:"redis_prefix_key"`
}

type S3Config struct {
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
	Region    string `mapstructure:"region"`
	Endpoint  string `mapstructure:"endpoint"`
}

type ObsConfig struct {
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
	Region    string `mapstructure:"region"`
	Endpoint  string `mapstructure:"endpoint"`
}

type MinioConfig struct {
	AccessKey string `mapstructure:"accessKey"`
	SecretKey string `mapstructure:"secretKey"`
	Region    string `mapstructure:"region"`
	Endpoint  string `mapstructure:"endpoint"`
}

type StorageKindConfig struct {
	S3    S3Config    `mapstructure:"s3"`
	OBS   ObsConfig   `mapstructure:"obs"`
	Minio MinioConfig `mapstructure:"minio"`
}

type EngineStudioConfig struct {
	ImageURL    string            `mapstructure:"imageUrl"`
	DatasetURL  string            `mapstructure:"datasetUrl"`
	StorageType string            `mapstructure:"storageType"`
	StorageKind StorageKindConfig `mapstructure:"storageKind"`
}

type WatermarkConfig struct {
	StartWater            bool   `mapstructure:"start_water"`
	WaterText             string `mapstructure:"water_text"`
	WaterLevel            int    `mapstructure:"water_level"`
	WaterDensity          int    `mapstructure:"water_density"`
	WaterFontSize         int    `mapstructure:"water_font_size"`
	WaterFontColor        string `mapstructure:"water_font_color"`
	WaterFontRotate       int    `mapstructure:"water_font_rotate"`
	WaterFontMergePercent int    `mapstructure:"water_font_merge_percent"`
	SwitchImg             bool   `mapstructure:"switch_img"`
}

type BasemapConfig struct {
	WithBasemaps          []string `mapstructure:"with-basemaps"`
	WithBasemap           bool     `mapstructure:"with-basemap"`
	InternetAccessAddress string   `mapstructure:"internet_access_address"`
}

type AppConfig struct {
	ConfigPath                    string        `mapstructure:"config_path"`
	LogLevel                      string        `mapstructure:"log_level"`
	CacheSize                     int           `mapstructure:"cache_size"`
	Port                          int           `mapstructure:"port"`
	AccessCtrl                    bool          `mapstructure:"access-ctrl"`
	AccessCtrlUrl                 string        `mapstructure:"access-ctrl-url"`
	AccessCtrlEffectiveTime       time.Duration `mapstructure:"access-ctrl-effective-time"`
	ShowSpendTime                 bool          `mapstructure:"show-spend-time"`
	OpentracingCollectionEndpoint string        `mapstructure:"opentracing_collection_endpoint"`
}

type Config struct {
	Redis           RedisConfig                   `mapstructure:"redis"`
	Tilers          map[string]EngineStudioConfig `mapstructure:"tilers"`
	WatermarkConfig WatermarkConfig               `mapstructure:"watermark"`

	BasemapConfig BasemapConfig `mapstructure:"basemap_config"`

	AppConfig AppConfig `mapstructure:"app_config"`
}

// Studio

func GetTilers() map[string]EngineStudioConfig {
	return C.Tilers
}

// BasemapConfig

func GetWithBasemaps() []string {
	return C.BasemapConfig.WithBasemaps
}

func GetWithBasemap() bool {
	return C.BasemapConfig.WithBasemap
}

func GetInternetAccessAddress() string {
	return C.BasemapConfig.InternetAccessAddress
}

// RedisConfig

func GetRedisHost() string {
	return C.Redis.Host
}

func GetRedisAddr() string {
	host := GetRedisHost()
	port := GetRedisPort()
	return fmt.Sprintf("%s:%d", host, port)
}

func GetRedisPort() int {
	return C.Redis.Port
}

func GetRedisDB() int {
	return C.Redis.DB
}

func GetRedisUsername() string {
	return C.Redis.Username
}

func GetRedisPassword() string {
	return C.Redis.Password
}

func GetRedisTimeout() time.Duration {
	return time.Duration(C.Redis.Timeout.Seconds())
}

func GetRedisReadTimeout() time.Duration {
	return time.Duration(C.Redis.ReadTimeout.Seconds())
}

func GetRedisWriteTimeout() time.Duration {
	return time.Duration(C.Redis.WriteTimeout.Seconds())
}

func GetRedisConnMinIdle() int {
	return C.Redis.ConnMinIdle
}

func GetRedisConnMaxOpen() int {
	return C.Redis.ConnMaxOpen
}

func GetRedisConnMaxLifetime() time.Duration {
	return time.Duration(C.Redis.ConnMaxLifetime.Seconds())
}

func GetRedisConnMaxIdleTime() time.Duration {
	return time.Duration(C.Redis.ConnMaxIdleTime.Seconds())
}

func GetRedisPrefixKey() string {
	return C.Redis.RedisPrefixKey
}

// AppConfig

func GetAppConfigPath() string {
	return C.AppConfig.ConfigPath
}

func GetAppConfigLogLevel() slog.Level {
	level := C.AppConfig.LogLevel
	var l slog.Level
	parse := func(s string) slog.Level {
		var err error
		defer func() {
			if err != nil {
				slog.Error("slog: level string %q: %w", s, err)
				return
			}
		}()

		name := s
		offset := 0
		if i := strings.IndexAny(s, "+-"); i >= 0 {
			name = s[:i]
			offset, err = strconv.Atoi(s[i:])
			if err != nil {
				return slog.LevelInfo
			}
		}
		switch strings.ToUpper(name) {
		case "DEBUG":
			l = slog.LevelDebug
		case "INFO":
			l = slog.LevelInfo
		case "WARN":
			l = slog.LevelWarn
		case "ERROR":
			l = slog.LevelError
		default:
			return slog.LevelInfo
		}
		l += slog.Level(offset)
		return l
	}

	slog.Info("log level", "level", level)

	return parse(level)
}

func GetAppConfigCacheSize() int {
	return C.AppConfig.CacheSize
}

func GetAppConfigPort() int {
	return C.AppConfig.Port
}

func GetAppConfigAccessCtrl() bool {
	return C.AppConfig.AccessCtrl
}

func GetAppConfigAccessCtrlUrl() string {
	return C.AppConfig.AccessCtrlUrl
}

func GetAppConfigAccessCtrlEffectiveTime() time.Duration {
	return time.Duration(C.AppConfig.AccessCtrlEffectiveTime.Seconds())
}

func GetAppConfigShowSpendTime() bool {
	return C.AppConfig.ShowSpendTime
}

func GetAppConfigOpentracingCollectionEndpoint() string {
	return C.AppConfig.OpentracingCollectionEndpoint
}
