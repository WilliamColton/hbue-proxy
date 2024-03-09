package aes

import (
	"crypto/aes"
	cipherpkg "crypto/cipher"

	"github.com/WilliamColton/hbue-proxy/pkg/crypto"
)

const blockSize = 16

type Cipher struct {
	key []byte
	iv  []byte
}

func NewAESCipher(key, iv []byte) crypto.Cipher {
	return &Cipher{
		key: key,
		iv:  iv,
	}
}

func (a *Cipher) Encode(buf []byte) ([]byte, error) {
	block := crypto.PKCS7Padding(buf, blockSize)
	cipher, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	CBCEncrypt := cipherpkg.NewCBCEncrypter(cipher, a.iv)
	CBCEncrypt.CryptBlocks(block, block)
	return block, nil
}

func (a *Cipher) Decode(buf []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, err
	}
	CBCDecrypt := cipherpkg.NewCBCDecrypter(cipher, a.iv)
	CBCDecrypt.CryptBlocks(buf, buf)
	return crypto.UnPKCS7Padding(buf), nil
}
