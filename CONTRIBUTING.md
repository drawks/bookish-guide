# Contributing to bookish-guide

Thank you for your interest in contributing to bookish-guide! This guide will help you create custom commands and plugins.

## Creating a Custom Command

Commands in bookish-guide are simple Go structs that implement the `bot.Command` interface. Here's a step-by-step guide:

### Step 1: Create Your Command

Create a new file in the `commands/` directory (e.g., `commands/mycommands.go`):

```go
package commands

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

// MyCustomCommand does something cool
type MyCustomCommand struct {
    // Add any state your command needs
    counter int
}

// Name returns the command name (used to trigger the command)
func (c *MyCustomCommand) Name() string {
    return "mycmd"
}

// Description provides a brief description shown in help
func (c *MyCustomCommand) Description() string {
    return "Does something awesome"
}

// Execute runs when the command is triggered
func (c *MyCustomCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
    // Your command logic here
    c.counter++
    message := fmt.Sprintf("Command executed %d times!", c.counter)
    _, err := s.ChannelMessageSend(m.ChannelID, message)
    return err
}
```

### Step 2: Register Your Command

In `main.go`, register your command:

```go
import "github.com/drawks/bookish-guide/commands"

func main() {
    // ... (existing setup code)
    
    // Register your custom command
    b.RegisterCommand(&commands.MyCustomCommand{})
    
    // ... (start bot)
}
```

### Step 3: Build and Test

```bash
go build -o bookish-guide main.go
./bookish-guide
```

## Command Interface

The `bot.Command` interface requires three methods:

```go
type Command interface {
    Name() string                                                              // Command name
    Description() string                                                       // Brief description
    Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error  // Command logic
}
```

### Execute Method Parameters

- `s *discordgo.Session`: The Discord session for sending messages and interacting with Discord
- `m *discordgo.MessageCreate`: The message that triggered the command (contains author, channel, content, etc.)
- `args []string`: Command arguments (everything after the command name, split by spaces)

## Best Practices

1. **Error Handling**: Always return errors from `Execute()`. The framework will log them and send an error message to the user.

2. **State Management**: If your command needs to maintain state:
   - Use struct fields
   - Consider thread-safety for concurrent access
   - Use constructor functions (e.g., `NewMyCommand()`) for initialization

3. **Argument Parsing**: Validate and parse arguments before use:
   ```go
   if len(args) < 1 {
       return fmt.Errorf("missing required argument")
   }
   ```

4. **Response Messages**: Use `s.ChannelMessageSend(m.ChannelID, "message")` to send responses

5. **Embed Messages**: For rich responses, use Discord embeds:
   ```go
   embed := &discordgo.MessageEmbed{
       Title:       "My Title",
       Description: "My Description",
       Color:       0x00ff00,
   }
   _, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
   ```

## Example Commands

Check out `commands/examples.go` for more examples:
- **RollCommand**: Demonstrates argument parsing
- **QuoteCommand**: Shows state management with a constructor
- **UpperCommand/LowerCommand**: Simple text transformation commands

## Testing Your Commands

Create tests in the `commands/` directory:

```go
package commands

import (
    "testing"
)

func TestMyCommand(t *testing.T) {
    cmd := &MyCustomCommand{}
    
    if cmd.Name() != "mycmd" {
        t.Errorf("Expected name 'mycmd', got '%s'", cmd.Name())
    }
}
```

## Common Patterns

### Command with Configuration

```go
type ConfigurableCommand struct {
    apiKey string
}

func NewConfigurableCommand(apiKey string) *ConfigurableCommand {
    return &ConfigurableCommand{apiKey: apiKey}
}
```

### Command with Database Access

```go
type DatabaseCommand struct {
    db *sql.DB
}

func NewDatabaseCommand(db *sql.DB) *DatabaseCommand {
    return &DatabaseCommand{db: db}
}
```

### Command with External API

```go
func (c *APICommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
    resp, err := http.Get("https://api.example.com/data")
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    // Process response and send to Discord
    // ...
}
```

## Submitting Your Command

If you've created a useful command and want to share it:

1. Ensure your code follows Go best practices
2. Add tests for your command
3. Update this guide if you've created a new pattern
4. Submit a pull request with a clear description

## Questions?

If you have questions or need help, please open an issue on GitHub.
