package mtproto

import (
	"net"
	"strconv"
)

var DefaultDialer = NewTcpDialer()

type Dialer interface {
	Dial(address string) (net.Conn, error)
	IsRoutable(address string, port int) bool
}

type tcpDialer struct{}

func NewTcpDialer() Dialer {
	return tcpDialer{}
}

func (i tcpDialer) Dial(address string) (net.Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (i tcpDialer) IsRoutable(address string, port int) bool {
	_, err := net.ResolveTCPAddr("tcp", address+":"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	return true
}

type FuncDialer struct {
	connect    func(address string) (net.Conn, error)
	isRoutable func(address string, port int) bool
}

func NewFuncDialer(
	connect func(address string) (net.Conn, error),
) FuncDialer {
	return FuncDialer{
		connect: connect,
	}
}

func (t FuncDialer) WithRoutableCheck(isRoutable func(address string, port int) bool) FuncDialer {
	t.isRoutable = isRoutable
	return t
}

func (t FuncDialer) Dial(address string) (net.Conn, error) {
	return t.connect(address)
}

func (t FuncDialer) IsRoutable(address string, port int) bool {
	if t.isRoutable == nil {
		return true
	}
	return t.isRoutable(address, port)
}
