package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the first number")
	num1, _ := reader.ReadString('\n')
	num1 = strings.Replace(num1, "\n", "", -1)
	IntNum1, _ := strconv.Atoi(num1)
	fmt.Println("Please enter the second number")
	num2, _ := reader.ReadString('\n')
	num2 = strings.Replace(num2, "\n", "", -1)
	IntNum2, _ := strconv.Atoi(num2)
	fmt.Println("Please enter the second number")
	num3, _ := reader.ReadString('\n')
	num3 = strings.Replace(num3, "\n", "", -1)
	IntNum3, _ := strconv.Atoi(num3)
	average := (IntNum1 + IntNum2 + IntNum3) / 3
	fmt.Printf("the average is %v \n", average)
}
