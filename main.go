package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/slack-go/slack"
)

type message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type config struct {
	Token    string    `json:"bot_token"`
	Channels []message `json:"channels"`
}

func read_config(filename string) config {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Panicf("Failed to read file: %v", err)
	}

	log.Printf("Running with config: %v", string(file))
	config := config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Panicf("Failed to parse: %v", err)
	}

	return config
}

type name_to_id map[string]string

func get_channel_name_to_id(api *slack.Client) name_to_id {
	name_to_id := make(name_to_id)

	cursor := ""

	for {
		var groups []slack.Channel
		var err error
		groups, cursor, err = api.GetConversations(&slack.GetConversationsParameters{
			ExcludeArchived: true,
			Limit:           1000,
			Types: []string{
				"public_channel",
			},
			Cursor: cursor,
		})

		if err != nil {
			log.Panicf("Failed to get conversations: %v", err)
			return nil
		}

		for _, group := range groups {
			_, ok := name_to_id[group.Name]
			if ok {
				log.Panicf("There are two channels with name %v", group.Name)
			}
			name_to_id[group.Name] = group.ID
		}

		if cursor == "" {
			break
		}
	}

	return name_to_id
}

func send_single_message(api *slack.Client, name_to_id *name_to_id, message *message) {
	log.Printf("Sending message '%v' to channel '%v'", message.Text, message.Channel)
	channel_id := (*name_to_id)[message.Channel]
	_, _, err := api.PostMessage(
		channel_id,
		slack.MsgOptionText(message.Text, false))

	if err != nil {
		log.Panicf("Failed to send message to channel %v: %v", message.Channel, err)
	}
}

func send_messages(config *config) {
	api := slack.New(config.Token)
	name_to_id := get_channel_name_to_id(api)

	for _, message := range config.Channels {
		send_single_message(api, &name_to_id, &message)
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Panicf("Usage: %v config_filename", os.Args[0])
	}

	filename := os.Args[1]

	log.Printf("Reading config from file %v", filename)
	config := read_config(filename)

	log.Printf("Sending %v messages", len(config.Channels))
	send_messages(&config)

	log.Println("Done")
}
