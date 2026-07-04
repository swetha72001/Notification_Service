package internal

import (
	"Backend/config"
	"Backend/service/sc"
	"context"
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(ctx context.Context, deps *sc.Dependencies) error {
	viper.SetConfigFile("config.toml")
	viper.ReadInConfig()

	nsConfig := &config.NSConfig{}

	err := viper.Unmarshal(nsConfig)
	if err != nil {
		log.Println("error while unmarshal config")
		return err
	}
	deps.NsConfig = nsConfig

	return nil
}
