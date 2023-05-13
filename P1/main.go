package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter the Length:")
	reader := bufio.NewReader(os.Stdin)
	length, _ := reader.ReadString('\n')
	length = strings.Replace(length, "\n", "", -1)
	IntLength, _ := strconv.Atoi(length)
	fmt.Println("Please enter the width:")
	width, _ := reader.ReadString('\n')
	width = strings.Replace(width, "\n", "", -1)
	Intwidth, _ := strconv.Atoi(width)
	area := IntLength * Intwidth
	perimeter := (Intwidth + IntLength) * 2
	fmt.Printf("the area is %v. \nand the perimeter is %v. \n", area, perimeter)

}
