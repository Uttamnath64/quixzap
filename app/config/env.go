package config

import (
	"crypto/rsa"
	"time"

	"github.com/spf13/viper"
)

// env config
type AppEnv struct {
	MySQL  MySQLConfig
	Mongo  MongoConfig
	Redis  RedisConfig
	Kafka  KafkaConfig
	Server ServerConfig
	Auth   AuthConfig
}

type MySQLConfig struct {
	DNS struct {
		ReadWrite string `mapstructure:"MYSQL_MAIN_DNS"`
		ReadOnly  string `mapstructure:"MYSQL_MAIN_READ_DNS"`
		Log       string `mapstructure:"MYSQL_LOG_DNS"`
	}
	DB struct {
		Main string `mapstructure:"MYSQL_MAIN_DB"`
		Log  string `mapstructure:"MYSQL_LOG_DB"`
	}
}

type MongoConfig struct {
	Url      string `mapstructure:"MONGO_URL"`
	DB       string `mapstructure:"MONGO_DB"`
	User     string `mapstructure:"MONGO_USER"`
	Password string `mapstructure:"MONGO_PASSWORD"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"REDIS_ADDR"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type KafkaConfig struct {
	Brokers  string `mapstructure:"KAFKA_BROKERS"`
	GroupID  string `mapstructure:"KAFKA_GROUP_ID"`
	Username string `mapstructure:"KAFKA_USERNAME"`
	Password string `mapstructure:"KAFKA_PASSWORD"`
}

type ServerConfig struct {
	Port         int    `mapstructure:"APP_PORT"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	Environment  string `mapstructure:"ENVIRONMENT"`
	IsLive       bool   `mapstructure:"IS_LIVE"`
	Smtp         struct {
		Host     string `mapstructure:"SMTP_HOST"`
		Port     int    `mapstructure:"SMTP_PORT"`
		Email    string `mapstructure:"SMTP_EMAIL"`
		Password string `mapstructure:"SMTP_PASSWORD"`
	}
}

type AuthConfig struct {
	Tokens struct {
		AccessPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
		AccessPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
		RefreshPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
		RefreshPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	}
	AccessTokenExpired  time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED"`
	RefreshTokenExpired time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED"`

	AccessPublicKey   *rsa.PublicKey
	AccessPrivateKey  *rsa.PrivateKey
	RefreshPrivateKey *rsa.PrivateKey
	RefreshPublicKey  *rsa.PublicKey
}

func LoadEnv(fileName string) (env AppEnv, err error) {

	viper.SetConfigName(fileName)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	configs := []interface{}{
		&env.MySQL.DNS,
		&env.MySQL.DB,
		&env.Mongo,
		&env.Redis,
		&env.Kafka,
		&env.Server,
		&env.Server.Smtp,
		&env.Auth,
		&env.Auth.Tokens,
	}
	for _, config := range configs {
		err = viper.Unmarshal(config)
		if err != nil {
			return
		}
	}

	return
}
