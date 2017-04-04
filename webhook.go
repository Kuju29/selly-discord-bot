package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
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
			botInstance.ChannelMessageSendEmbed(config.ChannelID, formatMessage(webhook, value))
		}

		io.WriteString(w, "Received correctly")
	} else {
		http.Error(w, "Invalid secret", http.StatusForbidden)
		log.Print("warn: Invalid secret attempted")
	}
}

func formatMessage(webhook Webhook, value float64) *discordgo.MessageEmbed {
	log.Print(webhook)
	embed := new(discordgo.MessageEmbed)
	fields := make([]*discordgo.MessageEmbedField, 0, 10)
	thumbnail := new(discordgo.MessageEmbedThumbnail)

	thumbnail.URL = "https://selly.gg/images/apple-touch-icon-180x180.png"
	embed.Thumbnail = thumbnail
	embed.Description = webhook.ID
	embed.URL = fmt.Sprintf("https://selly.gg/orders/%s", webhook.ID)

	// Value
	valueField := new(discordgo.MessageEmbedField)
	valueField.Name = "Amount"
	valueField.Value = fmt.Sprintf("%0.2f %s", value, webhook.Currency)
	valueField.Inline = true

	// Email
	emailField := new(discordgo.MessageEmbedField)
	emailField.Name = "Email"
	emailField.Value = webhook.Email
	emailField.Inline = true

	// IP
	ipField := new(discordgo.MessageEmbedField)
	ipField.Name = "IP Address"
	ipField.Value = webhook.IPAddress
	ipField.Inline = true

	// Country Code
	countryField := new(discordgo.MessageEmbedField)
	countryField.Name = "Country Code"
	countryField.Value = webhook.CountryCode
	countryField.Inline = true

	// Gateway Code
	gatewayField := new(discordgo.MessageEmbedField)
	gatewayField.Name = "Gateway"
	gatewayField.Value = webhook.Gateway
	gatewayField.Inline = true

	// Risk Level Code
	riskLevelField := new(discordgo.MessageEmbedField)
	riskLevelField.Name = "Risk Level"
	riskLevelField.Value = fmt.Sprintf("%d", webhook.RiskLevel)
	riskLevelField.Inline = true

	// Created
	createdAtField := new(discordgo.MessageEmbedField)
	createdAtField.Name = "Created At"
	createdAtField.Value = webhook.CreatedAt
	createdAtField.Inline = false

	fields = append(fields, valueField)
	fields = append(fields, emailField)
	fields = append(fields, countryField)
	fields = append(fields, ipField)
	fields = append(fields, gatewayField)
	fields = append(fields, riskLevelField)
	fields = append(fields, createdAtField)

	switch webhookType := webhook.WebhookType; webhookType {
	case 1:
		embed.Title = "Order Received"
	case 2:
		embed.Title = "PayPal Chargeback"
	}
	embed.Fields = fields
	return embed
}
