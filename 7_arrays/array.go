package main

import "fmt"

func main() {
	//in go array size is fixed if you give value [] in this bracket
	var nums [10]int
	fmt.Println(len(nums)) //10

	//how to declare string array
	fruits := [3]string{"apple", "banana", "mango"}
	fmt.Println(fruits)    //[apple banana mango]
	fmt.Println(fruits[1]) //banana

	//ellipsis (...) se -compile khud size count karta hai

	numsElipsis := [...]int{10, 10, 20, 30, 40, 50}
	fmt.Println(numsElipsis)      //[10 10 20 30 40 50]
	fmt.Println(len(numsElipsis)) //6

	//specific index pe value set karna

	arrs := [5]int{0: 100, 4: 500}
	fmt.Println(arrs) //[100,0,0,0,500]

	//Go mein arrays directly == se compare ho sakte hain:
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	c := [3]int{1, 3, 2}

	fmt.Println(a == c) //false
	fmt.Println(a == b) //true

	//Array Copy — Value Type Hai!
	//Go mein array value type hai, matlab jab assign karo toh copy banta hai, reference nahi.

	noncopy := [4]int{1, 3, 4, 5}
	copyarray := noncopy

	fmt.Println(copyarray) //[1,3,4,5]
	copyarray[1] = 500
	fmt.Println(copyarray) //[1,500,4,5]

	//add element
	nums[9] = 1
	for index, item := range nums {
		fmt.Println(item, index) //0.0,......1
	}

}
