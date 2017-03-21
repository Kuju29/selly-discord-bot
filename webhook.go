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
		_ = json.NewDecoder(r.Body).Decode(&webhook)

		value, err := strconv.ParseFloat(webhook.Value, 64)
		if err != nil {
			log.Print("error: Error parsing value: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if config.SendToChannel {
			botInstance.ChannelMessageSend(config.ChannelID, formatMessage(webhook, value))
		}

		io.WriteString(w, "Received correctly")
	} else {
		http.Error(w, "Invalid secret", http.StatusForbidden)
		log.Print("warn: Invalid secret attempted")
	}
}

func formatMessage(webhook Webhook, value float64) string {
	log.Print(webhook)
	switch webhookType := webhook.WebhookType; webhookType {
	case 1:
		return fmt.Sprintf("Order received of %0.2f %s - %s - https://selly.gg/orders/%s", value, webhook.Currency, webhook.Email, webhook.ID)
	case 2:
		return fmt.Sprintf("PayPal chargeback of %0.2f %s - %s - https://selly.gg/orders/%s", value, webhook.Currency, webhook.Email, webhook.ID)
	}
	return "Invalid webhook type"
}
