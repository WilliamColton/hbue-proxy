package crypto

import "bytes"

type Cipher interface {
	Encode([]byte) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

func PKCS7Padding(buf []byte, blockSize int) []byte {
	paddingSize := blockSize - len(buf)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	return append(buf, padding...)
}

func UnPKCS7Padding(buf []byte) []byte {
	paddingSize := buf[len(buf)-1]
	return buf[:len(buf)-int(paddingSize)]
}

func KeyPadding(key []byte) []byte {
	keyLen := len(key)
	paddingLen := 16 - keyLen

	if keyLen >= 32 {
		return key[:32]
	} else if keyLen >= 24 {
		return key[:24]
	} else if keyLen >= 16 {
		return key[:16]
	} else {
		return append(key, bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)...)
	}
}
