package main

import (
	"fmt"
	"time"
)

func main() {

	// var CurrentTime time.Time

	// CurrentTime = time.Now()
	CurrentTime := time.Now() //you dont want to define variable type
	fmt.Println(CurrentTime)
	fmt.Println("Year", CurrentTime.Year())
	fmt.Println("Day", CurrentTime.Day())
	fmt.Println("MONTH", CurrentTime.Month())
	fmt.Println("HOUR", CurrentTime.Hour())
	fmt.Println("MINUT", CurrentTime.Minute())
	fmt.Println("SECOND", CurrentTime.Second())

	//Sleep function progam will close after 3 second--Yeh time ka predefined duration constant hai.
	time.Sleep(3 * time.Second)         //line by line execute hoga 1,agr upar niche sequnce krna hai to goroutines use kro
	time.Sleep(200 * time.Nanosecond)   //2
	time.Sleep(3000 * time.Millisecond) //3
	time.Sleep(4000 * time.Microsecond) //4
}
