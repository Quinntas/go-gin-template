package envLoader

import (
	"github.com/joho/godotenv"
	"os"
)

func getEnv(key string, required bool) string {
	value := os.Getenv(key)
	if required && value == "" {
		panic("Missing required environment variable: " + key)
	}
	return value
}

type Variables struct {
	Port        string
	DatabaseURL string
	RedisURL    string
	Pepper      string
	JwtSecret   string
}

func NewVariables() (*Variables, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Variables{
		Port:        getEnv("PORT", true),
		DatabaseURL: getEnv("DATABASE_URL", true),
		RedisURL:    getEnv("REDIS_URL", true),
		Pepper:      getEnv("PEPPER", true),
		JwtSecret:   getEnv("JWT_SECRET", true),
	}, nil
}
