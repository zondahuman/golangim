//粘包问题演示服务端
package main

import (
	"fmt"
	"golangim/im/handler"
	"net"
	"os"
)

func main() {
	netListen, err := net.Listen("tcp", ":9988")
	CheckError(err)

	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data length:", n)
		Log(conn.RemoteAddr().String(), "receive data:", buffer[:n])
		Log(conn.RemoteAddr().String(), "receive data string:", string(buffer[:n]))
		go handler.ProcessHandler(string(buffer[:n]))
		var result = "receive success"
		conn.Write([]byte(result))
		conn.Close()
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
