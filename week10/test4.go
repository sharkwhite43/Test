package main

import "time"

func say(txt string, sleep time.Duration, wg *sync.WaitGrooup) {
	defer wg.Done()

}
