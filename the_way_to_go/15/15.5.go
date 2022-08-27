package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage:host port")
	}
	hostAndPord := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPord)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHaldner(conn)
	}
}
func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving addr:port failed")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP:")
	println("Listening to :" + listener.Addr().String())
	return listener
}

func connectionHaldner(conn net.Conn) {
	connForm := conn.RemoteAddr().String()
	println("Connection from :", connForm)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		lengh, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(lengh, err, ibuf)
		case syscall.EAGAIN:
			continue
		default:
			goto DISCONNECT

		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection:", connForm)
	checkError(err, "Close:")
}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	checkError(err, "Write:wrote "+string(wrote)+"bytes.")
}

func handleMsg(lengh int, err error, ibuf []byte) {
	if lengh > 0 {
		print("<", lengh, ":")
		for i := 0; ; i++ {
			if ibuf[i] == 0 {
				break
			}
			fmt.Printf("%c", ibuf[i])
		}
		print(">")
	}
}
func checkError(err error, info string) {
	if err != nil {
		panic("Error:" + info + " " + err.Error())
	}
}
