package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACKBOT_TOKEN_ID", "xoxb-3150876980551-3158832191414-0KHeSzxaJ5mvbCqGFPlisYAA")
	os.Setenv("SLACKBOT_CHANNEL_ID", "C034ERSVAEB")
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
