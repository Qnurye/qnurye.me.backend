package config

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
	"time"
)

type AppConfig struct {
	App      app
	Server   server
	Database database
	Redis    redis
}

type app struct {
	JwtSecret string
	BaseUrl   string
}

type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type database struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type redis struct {
	Host        string
	Port        int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var (
	cfg  *AppConfig
	once sync.Once
)

func Config() *AppConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("./conf/app.toml")
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
