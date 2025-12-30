package main

import (
	"context"
	"log"
	"os"

	"github.com/kolosys/discord"
)

func main() {
	bot, err := discord.New(&discord.Options{
		Token: os.Getenv("DISCORD_TOKEN"),
	})

	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	if err := bot.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}
