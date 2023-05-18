package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	Table := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	fmt.Println("Please enter the number :")
	reader := bufio.NewReader(os.Stdin)
	Input, _ := reader.ReadString('\n')
	Input = strings.Replace(Input, "\n", "", -1)
	Number := 0
	for i := 0; i < len(Input); i++ {
		switch rune(Input[i]) {
		case 'I':
			Number += Table['I']
		case 'V':
			Number += Table['V']
		case 'X':
			Number += Table['X']
		case 'L':
			Number += Table['L']
		case 'C':
			Number += Table['C']
		case 'D':
			Number += Table['D']
		case 'M':
			Number += Table['M']
		default:
			log.Fatal("INVALID INPUT")
		}
	}
	fmt.Printf("the number is %v \n", Number)
}
