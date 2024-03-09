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
	Domain = "Domain"
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
	case Domain:
		return string(address), nil
	default:
		return "", errors.New("exit directly")
	}
}

func GetAddress(addr string, port string) string {
	return fmt.Sprintf("%v:%v", addr, port)
}

func Handle(s conn.CipherConn) (net.Conn, error) {
	buf := make([]byte, 512)
	if _, err := s.DecodeRead(buf); err != nil {
		return nil, err
	}
	if _, err := s.EncodeWrite([]byte{0x05, 0x00}); err != nil {
		return nil, err
	}

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
		addr, err = GetAddr(Domain, buf[5:n-2])
		if err != nil {
			return nil, err
		}
	case 0x04:
		return nil, errors.New("does not support IPv6")
	default:
		return nil, errors.New("unknown error")
	}
	port := GetPort(buf[n-2:])
	address := GetAddress(addr, port)

	serverConn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	if _, err := s.EncodeWrite([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}); err != nil {
		return nil, err
	}

	return serverConn, err
}
