package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	maxp := flag.Int("max", 6, "the max value")
	fmt.Println(*maxp)
	fmt.Println(rand.Intn(*maxp))
}
