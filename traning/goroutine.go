package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloworld(&wg , "sdcard")
	go goodbye(&wg ,"sdssd")
	wg.Wait()
}

func helloworld(wg *sync.WaitGroup ,st string) {
	defer wg.Done()
	fmt.Println(st)
}

func goodbye(wg *sync.WaitGroup ,st string) {
	defer wg.Done()
	fmt.Println(st)
}
