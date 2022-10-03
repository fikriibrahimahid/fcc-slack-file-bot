package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-775315994100-4161079360325-DNcPlrsCQP0XCVBNfpHgInsy")
	os.Setenv("CHANNEL_ID", "CNTNYCJMN")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{}
	// fileArr = append(fileArr, "C01WORLDUNIVERSITYRANKINGS.csv")
	fileArr = append(fileArr, "C02BANKMARKETING.csv")
	// fileArr = append(fileArr, "C03NFLOFFENSEWEEK3.csv")

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Println()
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	}
}
