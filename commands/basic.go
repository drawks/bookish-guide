package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/drawks/bookish-guide/bot"
)

// PingCommand implements a simple ping command
type PingCommand struct{}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Responds with 'Pong!'"
}

func (c *PingCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
	return err
}

// EchoCommand implements an echo command
type EchoCommand struct{}

func (c *EchoCommand) Name() string {
	return "echo"
}

func (c *EchoCommand) Description() string {
	return "Echoes back the provided message"
}

func (c *EchoCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no message provided")
	}
	message := strings.Join(args, " ")
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

// HelpCommand implements a help command
type HelpCommand struct {
	bot *bot.Bot
}

func NewHelpCommand(b *bot.Bot) *HelpCommand {
	return &HelpCommand{bot: b}
}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Shows available commands"
}

func (c *HelpCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	commands := c.bot.GetCommands()
	prefix := c.bot.GetPrefix()

	var message strings.Builder
	message.WriteString("**Available Commands:**\n")
	
	for _, cmd := range commands {
		message.WriteString(fmt.Sprintf("`%s%s` - %s\n", prefix, cmd.Name(), cmd.Description()))
	}

	_, err := s.ChannelMessageSend(m.ChannelID, message.String())
	return err
}
