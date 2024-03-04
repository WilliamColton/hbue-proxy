//将具体的加密方式抽象出来，在具体的Local和Server实现中再传入加解密的对象

package conn

import (
	"net"
)
import "github.com/WilliamColton/hbue-proxy/pkg/crypto"

type CipherConn interface {
	net.Conn
	EncodeWrite([]byte) (int, error)
	DecodeRead([]byte) (int, error)
}

type SecureConn struct {
	net.Conn
	cipher crypto.Cipher
}

func NewCipherConn(c net.Conn, cipher crypto.Cipher) CipherConn {
	return &SecureConn{
		Conn:   c,
		cipher: cipher,
	}
}

func (s *SecureConn) EncodeWrite(buf []byte) (n int, err error) {
	bufLen := len(buf)
	buf, err = s.cipher.Encode(buf)
	if err != nil {
		return 0, err
	}
	_, err = s.Write(buf)
	return bufLen, err
}

func (s *SecureConn) DecodeRead(buf []byte) (n int, err error) {
	bufLen, err := s.Read(buf)
	buf, err = s.cipher.Decode(buf[:bufLen])
	return len(buf), err
}
