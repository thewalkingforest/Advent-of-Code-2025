package utils

import (
	"log"
	"os"
	"strconv"
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

func ParseI64(i string) int64 {
	s, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	return int64(s)
}

func ParseI32(i string) int32 {
	s, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	return int32(s)
}
