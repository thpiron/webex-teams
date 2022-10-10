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

	myRoomID := ""   // Change to your testing room
	webHookURL := "" // Change this to your test URL

	// POST webhooks

	webhookRequest := &webexteams.WebhookCreateRequest{
		Name:      "Webhook - Test",
		TargetURL: webHookURL,
		Resource:  "messages",
		Event:     "created",
		Filter:    "roomId=" + myRoomID,
	}

	testWebhook, _, err := Client.Webhooks.CreateWebhook(webhookRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("POST:", testWebhook.ID, testWebhook.Name, testWebhook.TargetURL, testWebhook.Created)

	// GET webhooks

	webhooksQueryParams := &webexteams.ListWebhooksQueryParams{
		Max: 10,
	}

	webhooks, _, err := Client.Webhooks.ListWebhooks(webhooksQueryParams)
	if err != nil {
		log.Fatal(err)
	}
	for id, webhook := range webhooks.Items {
		fmt.Println("GET:", id, webhook.ID, webhook.Name, webhook.TargetURL, webhook.Created)
	}

	// GET webhooks/<ID>
	webhook, _, err := Client.Webhooks.GetWebhook(testWebhook.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GET <ID>:", webhook.ID, webhook.Name, webhook.TargetURL, webhook.Created)

	updateWebhookRequest := &webexteams.WebhookUpdateRequest{
		Name:      "Webhook Update - Test",
		TargetURL: webHookURL,
	}

	// PUT webhooks/<ID>
	updatedWebhook, _, err := Client.Webhooks.UpdateWebhook(testWebhook.ID, updateWebhookRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PUT:", updatedWebhook.ID, updatedWebhook.Name, updatedWebhook.TargetURL, updatedWebhook.Created)

	// DELETE webhooks/<ID>
	resp, err := Client.Webhooks.DeleteWebhook(testWebhook.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DELETE:", resp.StatusCode())

}
