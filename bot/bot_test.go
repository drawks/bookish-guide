package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "echo test message",
			expected: []string{"echo", "test", "message"},
		},
		{
			input:    `echo "hello world"`,
			expected: []string{"echo", "hello world"},
		},
		{
			input:    "single",
			expected: []string{"single"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  spaces  between  ",
			expected: []string{"spaces", "between"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := parseArgs(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("parseArgs(%q) length = %d, want %d", tt.input, len(result), len(tt.expected))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("parseArgs(%q)[%d] = %q, want %q", tt.input, i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func TestBotPrefix(t *testing.T) {
	// Create a bot with empty token (won't connect but ok for testing)
	bot := &Bot{
		commands: make(map[string]Command),
		prefix:   "!",
	}

	if bot.GetPrefix() != "!" {
		t.Errorf("GetPrefix() = %q, want %q", bot.GetPrefix(), "!")
	}
}

func TestRegisterCommand(t *testing.T) {
	bot := &Bot{
		commands: make(map[string]Command),
		prefix:   "!",
	}

	cmd := &mockCommand{name: "test"}
	bot.RegisterCommand(cmd)

	commands := bot.GetCommands()
	if len(commands) != 1 {
		t.Errorf("Expected 1 command, got %d", len(commands))
	}

	if _, ok := commands["test"]; !ok {
		t.Error("Command 'test' not found in registered commands")
	}
}

// mockCommand for testing
type mockCommand struct {
	name string
}

func (m *mockCommand) Name() string {
	return m.name
}

func (m *mockCommand) Description() string {
	return "Mock command"
}

func (m *mockCommand) Execute(s *discordgo.Session, msg *discordgo.MessageCreate, args []string) error {
	return nil
}
