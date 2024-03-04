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
