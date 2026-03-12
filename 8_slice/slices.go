//dynamic arrays
//most used

package main

import "fmt"

func main() {
	//declared not initialized,if empty is nil
	var nums []int
	var notnil = make([]int, 10, 100) //100 represent is capacity,10 is pre fild data
	fmt.Println(notnil, nums)         //[000000000],[]
	// fmt.Println(notnil[11])
	fmt.Println(cap(notnil)) //10
	notnil = append(notnil, 19)
	fmt.Println(notnil) //on last position add 19,[0 0 0 0 0 0 0 0 0 0 19]
	fmt.Println(append(notnil, 19))

	s := []int{2, 4, 6, 8, 10, 12, 14}

	fmt.Println(s[1:2])
	fmt.Println(s[1:4])
	fmt.Println(s[1:1])
	fmt.Println(s[1:])
	fmt.Println(s[:])
	fmt.Print(s[:2])

}
