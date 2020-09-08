package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

func main() {

	token := os.Getenv("SLACK_TOKEN")
	fmt.Println(token)
	api := slack.New(token)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			msg := ev.Msg
			// info := rtm.GetInfo()
			// prefix := fmt.Sprintf("<@%s> ", info.User.ID)

			// if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {
			// 	respond(rtm, ev, prefix)
			// }

			response := "Good. How are you?"
			rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))

		case *slack.ConnectedEvent:
			fmt.Println("Connection counter:", ev.ConnectionCount)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid Token")
			return
		}
	}
}

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	acceptedGreetings := map[string]bool{
		"what's up?": true,
		"hey!":       true,
		"yo":         true,
	}
	acceptedHowAreYou := map[string]bool{
		"how's it going?": true,
		"how are ya?":     true,
		"feeling okay?":   true,
		"howya?":          true,
	}

	if acceptedGreetings[text] {
		response = "What's up buddy!?!?!"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "Good. How are you?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}
