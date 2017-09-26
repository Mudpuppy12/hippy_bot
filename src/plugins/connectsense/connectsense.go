package connectsense

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/go-chat-bot/bot"
	"github.com/spf13/viper"
)

const (
	pattern = "(?i)\\b(dogtemp)\\b"
)

var (
	re       = regexp.MustCompile(pattern)
	url      string
	apiToken string
	deviceID string
)

type apiResponse struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func getData(body []byte) (*apiResponse, error) {
	var s = new(apiResponse)

	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	fahrenheit := 9.0/5.0*s.Temperature + 32
	s.Temperature = fahrenheit

	return s, err
}

func tempStatus(command *bot.PassiveCmd) (string, error) {
	if re.MatchString(command.Raw) {

		apiURL := url + "/" + apiToken + "/devices/" + deviceID

		response, err := http.Get(apiURL)

		if err != nil {
			return "Error occured", nil
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			return "Error occured", nil
		}

		s, err := getData([]byte(contents))

		return fmt.Sprintf("It is %.1f \u2109  with a %.2f%% humidity.", s.Temperature, s.Humidity), nil

	}
	return "", nil
}

func init() {

	viper.SetConfigName("config") // no need to include file extension
	viper.AddConfigPath("/home/dennis/GoProjects/hippy_bot/src/bot")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		log.Fatal(err)
	}

	url = viper.GetString("connectsense.API_URL")
	apiToken = viper.GetString("connectsense.API_TOKEN")
	deviceID = viper.GetString("connectsense.DEVICE_SERIAL")

	bot.RegisterPassiveCommand(
		"dogtemp",
		tempStatus)
}
