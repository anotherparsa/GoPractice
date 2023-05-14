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
	fmt.Println("Please enter the first number:")
	Num1, _ := reader.ReadString('\n')
	Num1 = strings.Replace(Num1, "\n", "", -1)
	IntNum1, _ := strconv.Atoi(Num1)
	fmt.Println("Please enter the second number:")
	Num2, _ := reader.ReadString('\n')
	Num2 = strings.Replace(Num2, "\n", "", -1)
	IntNum2, _ := strconv.Atoi(Num2)
	fmt.Println("Please enter the third number:")
	Num3, _ := reader.ReadString('\n')
	Num3 = strings.Replace(Num3, "\n", "", -1)
	IntNum3, _ := strconv.Atoi(Num3)
	Max := IntNum1
	if IntNum2 > Max {
		Max = IntNum2
	} else if IntNum3 > Max {
		Max = IntNum3
	}
	Min := IntNum1
	if IntNum2 < Min {
		Min = IntNum2
	} else if IntNum3 < Min {
		Min = IntNum3
	}
	fmt.Printf("the Highest number is %v and the Lowest number is %v\n", Max, Min)

}
