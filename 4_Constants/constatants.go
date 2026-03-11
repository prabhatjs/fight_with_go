package main

import "fmt"

const isStudent = false

func main() {
	const onlyUser = "Prabhat"
	const age = 30
	// isStudent=true give error,const value will not chnage in future
	fmt.Println(isStudent)

	//hots multiple const group
	const (
		port = 5000
		host = "localhost:8080"
	)

	fmt.Println(port, host)
}
