package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-6244363649893-6247208597795-8YQTAQBjid875I3JoDB2lls9")
	os.Setenv("CHANNEL_ID", "C067MT9U59P")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	filesArr := []string{"fileToUpload.txt"}

	for i := 0; i < len(filesArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     filesArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	}
}
