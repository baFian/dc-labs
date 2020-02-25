// Clock2 is a concurrent TCP server that periodically writes the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
	"os"
	"strings"
	
)
func TimeIn(t time.Time, name string) (time.Time, error) {
    loc, err := time.LoadLocation(name)
    if err == nil {
        t = t.In(loc)
    }
    return t, err
}

func handleConn(zone string,c net.Conn) {
	defer c.Close()
	for {
		
		realTime,_ := TimeIn(time.Now(),zone)
		_, err := io.WriteString(c,zone+"        :"+realTime.Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	input := os.Args[1]

	port := os.Args[4]
	
	var timeZone []string
	timeZone=strings.Split(input,"=")

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(timeZone[1], conn) // handle connections concurrently
	}
}
