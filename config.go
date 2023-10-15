package main

import (
	"github.com/spf13/viper"
	"strings"
)

// Config structure holds all configuration data
type Config struct {
	BotToken           string `mapstructure:"telegram_bot_token"`
	DbURL              string `mapstructure:"mongodb_connection_string"`
	DbName             string `mapstructure:"mongodb_database_name"`
	OpenAIKey          string `mapstructure:"openai_api_key"`
	ImageAPIKey        string `mapstructure:"sd_key"`
	S3BucketName       string `mapstructure:"s3_bucket_name"`
	AWSKey             string `mapstructure:"aws_key"`
	AWSSecret          string `mapstructure:"aws_secret"`
	AWSRegion          string `mapstructure:"aws_region"`
	GoogleSearchAPIKey string `mapstructure:"serper_api_key"`
	// Add other fields as needed
}

// LoadConfig loads configuration data
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("telegram_bot") // will be uppercased automatically
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
