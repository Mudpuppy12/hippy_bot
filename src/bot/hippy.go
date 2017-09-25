package main

import (
	"log"

	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/chucknorris"
	// Import all the commands you wish to use

	"github.com/spf13/viper"
)

var (
	SLACK_API string
)

func init() {
	viper.SetConfigName("config") // no need to include file extension
	viper.AddConfigPath("/home/dennis/GoProjects/hippy_bot/src/bot")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Fatal(err)
	}

	SLACK_API = viper.GetString("bot.SLACK_API")
}

func main() {

	slack.Run(SLACK_API)
}
