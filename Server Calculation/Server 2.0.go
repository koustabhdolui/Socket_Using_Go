package main
import "net"
import "fmt"
import "bufio"
import (

	"time"
	"bytes"
)


func main() {
	fmt.Println("\nLaunching server...\n")								// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8082")								// accept connection on port
	conn, _ := ln.Accept()											// run loop forever (or until ctrl-c)
	fmt.Print("Server Started \n")
	start := time.Now()
	i:=0
	size:=0
	y:=1
	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')
		conn.Write([]byte(""+"\n"))										//Reply for ACK
		m_bytes:=[]byte(message)									//Convert to bytes
		x:=bytes.NewReader(m_bytes)
		y=x.Len();													//Calculate length in bytes
		size=size+y													//Calculating and adding byte size

		if y!=0 {													//check if empty
			i++
		} else {
			fmt.Print("Disconnected\n")
			break
		}



	}
	value := time.Since(start)
	fmt.Printf("\n\nTime Consumed is ")								//Stip timer and calculate details
	fmt.Printf("%v\n", value)
	fmt.Printf("Number of Hellos received: ")
	fmt.Printf("%v\n", i)
	fmt.Printf("Number of Bytes: ")
	fmt.Printf("%v\n", size)
	size=int(size)
	int_value:=int(value)
	int_value=int_value/1000000
	band:=(size/int_value)*8
	fmt.Printf("The bandwidth is: ")
	fmt.Printf("%v ", band)
	fmt.Printf("Kbps")
}