package main

import "fmt"

func main() { //entry point of any go program
	var name string
	var age int

	fmt.Print("Enter You Name")
	fmt.Scan(&name)
	fmt.Print("Enter You Age")
	fmt.Scan(&age)
	fmt.Print("Hello ", name, " You age is ", age)

}
