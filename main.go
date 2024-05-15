package main

import (
	"flag"
	"log"

	"github.com/firman-alam/tele_bot.git/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	
	tgClient = telegram.New(tgBotHost, mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() (string, error) {
	// bot -tg-bot-token
	token := flag.String("tg-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token, nil
}