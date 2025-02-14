package shortCode

import "math/rand"

// GenerateShortCode() string

type ShortCode struct {
	length int
}

func NewShortCode(length int) *ShortCode {
	return &ShortCode{length}
}

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (s *ShortCode) GenerateShortCode() string {
	length := len(chars)
	result := make([]byte, s.length)
	for i := 0; i < s.length; i++ {
		result[i] = chars[rand.Intn(length)]
	}
	return string(result)
}
