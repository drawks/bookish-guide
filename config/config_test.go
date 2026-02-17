package config

import (
	"os"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	// Save original env vars
	originalToken := os.Getenv("DISCORD_TOKEN")
	originalPrefix := os.Getenv("DISCORD_PREFIX")
	defer func() {
		os.Setenv("DISCORD_TOKEN", originalToken)
		os.Setenv("DISCORD_PREFIX", originalPrefix)
	}()

	t.Run("with token and prefix", func(t *testing.T) {
		os.Setenv("DISCORD_TOKEN", "test_token")
		os.Setenv("DISCORD_PREFIX", "$$")

		cfg, err := LoadConfigFromEnv()
		if err != nil {
			t.Fatalf("LoadConfigFromEnv() error = %v", err)
		}

		if cfg.Token != "test_token" {
			t.Errorf("Token = %q, want %q", cfg.Token, "test_token")
		}

		if cfg.Prefix != "$$" {
			t.Errorf("Prefix = %q, want %q", cfg.Prefix, "$$")
		}
	})

	t.Run("with token only (default prefix)", func(t *testing.T) {
		os.Setenv("DISCORD_TOKEN", "test_token")
		os.Unsetenv("DISCORD_PREFIX")

		cfg, err := LoadConfigFromEnv()
		if err != nil {
			t.Fatalf("LoadConfigFromEnv() error = %v", err)
		}

		if cfg.Prefix != "!" {
			t.Errorf("Prefix = %q, want %q (default)", cfg.Prefix, "!")
		}
	})

	t.Run("without token", func(t *testing.T) {
		os.Unsetenv("DISCORD_TOKEN")

		_, err := LoadConfigFromEnv()
		if err == nil {
			t.Error("LoadConfigFromEnv() expected error when token is missing")
		}
	})
}
