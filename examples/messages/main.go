package main

import (
	"fmt"
	"log"

	webexteams "github.com/thpiron/webex-teams/sdk"
)

// Client is Webex Teams API client
var Client *webexteams.Client

func main() {
	Client = webexteams.NewClient()

	/*

		MESSAGES

	*/

	myRoomID := "" // Change to your testing room

	// POST messages - Text Message

	message := &webexteams.MessageCreateRequest{
		Text:   "This is a text message",
		RoomID: myRoomID,
	}
	newTextMessage, _, err := Client.Messages.CreateMessage(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newTextMessage.ID, newTextMessage.Text, newTextMessage.Created)

	// POST messages - Markdown Message

	markDownMessage := &webexteams.MessageCreateRequest{
		Markdown: "This is a markdown message. *Italic*, **bold** and ***italic/bold***.",
		RoomID:   myRoomID,
	}
	newMarkDownMessage, _, err := Client.Messages.CreateMessage(markDownMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("POST:", newMarkDownMessage.ID, newMarkDownMessage.Markdown, newMarkDownMessage.Created)

	// GET messages
	messageQueryParams := &webexteams.ListMessagesQueryParams{
		RoomID: myRoomID,
	}

	messages, _, err := Client.Messages.ListMessages(messageQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, message := range messages.Items {
		fmt.Println("GET:", id, message.ID, message.Text, message.Created)
	}

	// GET messages/<ID>

	htmlMessageGet, _, err := Client.Messages.GetMessage(newMarkDownMessage.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET <ID>:", htmlMessageGet.ID, htmlMessageGet.Text, htmlMessageGet.Created)

	// DELETE messages<ID>

	resp, err := Client.Messages.DeleteMessage(newTextMessage.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DELETE:", resp.StatusCode())

}
