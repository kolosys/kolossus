package main

import (
	"context"
	"log"
	"os"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/server"
)

func NewBot() (*discord.Bot, error) {
	return discord.New(&discord.Options{
		Token:   os.Getenv("DISCORD_TOKEN"),
		Intents: gateway.IntentsAll,
		Server: &server.Options{
			Enabled: true,
		},
	})
}

func main() {
	bot, err := NewBot()
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	bot.Start(context.Background())
	defer bot.Stop(context.Background())
}
