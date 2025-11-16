package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	ServerHost      string
	ServerPort      string
	ServerMode      string
	ServerSecretKey string
	MySQLHost       string
	MySQLPort       string
	MySQLUser       string
	MySQLPassword   string
	MysqlDatabase   string
	RedisHost       string
	RedisPort       string
	RedisPassword   string
}

var (
	SERVER_CONFIG           *Config
	onceConfigInitilization sync.Once
)

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, using system environment variables: %v", err)
	}

	config := &Config{
		ServerHost:      getEnv("SERVER_HOST", "0.0.0.0"),
		ServerPort:      getEnv("SERVER_PORT", "4387"),
		ServerMode:      getEnv("SERVER_MODE", "Debug"),
		ServerSecretKey: getEnv("SERVER_SECRET_KEY", "12345678"),
		MySQLHost:       getEnv("MYSQL_HOST", "127.0.0.1"),
		MySQLPort:       getEnv("MYSQL_PORT", "3306"),
		MySQLUser:       getEnv("MYSQL_USER", "root"),
		MySQLPassword:   getEnv("MYSQL_PASSWORD", "12345678"),
		MysqlDatabase:   getEnv("MYSQL_DATABASE", "gin-web-template"),
		RedisHost:       getEnv("REDIS_HOST", "127.0.0.1"),
		RedisPort:       getEnv("REDIS_PORT", "6379"),
		RedisPassword:   getEnv("REDIS_PASSWORD", "12345678"),
	}

	return config, nil
}

// getEnv retrieves environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func (config *Config) GetBindAddr() string {
	return fmt.Sprintf("%s:%s", config.MySQLHost, config.ServerPort)
}

func (config *Config) GetMysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLHost,
		config.MySQLPort,
		config.MysqlDatabase)
}

func (config *Config) GetRedisAddr() string {
	return fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)
}

func (config *Config) GetRedisPassword() string {
	return config.RedisPassword
}

func InitConfig() error {
	var err error
	onceConfigInitilization.Do(func() {
		SERVER_CONFIG, err = LoadConfig()
		if err != nil {
			log.Println("Unable to load the config.")
		}
	})
	return err
}

func ServerConfig() (*Config, error) {
	if SERVER_CONFIG == nil {
		log.Println("Server config has not been initialized!")
		return nil, errors.New("server config has not been initialized")
	}
	return SERVER_CONFIG, nil
}
