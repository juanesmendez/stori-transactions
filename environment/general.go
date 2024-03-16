package environment

import "os"

var (
	Email         = getEnv("EMAIL", "")
	EmailPassword = getEnv("EMAIL_PASSWORD", "")
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
