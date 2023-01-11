package config

import (
	"github.com/spf13/viper"
)

const (
	tokenKey = "LINK_SCREEN_TGBOT_TOKEN"
)

type Messages struct {
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
	Commands         Commands
	Localizations    map[string]Messages
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

	if err := viper.UnmarshalKey("commands", &cfg.Commands); err != nil {
		return err
	}

	cfg.Localizations = make(map[string]Messages)

	engMessages := Messages{}
	if err := viper.UnmarshalKey("messages.en.response", &engMessages.Responses); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("messages.en.errors", &engMessages.Errors); err != nil {
		return err
	}
	cfg.Localizations["en"] = engMessages

	ruMessages := Messages{}
	if err := viper.UnmarshalKey("messages.ru.response", &ruMessages.Responses); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("messages.ru.errors", &ruMessages.Errors); err != nil {
		return err
	}
	cfg.Localizations["ru"] = ruMessages

	return nil
}

func grabEnvVars(cfg *Config) error {
	if err := viper.BindEnv(tokenKey); err != nil {
		return err
	}
	cfg.TelegramBotToken = viper.GetString(tokenKey)

	return nil
}

func setUp() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
