package functions

import (
	"context"
	"kolossus/internal"
	"log"
)

func main() {
	bot, err := internal.NewBot()
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	bot.Start(context.Background())
	defer bot.Stop(context.Background())
}
