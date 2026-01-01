package main

import (
	"context"
	"kolossus/cmd/commands"
	"kolossus/cmd/config"
	"kolossus/cmd/events"
	"kolossus/cmd/routes"
	"log"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/gateway"
	"github.com/kolosys/discord/server"
)

func main() {
	bot, err := discord.New(&discord.Options{
		Token:   config.DiscordToken,
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentMessageContent,
		Server: &server.Options{
			Enabled: true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	if err := discord.RegisterModules(bot,
		&commands.CommandModule{},
		&events.EventsModule{},
		&routes.RoutesModule{},
	); err != nil {
		log.Fatalf("Failed to register modules: %v", err)
	}

	if err := bot.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}
}
