//
// 加密解密相关函数
//

package cryptoutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"log"
)

// 使用aes算法,加密指定内容
func EncryptAes(rawData, key []byte) ([]byte, error) {
	if (len(rawData) == 0) || (len(key) == 0) {
		return nil, errors.New("encrypt data or secret key is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	padding := aes.BlockSize - len(rawData)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	rawData = append(rawData, padtext...)

	blockMode := cipher.NewCBCEncrypter(block, key[:aes.BlockSize])
	// CryptBlocks(dst, src []byte), Dst and src may point to the same memory.
	blockMode.CryptBlocks(rawData, rawData)
	return rawData, nil
}

// 使用aes算法,解密指定内容
func DecryptAes(cryptData, key []byte) ([]byte, error) {
	if (len(cryptData) == 0) || (len(key) == 0) {
		return nil, errors.New("decrypt data or secret key is empty")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])

	rawData := make([]byte, len(cryptData))
	// CryptBlocks(dst, src []byte), Dst and src may point to the same memory.
	blockMode.CryptBlocks(rawData, cryptData)

	unpadding := int(rawData[len(rawData)-1])

	if len(rawData) > unpadding {
		rawData = rawData[:(len(rawData) - unpadding)]
	} else {
		return nil, errors.New("decrypt raw data error, key is correct?")
	}

	return rawData, nil
}
