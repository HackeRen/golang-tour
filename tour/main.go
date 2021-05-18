package main

import (
	"log"
	"tour/cmd"
)

func main() {
	err := cmd.Excute()
	if err != nil {
		log.Fatalf("cmd.Excute err: %v", err)
	}
}
