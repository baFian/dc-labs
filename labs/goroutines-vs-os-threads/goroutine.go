package main

import(
	"log"
	"fmt"
	"os"
	"strconv"
	"time"
)

const numGoroutines = 1300000

func main() {

	greenLight := make(chan int)
	final := greenLight

	for i := 0; i < numGoroutines; i++ {
		end := make(chan int)
		go gorut(final, end)
		final = end
		end = make(chan int)
	}
	
	timer := time.Now()
	greenLight <- 2
	total :=time.Since(timer)
	fmt.Println(total)

}

func gorut(input chan int, output chan int) {
	num := <-input
	num ++
	f, err := os.OpenFile("report.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(strconv.Itoa(num))); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }

	output <- num
}