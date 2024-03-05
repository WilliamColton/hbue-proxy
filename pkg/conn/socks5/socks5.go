package socks5

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/WilliamColton/hbue-proxy/pkg/conn"
	"net"
	"strconv"
)

const (
	IPv4   = "IPv4"
	IPv6   = "IPv6"
	Domaim = "Domain"
)

func GetPort(portData []byte) string {
	return strconv.Itoa(int(binary.BigEndian.Uint16(portData)))
}

func GetAddr(dataType string, address []byte) (string, error) {
	switch dataType {
	case IPv4:
		var ip net.IP
		ip = address
		return ip.String(), nil
	case Domaim:
		return string(address), nil
	default:
		return "", errors.New("Exit directly.")
	}
}

func GetAddress(addr string, port string) string {
	return fmt.Sprintf("%v:%v", addr, port)
}

func Socks5Handle(s conn.CipherConn) (net.Conn, error) {
	buf := make([]byte, 512)
	s.DecodeRead(buf)
	s.EncodeWrite([]byte{0x05, 0x00})

	n, _ := s.DecodeRead(buf)
	var addr string
	var err error
	switch buf[3] {
	case 0x01:
		addr, err = GetAddr(IPv4, buf[4:4+net.IPv4len])
		if err != nil {
			return nil, err
		}
	case 0x03:
		addr, err = GetAddr(Domaim, buf[5:n-2])
		if err != nil {
			return nil, err
		}
	case 0x04:
		return nil, errors.New("Does not support IPv6.")
	default:
		return nil, errors.New("Unknown error.")
	}
	port := GetPort(buf[n-2:])
	address := GetAddress(addr, port)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	s.EncodeWrite([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	return conn, err
}
