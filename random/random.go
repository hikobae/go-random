package random

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

const alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Init initializes random.
func Init() {
	var seed int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &seed); err != nil {
		panic(err)
	}
	rand.Seed(seed)
}

// Alphabetn returns an string that consists only alphabet.
func Alphabetn(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = alphabets[rand.Intn(len(alphabets))]
	}
	return string(b)
}

// ShuffleSlice shuffles slice with Fisher–Yates shuffle.
func ShuffleSlice(p []int) {
	n := len(p)
	if n <= 1 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		p[i], p[j] = p[j], p[i]
	}
}
