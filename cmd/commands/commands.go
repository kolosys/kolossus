package commands

import (
	"github.com/kolosys/discord"
	"github.com/kolosys/discord/commands"
)

type CommandModule struct{}

func (c CommandModule) Register(bot *discord.Bot) error {
	bot.Slash("ping", "Check if the bot is alive", handlePing)
	bot.Slash("hello", "Say hello", handleHello)
	return nil
}

func handlePing(ctx *commands.Context) error {
	return ctx.Reply("Pong!")
}

func handleHello(ctx *commands.Context) error {
	return ctx.Reply("Hello, " + ctx.User.Username + "!")
}
