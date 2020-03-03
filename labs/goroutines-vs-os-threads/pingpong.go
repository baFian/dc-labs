package main

import(
	"fmt"
	"time"
	"strconv"
	"os"
)
/*var value=0
var sum=0*/

func main(){

	ping :=make(chan int)
	pong :=make(chan int)

	go send(ping, pong)
	go send(pong, ping)

	seconds, _ := strconv.Atoi(os.Args[1])
	fmt.Println(seconds)
	ping <-1
	fmt.Println("Test:",seconds,"seconds")
	time.Sleep(time.Duration(seconds)*time.Second)
	
	select {

		case num := <-ping:
			fmt.Println("communication:", num)
		case num := <-pong:
			fmt.Println("communication:", num)
	}

	close(ping)
	close(pong)


}

func send(input chan int, output chan int){

	for score := range input {
		score++
		output <- score
	}

}