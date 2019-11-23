package main

import (
	"fmt"
	"os"
)

//PhysicsBall is
func PhysicsBall(h int) int {
	var s int
	s = h * 2 / 3
	return s
}

func main() {
	var height int
	fmt.Fscan(os.Stdin, &height)
	var FirstHeight int = height
	step := 0
	for i := 0; height > 32; i++ {
		height = PhysicsBall(height)
		step = i
	}
	fmt.Println("С высоты ", FirstHeight, " потребуется ", step, " шагов")
}
