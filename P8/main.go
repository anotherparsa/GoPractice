package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter the N:")
	reader := bufio.NewReader(os.Stdin)
	Num, _ := reader.ReadString('\n')
	Num = strings.Replace(Num, "\n", "", -1)
	IntNum, _ := strconv.Atoi(Num)
	if IntNum%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
