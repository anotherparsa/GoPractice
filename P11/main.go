package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Please enter the text")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("something went wrong")
	} else {
		text = strings.Replace(text, "\n", "", -1)
	}
	for i := 0; i < len(text); i++ {
		if unicode.IsLetter(rune(text[i])) {
			fmt.Println("it's a letter")
		} else {
			fmt.Println("it's a number")
		}
	}
}
