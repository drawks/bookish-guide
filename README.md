# bookish-guide

A basic extensible Discord bot framework in Golang

## Features

- **Extensible Command System**: Easy-to-use interface for adding custom commands
- **Flexible Configuration**: Support for both JSON config files and environment variables
- **Built-in Commands**: Includes ping, echo, and help commands as examples
- **Simple Plugin Architecture**: Register commands with a simple interface

## Installation

### Prerequisites

- Go 1.18 or higher
- A Discord Bot Token ([create one here](https://discord.com/developers/applications))

### Setup

1. Clone the repository:
```bash
git clone https://github.com/drawks/bookish-guide.git
cd bookish-guide
```

2. Install dependencies:
```bash
go mod download
```

3. Configure your bot (choose one method):

   **Method A: Using environment variables**
   ```bash
   export DISCORD_TOKEN="your_bot_token_here"
   export DISCORD_PREFIX="!"  # Optional, defaults to "!"
   ```

   **Method B: Using a config file**
   ```bash
   cp config.example.json config.json
   # Edit config.json with your bot token
   ```

## Usage

### Running the Bot

Using environment variables:
```bash
go run main.go
```

Using a config file:
```bash
go run main.go -config config.json
```

### Building the Bot

```bash
go build -o bookish-guide main.go
./bookish-guide
```

## Creating Custom Commands

The framework is designed to be easily extensible. Here's how to create custom commands:

1. Create a new command by implementing the `bot.Command` interface:

```go
package commands

import (
    "github.com/bwmarrin/discordgo"
)

type MyCommand struct{}

func (c *MyCommand) Name() string {
    return "mycommand"
}

func (c *MyCommand) Description() string {
    return "Description of what my command does"
}

func (c *MyCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
    // Your command logic here
    _, err := s.ChannelMessageSend(m.ChannelID, "Command executed!")
    return err
}
```

2. Register your command in `main.go`:

```go
b.RegisterCommand(&commands.MyCommand{})
```

## Built-in Commands

### Basic Commands
- `!ping` - Responds with "Pong!"
- `!echo <message>` - Echoes back the provided message
- `!help` - Shows all available commands

### Example Commands (demonstrating extensibility)
- `!roll [sides]` - Rolls a dice (default 6-sided, or specify number of sides)
- `!quote` - Get a random inspirational quote
- `!upper <text>` - Converts text to UPPERCASE
- `!lower <text>` - Converts text to lowercase

## Project Structure

```
bookish-guide/
├── bot/           # Core bot framework
│   └── bot.go     # Bot logic and command handling
├── commands/      # Built-in and custom commands
│   └── basic.go   # Example commands
├── config/        # Configuration management
│   └── config.go  # Config loading utilities
├── main.go        # Entry point
└── README.md      # This file
```

## Architecture

The framework follows a plugin-based architecture:

1. **Bot Core** (`bot/bot.go`): Handles Discord connection, message routing, and command execution
2. **Command Interface**: Simple interface that all commands must implement
3. **Command Registration**: Commands are registered with the bot at startup
4. **Message Handler**: Routes incoming messages to appropriate commands based on prefix

## License

MIT License - see LICENSE file for details
