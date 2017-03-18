package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/comail/colog"
	"net/http"
)

var (
	botInstance *discordgo.Session
	config      Configuration
)

func main() {
	colog.Register()

	//discord.ChannelMessageSend("292420088236343307", "Test")
	config = loadConfig()
	botInstance = botStart()

	http.HandleFunc("/webhook", incomingWebhook)
	http.ListenAndServe(":"+config.Port, nil)

	<-make(chan struct{})
	return
}
