package config

import (
	"os"

	"github.com/kolosys/helix/logs"
)

var DiscordToken = os.Getenv("DISCORD_TOKEN")
var SupportServerID = os.Getenv("DISCORD_GUILD_ID")

func init() {
	if DiscordToken == "" {
		logs.Fatal("DISCORD_TOKEN is not set")
	}
	if SupportServerID == "" {
		logs.Warn("DISCORD_GUILD_ID is not set, some features will be disabled")
	}
}
