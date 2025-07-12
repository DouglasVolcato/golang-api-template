package main

import (
	"fmt"
	"time"
)

func main() {
	itr := 10

	for itr > 0 {
		itr -= 1
		time.Sleep(time.Second * 1)
		fmt.Println(itr)
	}
}
