package main

import (
	"fmt"
	"os"
	"pacgobot/anilisthandler"
	"pacgobot/handlers"

	"github.com/bwmarrin/discordgo"
)

type BotHandler struct {
	Name        string
	Description string
	ShortHand   string
	HandlerFunc func(s *discordgo.Session, m *discordgo.MessageCreate)
}

var PREFIX = "p."

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Println("Error: BOT_TOKEN environment variable not set!")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	handlers.PREFIX = PREFIX
	anilisthandler.PREFIX = PREFIX

	BasicHandlers := []BotHandler{
		{
			Name:        "Avatar",
			Description: "Steals users's current avatar",
			ShortHand:   "av",
			HandlerFunc: handlers.AvatarHandler,
		},
		{
			Name:        "Ping",
			Description: "Ping to bot's backend server",
			ShortHand:   "ping",
			HandlerFunc: handlers.PingHandler,
		},
	}

	AnilistHandlers := []BotHandler{
		{
			Name:        "Anilist User",
			Description: "Fetches Anilist user",
			ShortHand:   "aniuser",
			HandlerFunc: anilisthandler.AnilistUserHandler,
		},
	}

	for _, funcs := range BasicHandlers {
		fmt.Printf("[INFO][BasicHandlers] Adding %s\n", funcs.Name)
		dg.AddHandler(funcs.HandlerFunc)
	}

	for _, funcs := range AnilistHandlers {
		fmt.Printf("[INFO][AnilistHandlers] Adding %s\n", funcs.Name)
		dg.AddHandler(funcs.HandlerFunc)
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening websocket connection:", err)
		return
	}

	fmt.Println("[INFO] Bot is now running. Press CTRL+C to exit.")
	<-make(chan struct{})
}
