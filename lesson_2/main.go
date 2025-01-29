package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Printf("Done %d \n", i)
		}()
	}

	time.Sleep(30 * time.Second)
}
