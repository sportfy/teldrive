package config

import (
	"time"
)

type Config struct {
	Server ServerConfig
	Log    LoggingConfig
	JWT    JWTConfig
	DB     DBConfig
	TG     TGConfig
}

type ServerConfig struct {
	Port             int
	GracefulShutdown time.Duration
}

type TGConfig struct {
	AppId             int
	AppHash           string
	RateLimit         bool
	RateBurst         int
	Rate              int
	DeviceModel       string
	SystemVersion     string
	AppVersion        string
	LangCode          string
	SystemLangCode    string
	LangPack          string
	SessionFile       string
	BgBotsLimit       int
	DisableStreamBots bool
	Uploads           struct {
		EncryptionKey string
		Threads       int
		Retention     time.Duration
	}
}

type LoggingConfig struct {
	Level       int
	Development bool
	File        string
}

type JWTConfig struct {
	Secret       string
	SessionTime  time.Duration
	AllowedUsers []string
}

type DBConfig struct {
	DataSource string
	LogLevel   int
	Migrate    struct {
		Enable bool
	}
	Pool struct {
		MaxOpenConnections int
		MaxIdleConnections int
		MaxLifetime        time.Duration
	}
}
