package cfg

import (
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	PostgresDSN string
	Port        string
}

func LoadConfig(filePath string) (cfg Config, err error) {
	dir := filepath.Dir(filePath)
	file := filepath.Base(filePath)
	fileName := file[:len(file)-len(filepath.Ext(file))]

	viper.AddConfigPath(dir)
	viper.SetConfigName(fileName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&cfg)
	return
}
