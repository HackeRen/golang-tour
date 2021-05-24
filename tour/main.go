package main

import (
	"fmt"
	"log"
	"tour/cmd"
)

type entry struct {
	off uint32
	len uint32
}

func main() {
	err := cmd.Excute()
	if err != nil {
		log.Fatalf("cmd.Excute err: %v", err)
	}

	var data []entry

	var ent = &entry{off: 1, len: 2}
	data = append(data, *ent)

	fmt.Println(data)
	data[0] = entry{2, 3}
	fmt.Println(data)
}
