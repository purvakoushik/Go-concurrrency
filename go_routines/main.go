package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now() //current time storing in the start var.

	defer func() { fmt.Println(time.Since(start)) }() //defer is used to execute this IIF or lambda function immediately at last in the main function.

	// fmt.Println("Welcome to the chapter of go routine.")

	badPersons := []string{"carl", "larc", "kevin", "louis"} //Go slice

	// Loops to call the punch() func parallely for all the respective values in slice.
	for _, person := range badPersons {
		go punch(person) // Go statement is used to call the punch function as go routine.
	}
	time.Sleep(time.Second * 4) // Waiting for the go routines to finish their execution in main func, basically this func time.sleep is used here to prevent the main functions form execute out before the go routines come back.
}

// Function punch to get a person name as a string.
func punch(person string) {
	time.Sleep(time.Second * 3) //Three second wait just to test.
	fmt.Println("punched the bad person ", person)
}

//Total exec time will be approx 4 sec. because all function call will be at the same time parallely so all four calls will take approx. 3 sec. and 1 sec. for main func.
//In normal func call will like every call will take 3 sec. eac in total '3 sec * 4 func calls' 12 sec for func calls and 1 sec for main func so total time gonna be 13 sec. without multi-threading/go routine.
