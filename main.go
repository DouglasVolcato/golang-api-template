package main

import "log"

func main() {
	api := NewApi()
	api.Mount()
	err := api.Run()
	if err != nil {
		log.Fatal(err)
	}
}
