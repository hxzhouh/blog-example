package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

type Connection struct {
	*tls.Conn
}

func main() {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		MaxVersion:         tls.VersionTLS12,
	}

	conn, err := net.Dial("tcp4", "httpbin.org:443")
	if err != nil {
		panic(err)
	}

	tlsConn := tls.Client(&TimedConn{Conn: conn}, tlsConfig)
	tlsConn.Handshake()
	fmt.Println("TLS handshake complete")

	buff := make([]byte, 1500)
	tlsConn.Write([]byte("GET /get HTTP/1.1\r\nHost:httpbin.org\r\n\r\n"))
	n, _ := tlsConn.Read(buff)
	fmt.Println(string(buff[:n]))
}

// TimedConn 实现 net.Conn 接口，并在每次读写时记录时间
type TimedConn struct {
	Conn      net.Conn
	lastRead  time.Time
	lastWrite time.Time
}

func (t *TimedConn) Read(b []byte) (n int, err error) {
	startTime := time.Now()
	fmt.Println(startTime, "net.Conn start reading")
	n, err = t.Conn.Read(b)
	t.lastRead = time.Now()
	fmt.Println(time.Now(), "net.Conn last read", time.Since(startTime))
	//fmt.Println(hex.Dump(b[:n]))
	return
}

func (t *TimedConn) Write(b []byte) (n int, err error) {
	startTime := time.Now()
	fmt.Println(startTime, "net.Conn start writing")
	n, err = t.Conn.Write(b)
	t.lastWrite = time.Now()
	fmt.Println(time.Now(), "net.Conn last write", time.Since(startTime))
	//fmt.Println(hex.Dump(b[:n]))
	return
}

func (t *TimedConn) Close() error {
	return t.Conn.Close()
}

func (t *TimedConn) LocalAddr() net.Addr {
	return t.Conn.LocalAddr()
}

func (t *TimedConn) RemoteAddr() net.Addr {
	return t.Conn.RemoteAddr()
}

func (t *TimedConn) SetDeadline(td time.Time) error {
	return t.Conn.SetDeadline(td)
}

func (t *TimedConn) SetReadDeadline(td time.Time) error {
	return t.Conn.SetReadDeadline(td)
}

func (t *TimedConn) SetWriteDeadline(td time.Time) error {
	return t.Conn.SetWriteDeadline(td)
}
