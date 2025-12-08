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

func ParseInt64(i string) int64 {
	s, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		log.Fatalf("failed to parse int64: %s", err)
	}
	return int64(s)
}

func ParseInt32(i string) int32 {
	s, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse int32: %s", err)
	}
	return int32(s)
}

func ParseInt(i string) int {
	s, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		log.Fatalf("failed to parse int: %s", err)
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

func IndexAll(s string, substr string) []int {
	var idxs []int

	i := 0
	found := strings.Index(s, substr)
	for found != -1 && i < len(s) {
		i += found
		idxs = append(idxs, i)
		i += len(substr)
		found = strings.Index(s[i:], substr)
	}

	return idxs
}
