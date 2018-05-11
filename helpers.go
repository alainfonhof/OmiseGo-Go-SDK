package omisego

import (
	"crypto/rand"
	"fmt"
)

// newIdempotencyToken returns a randomly generated idempotency token.
func NewIdempotencyToken() string {
	b := make([]byte, 16)
	rand.Reader.Read(b)
	return UUIDVersion4(b)
}

// uUIDVersion4 returns a Version 4 random UUID from the byte slice provided
func UUIDVersion4(u []byte) string {
	// https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29
	// 13th character is "4"
	u[6] = (u[6] | 0x40) & 0x4F
	// 17th character is "8", "9", "a", or "b"
	u[8] = (u[8] | 0x80) & 0xBF

	return fmt.Sprintf(`%X-%X-%X-%X-%X`, u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}
