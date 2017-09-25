package main

import (
	"github.com/go-chat-bot/bot/irc"
	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/chucknorris"
	// Import all the commands you wish to use
	"os"
	"strings"
)

func main() {

	go irc.Run(&irc.Config{
		Server:   "chat.freenode.net:6697",
		Channels: strings.Split("#dogpile", ","),
		User:     "Moonie-bot",
		Nick:     "Moonie-bot",
		Password: "",
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != ""})

	slack.Run("xoxb-233695181313-vkKL3ICGSxZghe6niuyhPr7l")
}
