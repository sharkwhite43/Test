package main

import (
	"fmt"
	"time"
)

func say(txt string, sleep time.Duration, wg *sync.WaitGrooup) {
	defer wg.Done()
	fmt.Println(txt)
	time.Sleep(time.Millisecond * sleep)

}
