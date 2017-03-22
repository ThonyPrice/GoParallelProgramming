/*
Task 3 - Alarm clock
DD1396 - Parallel and concurrent programming @ KTH
Author: Thony Price
Last revision: 2017-03-22
*/

package main

import (
	"fmt"
	"time"
)

// Remind takes a string and a delay (time) and prints various messages
// at intervals
func Remind(text string, delay time.Duration) {

	var summerTime time.Duration = 3600

	food := time.NewTicker(5 * time.Hour)
	work := time.NewTicker(8 * time.Hour)
	sleep := time.NewTicker(24 * time.Hour)

	for true {
		timeNow := time.Now().Add(summerTime).Format("15.04")
		select {
		case <-food.C:
			fmt.Printf("%v %v: Dags att ata\n", text, timeNow)
		case <-work.C:
			fmt.Printf("%v %v: Dags att arbeta\n", text, timeNow)
		case <-sleep.C:
			fmt.Printf("%v %v: Dags att sova\n", text, timeNow)
		default:
			fmt.Printf("%v %v\n", text, timeNow)
		}
		time.Sleep(time.Second * delay)
	}
}

func main() {
	Remind("Klockan är", 5)
}
