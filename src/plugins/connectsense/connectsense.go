package connectsense

import (
	"regexp"

	"github.com/go-chat-bot/bot"
)

const (
	pattern = "(?i)\\b(dogtemp)\\b"
)

var (
	re = regexp.MustCompile(pattern)
)

func tempStatus(command *bot.PassiveCmd) (string, error) {
	if re.MatchString(command.Raw) {
		return "tempStatus Return", nil
	}
	return "", nil
}

func init() {

	bot.RegisterPassiveCommand(
		"dogtemp",
		tempStatus)
}
