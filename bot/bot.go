package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Command represents a bot command
type Command interface {
	Name() string
	Description() string
	Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error
}

// Bot represents the Discord bot
type Bot struct {
	session  *discordgo.Session
	commands map[string]Command
	prefix   string
}

// New creates a new Bot instance
func New(token, prefix string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}

	bot := &Bot{
		session:  session,
		commands: make(map[string]Command),
		prefix:   prefix,
	}

	// Register message handler
	session.AddHandler(bot.messageHandler)

	// Set intents
	session.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	return bot, nil
}

// RegisterCommand registers a new command
func (b *Bot) RegisterCommand(cmd Command) {
	b.commands[cmd.Name()] = cmd
	log.Printf("Registered command: %s", cmd.Name())
}

// Start starts the bot
func (b *Bot) Start() error {
	err := b.session.Open()
	if err != nil {
		return fmt.Errorf("error opening connection: %w", err)
	}

	log.Printf("Bot is now running. Press CTRL-C to exit.")
	log.Printf("Command prefix: %s", b.prefix)

	// Wait for interrupt signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return b.Stop()
}

// Stop stops the bot
func (b *Bot) Stop() error {
	log.Println("Shutting down bot...")
	return b.session.Close()
}

// messageHandler handles incoming messages
func (b *Bot) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Check if message starts with prefix
	if len(m.Content) < len(b.prefix) || m.Content[:len(b.prefix)] != b.prefix {
		return
	}

	// Parse command and arguments
	content := m.Content[len(b.prefix):]
	args := parseArgs(content)
	if len(args) == 0 {
		return
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	// Find and execute command
	if cmd, ok := b.commands[cmdName]; ok {
		err := cmd.Execute(s, m, cmdArgs)
		if err != nil {
			log.Printf("Error executing command %s: %v", cmdName, err)
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error: %v", err))
		}
	}
}

// parseArgs parses command arguments from a string
func parseArgs(content string) []string {
	args := []string{}
	current := ""
	inQuotes := false

	for _, char := range content {
		switch char {
		case ' ':
			if inQuotes {
				current += string(char)
			} else if current != "" {
				args = append(args, current)
				current = ""
			}
		case '"':
			inQuotes = !inQuotes
		default:
			current += string(char)
		}
	}

	if current != "" {
		args = append(args, current)
	}

	return args
}

// GetPrefix returns the bot's command prefix
func (b *Bot) GetPrefix() string {
	return b.prefix
}

// GetCommands returns all registered commands
func (b *Bot) GetCommands() map[string]Command {
	return b.commands
}
