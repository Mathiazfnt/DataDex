package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	ClientID     string
	ClientSecret string
	BaseURL      string
	TokenURL     string
	Namespace    string
	Locale       string
	Port         string
	CacheTTL     time.Duration
}

func LoadEnv(path string) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineNumber int = 0

	for scanner.Scan() {
		lineNumber++
		var line string = scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		var parts []string = strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			return fmt.Errorf("Ligne invalide à la ligne %d", lineNumber)
		}

		var key string = strings.TrimSpace(parts[0])
		var value string = strings.TrimSpace(parts[1])

		if key == "" {
			return fmt.Errorf("Clé invalide à la ligne %d", lineNumber)
		}

		_, exists := os.LookupEnv(key)

		if !exists {
			err = os.Setenv(key, value)

			if err != nil {
				return err
			}
		}
	}

	err = scanner.Err()

	if err != nil {
		return err
	}

	return nil
}
