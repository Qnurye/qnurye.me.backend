package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func init() {
	var err error
	config = viper.New()
	config.SetConfigType("toml")
	config.SetConfigName("app")
	config.AddConfigPath("../conf/")
	config.AddConfigPath("conf/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file: ", err)
	}
}

func Get() *viper.Viper {
	return config
}

func Set(Key string, value any) {
	config.Set(Key, value)
	err := config.WriteConfig()
	if err != nil {
		log.Fatalln("Error on writing into config file: ", err)
	}
}

func GetFilePath() string {
	return config.ConfigFileUsed()
}
