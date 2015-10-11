package main
import "net"
import "fmt"
import (
	"bufio"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:8082")				// connect to this socket
	i:=0
	for ; i < 9000000; i++ {
		text:="Hello"
		fmt.Fprintf(conn, text + "\n")
		//listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')	//write message to socket
		_=message
	}
	fmt.Print("Connection Closed \n")
	conn.Close()

}
