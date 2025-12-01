package utils

import (
	"os"
	"strings"
)

func Ingest(name string) (string, error) {
	buff, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	input := strings.TrimSpace(string(buff))
	return input, nil
}
