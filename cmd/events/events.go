package events

import (
	"context"
	"log"

	"github.com/kolosys/discord"
	"github.com/kolosys/discord/events"
)

type EventsModule struct{}

func (e EventsModule) Register(bot *discord.Bot) error {
	bot.OnReady(onReady)
	bot.OnGuildCreate(onGuildCreate)
	return nil
}

func onReady(ctx context.Context, e *events.ReadyEvent) {
	log.Printf("Bot ready! Logged in as %s", e.User.Username)
	log.Printf("Serving %d guilds", len(e.Guilds))
}

func onGuildCreate(ctx context.Context, e *events.GuildCreateEvent) {
	log.Printf("Joined guild: %s (%s)", e.Name, e.ID)
}
