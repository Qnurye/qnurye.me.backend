package config

import (
	"log"
	"time"
)
import (
	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	BaseUrl   string
}

var AppConfig = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}

type database struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var DatabaseConfig = &database{}

type redis struct {
	Host        string
	Port        int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisConfig = &redis{}

var cfg *ini.File

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("Failed to parse \"conf/app.ini\": %v", err)
	}

	mapTo("app", AppConfig)
	mapTo("server", ServerConfig)
	mapTo("database", DatabaseConfig)
	mapTo("redis", RedisConfig)

	// TODO: Complete the config func
}
