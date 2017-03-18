package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func botStart() *discordgo.Session {
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Print("error: Encountered error when creating Discord session: " + err.Error())
	}

	err = discord.Open()
	if err != nil {
		log.Print("error: Encountered error when opening connection: " + err.Error())
	}

	log.Print("info: Bot started")

	return discord
}
