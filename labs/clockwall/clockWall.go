package main

import(
	"io"
	"log"
	"net"
	"fmt"
	"os"
	"strings"
)

func mustCopy(dst io.Writer, src io.Reader,c chan int) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}

func main() {
	clock1 := os.Args[1]
	clock2 := os.Args[2]
	clock3 := os.Args[3]

	var(
		port []string
		port2 []string
		port3 []string
	)

	port=strings.Split(clock1,":")
	port2=strings.Split(clock2,":")
	port3=strings.Split(clock3,":")
	
	
	c:=make(chan int)

	conn, err := net.Dial("tcp", "localhost:"+port[1])
	conn1, err := net.Dial("tcp", "localhost:"+port2[1])
	conn2, err := net.Dial("tcp", "localhost:"+port3[1])
    if err != nil {
        log.Fatal(err)
    }
	defer conn.Close()
	go mustCopy(os.Stdout,conn, c)
	go mustCopy(os.Stdout,conn1, c)
	go mustCopy(os.Stdout,conn2, c)
	result :=<-c
	fmt.Print(result)
	
}