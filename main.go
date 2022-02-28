package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACKBOT_TOKEN_ID", "xxxx-111111111111-1111111111111-11111111111111111111111")
	os.Setenv("SLACKBOT_CHANNEL_ID", "11111111111")
	api := slack.New(os.Getenv("SLACKBOT_TOKEN_ID"))
	channelArr := []string{os.Getenv("SLACKBOT_CHANNEL_ID")}
	fileArr := []string{"Sachin_Gold_Price.ipynb"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
