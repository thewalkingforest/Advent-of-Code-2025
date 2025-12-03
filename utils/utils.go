package utils

import (
	"cmp"
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

func ParseInt(i string) int {
	s, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		log.Fatalf("failed to ingest: %s", err)
	}
	return int(s)
}

func MaxIndex[S ~[]E, E cmp.Ordered](x S) (E, int) {
	var max E
	if len(x) == 0 {
		return max, -1
	}
	max = x[0]
	idx := 0
	for i, e := range x {
		if cmp.Less(max, e) {
			max = e
			idx = i
		}
	}
	return max, idx
}
