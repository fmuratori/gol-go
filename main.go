package main

import (
	"fmt"
	core "golgo/internal"
	"time"
)

func main() {
	v := core.NewBoard(9, 9)

	v.AddCell(3, 4)
	v.AddCell(4, 3)
	v.AddCell(4, 4)
	v.AddCell(4, 5)
	v.AddCell(5, 4)

	fmt.Print(v.String())
	for range 10 {
		v.Tick()
		fmt.Print(v.String())
		time.Sleep(1 * time.Second)
	}

}
