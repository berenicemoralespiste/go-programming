package main

import "fmt"

func main() {

	/*This is a unbuffered channel, we need a go routine to send or recive, because the only go routine
	is main
	*/

	//Go routine main sent 42 to channel c
	c := make(chan int)
	/*A way to avoid one goroutine more is have a buffer channel
	that gives us the chance of have data.
	*/
	//c:= make(chan int,1)
	c <- 42

	//Anonymous goroutine receives the value of we sent on channel "c"
	go func() {
		fmt.Println(<-c)
	}()

}
