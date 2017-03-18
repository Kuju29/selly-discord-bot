package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func incomingWebhook(w http.ResponseWriter, r *http.Request) {
	secret := r.URL.Query().Get("secret")
	if secret == config.WebhookSecret {
		var webhook Webhook
		err := json.NewDecoder(r.Body).Decode(&webhook)
		if err != nil {
			log.Print("error: Error decoding JSON: " + err.Error())
			http.Error(w, err.Error(), 400)
			//return
		}

		value, err := strconv.ParseFloat(webhook.Value, 64)
		if err != nil {
			log.Print("error: Error parsing value: " + err.Error())
			http.Error(w, err.Error(), 500)
		}

		// Nicer formatting to come
		message := fmt.Sprintf("Order received of %0.2f %s - %s - https://selly.gg/orders/%s", value, webhook.Currency, webhook.Email, webhook.ID)

		if config.SendToChannel {
			botInstance.ChannelMessageSend(config.ChannelID, message)
		}

		io.WriteString(w, "Received correctly")
	} else {
		io.WriteString(w, "Invalid secret")
		log.Print("warn: Invalid secret attempted")
	}
}
