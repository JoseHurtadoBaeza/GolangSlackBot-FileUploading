package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))  // Specify this environment variable in the .bashrc file
	channelArr := []string{os.Getenv("CHANNEL_ID")} // Specify this environment variable in the .bashrc file
	fileArr := []string{"./HelloWorld.txt"}         // The file that will be sending out uploading basically

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
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)

	}

}
