package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// RollCommand implements a dice rolling command
type RollCommand struct{}

func (c *RollCommand) Name() string {
	return "roll"
}

func (c *RollCommand) Description() string {
	return "Rolls a dice (1-6) or a custom sided dice (!roll 20)"
}

func (c *RollCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	sides := 6
	
	if len(args) > 0 {
		_, err := fmt.Sscanf(args[0], "%d", &sides)
		if err != nil || sides < 2 {
			return fmt.Errorf("invalid dice sides, must be a number >= 2")
		}
	}

	result := rand.Intn(sides) + 1
	message := fmt.Sprintf("🎲 You rolled a **%d** (d%d)", result, sides)
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	return err
}

// QuoteCommand implements a command that stores and recalls quotes
type QuoteCommand struct {
	quotes []string
}

func NewQuoteCommand() *QuoteCommand {
	return &QuoteCommand{
		quotes: []string{
			"The only way to do great work is to love what you do. - Steve Jobs",
			"Innovation distinguishes between a leader and a follower. - Steve Jobs",
			"Code is like humor. When you have to explain it, it's bad. - Cory House",
			"First, solve the problem. Then, write the code. - John Johnson",
		},
	}
}

func (c *QuoteCommand) Name() string {
	return "quote"
}

func (c *QuoteCommand) Description() string {
	return "Get a random inspirational quote"
}

func (c *QuoteCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(c.quotes) == 0 {
		return fmt.Errorf("no quotes available")
	}

	quote := c.quotes[rand.Intn(len(c.quotes))]
	_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("💬 %s", quote))
	return err
}

// UpperCommand converts text to uppercase
type UpperCommand struct{}

func (c *UpperCommand) Name() string {
	return "upper"
}

func (c *UpperCommand) Description() string {
	return "Converts text to UPPERCASE"
}

func (c *UpperCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no text provided")
	}
	
	text := strings.Join(args, " ")
	_, err := s.ChannelMessageSend(m.ChannelID, strings.ToUpper(text))
	return err
}

// LowerCommand converts text to lowercase
type LowerCommand struct{}

func (c *LowerCommand) Name() string {
	return "lower"
}

func (c *LowerCommand) Description() string {
	return "Converts text to lowercase"
}

func (c *LowerCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no text provided")
	}
	
	text := strings.Join(args, " ")
	_, err := s.ChannelMessageSend(m.ChannelID, strings.ToLower(text))
	return err
}
