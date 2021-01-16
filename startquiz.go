package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
)

//Global variable for debug flag
var debugFlag int

func debugPrint(text string) {
	if debugFlag == 1 {
		fmt.Println(text)
	}
}

func main() {

	funcargs := os.Args

	if len(funcargs) > 1 {
		var argserr error
		debugFlag, argserr = strconv.Atoi(funcargs[1])
		if argserr == nil {
			if debugFlag == 1 {
				fmt.Println("Debugging enabled")
			}
		} else {
			log.Fatal("Error getting debug flag")
		}
	}




	//Main game loop
	fmt.Println("Welcome to the piano quiz! Please select from the menu")

	var selection string
	var err error

	selection, err = displayMenu()

	if err != nil {
		fmt.Println("error!")
	}
	
	fmt.Printf("You selected %v\n", selection)

}

func displayMenu () (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Main menu:")
	fmt.Println("1. New game")
	fmt.Println("0. Exit")
	selection, _ := reader.ReadString('\n')
	debugPrint("Current selection: " + selection)

	return selection, nil		


}