# slack-file-bot
It is a file bot for uploading file to channels in a workspace of slack.

## Instructions
* set the env variables values in env file.
* The env variables name are:
    * "SLACK_BOT_TOKEN" (The value can be accessed or created from "api.slack.com/apps")
    * "CHANNEL_ID" (channel id of channel in which bot will do the file operations)
* run "go build" to build the executable file
* run "go run main.go" to start the program

#### Note: In the program, find the below line of code and change the path of the file for file upload
#### filesArr := []string{"fileToUpload.txt"}
