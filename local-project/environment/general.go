package environment

import "os"

var (
	Email         = getEnv("EMAIL", "")
	EmailPassword = getEnv("EMAIL_PASSWORD", "")
	ToEmail       = getEnv("TO_EMAIL", "")
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
