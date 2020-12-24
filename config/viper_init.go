package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigFile("./config.toml")// name of config file (without extension)
	viper.AddConfigPath(".")             // optionally look for config in the working directory
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil {                      // Handle errors reading the config file
		logrus.Fatalf("Fatal error config file: %s \n", err)
	}

	addr := viper.GetString("sshw.addr")
	user := viper.GetString("sshw.user")

	logrus.Info("[config]", addr, user)
}
