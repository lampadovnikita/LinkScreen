package config

import (
	"github.com/spf13/viper"
)

type Messages struct {
	Commands
	Responses
	Errors
}

type Commands struct {
	Start string `mapstructure:"start"`
	Help  string `mapstructure:"help"`
}

type Responses struct {
	Start          string `mapstructure:"start"`
	Help           string `mapstructure:"help"`
	UnknownCommand string `mapstructure:"unknown_command"`
}

type Errors struct {
	Default string `mapstructure:"default"`
	NonURL  string `mapstructure:"non_url"`
}

type Config struct {
	TelegramBotToken string
	Messages         Messages
}

func Init() (*Config, error) {
	if err := setUp(); err != nil {
		return nil, err
	}

	var cfg *Config = &Config{}
	if err := unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := grabEnvVars(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.commands", &cfg.Messages.Commands); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("messages.response", &cfg.Messages.Responses); err != nil {
		return err
	}

	return nil
}

func grabEnvVars(cfg *Config) error {
	if err := viper.BindEnv("LINK_SCREEN_TGBOT_TOKEN"); err != nil {
		return err
	}
	cfg.TelegramBotToken = viper.GetString("LINK_SCREEN_TGBOT_TOKEN")

	return nil
}

func setUp() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
