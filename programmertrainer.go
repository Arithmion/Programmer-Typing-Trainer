package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
)

var amountToPrint *int = flag.Int("amount", 40, "The amount of characters to print at once. Default 40.")

func main() {
	flag.Parse()
	core()
}

func core() {
	printedString := ""
	progChars := []rune{'<', '>', '|', '^', '!', '$', '%', '&', '*', '(', ')', '\\', '?', '/', '[', ']', '{', '}', '`', ':', '+', '=', '-', ';', '_', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

	for i := 0; i < *amountToPrint; i++ {
		printedString += string(progChars[rand.Intn(len(progChars))])
	}

	fmt.Println("Copy the following text:", printedString)
	fmt.Println("Press <Enter> to submit.")
	fmt.Println("Starting in 3")
	time.Sleep(1 * time.Second)
	fmt.Println("2")
	time.Sleep(1 * time.Second)
	fmt.Println("1")
	time.Sleep(1 * time.Second)
	fmt.Println("GO!")
	marker := time.Now()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	timeTaken := time.Since(marker)
	input = strings.TrimSpace(input)
	errorAmount := levenshtein.ComputeDistance(input, printedString)
	fmt.Println("Time taken:", timeTaken, "\n# of errors:", errorAmount, "\nWPM (assuming 5 chars per word):", (float64(len(printedString)/5))/(timeTaken.Minutes()), "\nWould you like to practice again? Y/N")
	decision, _ := reader.ReadString('\n')
	decision = strings.TrimSpace(decision)
	if decision != "Y" && decision != "y" {
		fmt.Println("Exiting. Thanks for playing!")
		os.Exit(0)
	}
	core()
}
