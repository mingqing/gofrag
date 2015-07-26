package randomutils

import (
	"math/rand"
	"time"
)

// 大小写字母与数字组合
func LettersDigits(length int) string {
	rand.Seed(time.Now().UnixNano())

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		keyType := rand.Intn(3)

		var keyValue byte
		switch keyType {
		case 0:
			keyValue = byte(rand.Intn(9) + '0')
		case 1:
			keyValue = byte(rand.Intn('z'-'a') + 'a')
		case 2:
			keyValue = byte(rand.Intn('Z'-'A') + 'A')
		}

		result = append(result, keyValue)
	}

	return string(result)
}
