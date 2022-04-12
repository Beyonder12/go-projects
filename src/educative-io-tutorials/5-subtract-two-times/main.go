// package main

// import "fmt"
// import "time"

// func main() {
//     time := time.Now()
//     time1:=time.AddDate(0,0,1)
//     time2:=time.AddDate(0,0,2)
//     diff := time2.Sub(time1)
//     fmt.Println("Current date and time is: ", diff.String())

// }

package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Now()
	time1 := time.AddDate(0, 0, 1)
	time2 := time.AddDate(0, 0, 2)
	diff := time2.Sub(time1)
	fmt.Println("Current date and time is ", diff.String())
}
