package main

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Configuration struct {
	Token         string
	ChannelID     string
	UserID        string
	SendToChannel bool
	SendToUser    bool
	WebhookSecret string
	Port          string
}

type Webhook struct {
	ID            string      `json:"id"`
	ProductID     string      `json:"product_id"`
	Email         string      `json:"email"`
	IPAddress     string      `json:"ip_address"`
	CountryCode   string      `json:"country_code"`
	UserAgent     string      `json:"user_agent"`
	Value         string      `json:"value"`
	Currency      string      `json:"currency"`
	Gateway       string      `json:"gateway"`
	RiskLevel     int         `json:"risk_level"`
	Status        int         `json:"status"`
	Delivered     string      `json:"delivered"`
	CryptoValue   interface{} `json:"crypto_value"`
	CryptoAddress interface{} `json:"crypto_address"`
	Referral      string      `json:"referral"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type MessageEmbed *discordgo.MessageEmbed
