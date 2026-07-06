package viper

import (
	"os"
	"strings"
)

func parseEnv(env string) (string, string) {
	parts := strings.SplitN(env, "=", 2)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}