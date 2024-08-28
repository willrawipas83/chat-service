package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Rawipass/chat-service/global_variable"
	"github.com/Rawipass/chat-service/logger"
	"github.com/spf13/viper"
)

var configFile string

func initTimezone() {
	loc, err := time.LoadLocation(global_variable.TimeZone)
	if err != nil {
		logger.Logger.Error("unable to set timezone cuz: %s", err)
		os.Exit(0)
	}
	time.Local = loc
	logger.Logger.Infof("Load Global Timezone to : %s", time.Local)
}

func initComponent() {
	// Init Global Variable
	global_variable.InitVariable()

	// Init Timezone
	initTimezone()
}

func initConfig() {
	// Init viper
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
	}
	// Read Config
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "unable to read config: %v\n", err)
		os.Exit(1)
	}
}
