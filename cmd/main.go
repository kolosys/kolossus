package main

import (
	"context"
	"kolossus/cmd/bot"
	"log"
)

func main() {
	bot, err := bot.NewBot()
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	bot.Start(context.Background())
	defer bot.Stop(context.Background())
}
