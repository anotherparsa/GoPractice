package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("this is the file's content before changing :")
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal("something went wrong while opening")
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	fmt.Println("now we're going to change the file's content")
	file, err = os.OpenFile("test.txt", 2, 0777)
	if err != nil {
		log.Fatal("something went wrong while opening")
	} else {
		_, err := file.WriteString("text message after change")
		if err != nil {
			log.Fatal("something went wrong while writing")
		} else {
			fmt.Println("the file's content has been changed")
		}
	}
}
