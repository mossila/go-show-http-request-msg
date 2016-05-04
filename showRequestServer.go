package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func httpBasic(c net.Conn) {
	content := `<html>
<body>
<h1>Hello, World!</h1>
</body>
</html>
`
	output := fmt.Sprintf("%s\r\n%s\r\n%s\r\n\r\n%s\r\n",
		"HTTP/1.1 200 OK",
		"Content-Type: text/html",
		"Server: ChatServer/1.0 (OSX)",
		content,
	)

	c.Write([]byte(output))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
func clientHandler(c net.Conn) {
	defer c.Close()
	buffSize := 1024
	bufString := ""
	for {
		m := make([]byte, buffSize)
		n1, err := c.Read(m)
		checkError(err)
		bufString += fmt.Sprint(string(m))
		if n1 < buffSize { //Finish read
			fmt.Print(bufString)
			httpBasic(c)
			fmt.Printf("\n---End Request---\n")
			return
		}

	}

}

/*MockServer start chatserver on tcp with specific port
 * Because this is exported function
 * if it hos no comment here you will get warning like this
 * > exported function ChatServer should have comment or be unexported
 */
func MockServer(port string) {
	l, err := net.Listen("tcp", ":"+port)
	fmt.Printf("try localhost:%s\n", port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go clientHandler(conn)
	}
}

func main() {
	if len(os.Args) < 2 {
		println("usage")
		println("showRequestServer <port>")
		println("showRequestServer 8080")
		return
	}
	port := os.Args[1]
	MockServer(port)
	
}
