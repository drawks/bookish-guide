package main

import (
	"flag"
	"log"

	"github.com/drawks/bookish-guide/bot"
	"github.com/drawks/bookish-guide/commands"
	"github.com/drawks/bookish-guide/config"
)

func main() {
	configFile := flag.String("config", "", "Path to config file (optional, uses env vars if not provided)")
	flag.Parse()

	var cfg *config.Config
	var err error

	// Load configuration
	if *configFile != "" {
		cfg, err = config.LoadConfig(*configFile)
		if err != nil {
			log.Fatalf("Error loading config from file: %v", err)
		}
	} else {
		cfg, err = config.LoadConfigFromEnv()
		if err != nil {
			log.Fatalf("Error loading config from environment: %v", err)
		}
	}

	// Create bot
	b, err := bot.New(cfg.Token, cfg.Prefix)
	if err != nil {
		log.Fatalf("Error creating bot: %v", err)
	}

	// Register commands
	b.RegisterCommand(&commands.PingCommand{})
	b.RegisterCommand(&commands.EchoCommand{})
	b.RegisterCommand(commands.NewHelpCommand(b))
	
	// Register example commands to demonstrate extensibility
	b.RegisterCommand(&commands.RollCommand{})
	b.RegisterCommand(commands.NewQuoteCommand())
	b.RegisterCommand(&commands.UpperCommand{})
	b.RegisterCommand(&commands.LowerCommand{})

	// Start bot
	if err := b.Start(); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
