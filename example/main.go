package main

import (
	"fmt"

	"github.com/hikobae/go-random/random"
)

func main() {
	random.Init()
	fmt.Println(random.Alphabetn(3))

	p := []int{1, 2, 3}
	random.ShuffleSlice(p)
	fmt.Println(p)
}
