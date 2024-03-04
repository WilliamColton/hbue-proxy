package aes

import (
	"crypto/aes"
	cipher2 "crypto/cipher"
	"github.com/WilliamColton/hbue-proxy/pkg/crypto"
)

const blockSize = 16

type AESCipher struct {
	key []byte
	iv  []byte
}

func NewAESCipher(key, iv []byte) crypto.Cipher {
	return &AESCipher{
		key: key,
		iv:  iv,
	}
}

func (a *AESCipher) Encode(buf []byte) ([]byte, error) {
	block := crypto.PKCS7Padding(buf, blockSize)
	cipher, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	CBCEncrypter := cipher2.NewCBCEncrypter(cipher, a.iv)
	CBCEncrypter.CryptBlocks(block, block)
	return block, nil
}

func (a *AESCipher) Decode(buf []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	CBCDecrypter := cipher2.NewCBCDecrypter(cipher, a.iv)
	CBCDecrypter.CryptBlocks(buf, buf)
	return crypto.UnPKCS7Padding(buf), nil
}
