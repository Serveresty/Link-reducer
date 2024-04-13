package pkg

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateLink(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(charset[r.Intn(len(charset))])
	}
	return builder.String()
}
