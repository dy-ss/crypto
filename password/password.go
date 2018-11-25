package password

import (
	"golang.org/x/crypto/argon2"
	"math/rand"
)


var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

// generates a random string of fixed size
func srand(size int) []byte {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return buf
}

func Make(password []byte) (key, salt []byte){
	salt = srand(10)
	key = argon2.Key(password, salt, 3, 32*1024, 4, 32)
	return
}
