package main

import (
	"fmt"
	"time"
)

// Main function is also a kind of go routine.
func main() {
	start := time.Now() // func to store the current time in var start
	defer func() {
		elapsed := time.Since(start)
		fmt.Println(elapsed)
	}() //IIF to print the elapsed time since the var start stored the time when main func started the execution.

	fmt.Println("welcome to the demonstration for Go channels.")

	myChannel := make(chan string) //Making a channel which will accept string type values.

	go sendData(myChannel) //Calling the func sendData to put some data into the channel.

	fmt.Println(<-myChannel) //Blocking statememt for main func, untill sendData func which is a go routine send some signal/data into the channel, this statement is called blocking because its listening for the data that will come into the channel.
}

func sendData(ch chan string) {
	time.Sleep(time.Second * 5) //Func to sleep the sendData func for 5 seconds to see if main func is really listening for the channel.
	ch <- "hello_world"         //sending signal/data to the channel
}

//generally, Go channels are used to set up communication between the go routines.

//Consider a scenario where one go routine wants to send a message and another go routine has to read that message, in these types of cases we utilize the channels.

//Channels are somewhat similar to queues, they work in FIFO manner, First In First Out, means whatever value is going first into the channel going to be read first.

//Channel can be used as blocking statement, to block the execution of main func go routine, if there's a reading statement in main func go routine, it will wait for the go routine thats outside of main func to put some data into that channel.

// go func() {
// 	myChannel <- "Hello, world!"
// }()
