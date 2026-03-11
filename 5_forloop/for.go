package main

import "fmt"

//for is only use in loop
func main() {
	//while loop use of for
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	//infinite loop just do for{}

	//classic loop

	for j := 20; j < 30; j++ {
		//	break//break the code from here not go to print line

		if j == 24 { //skip this line and continue work after 24
			continue
		}
		fmt.Println(j)

	}

	for K := range 5 {
		fmt.Println(K)
	}

}
