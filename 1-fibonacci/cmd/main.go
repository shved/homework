package main

import (
	"flag"
	"fmt"
	"os"

	"fib"
)

func main() {
	var step = flag.Int("step", 20, "fibonacci sequence step number")

	flag.Parse()

	if *step < 1 {
		fmt.Println("step is too low")
		os.Exit(1)
	}

	if *step > 20 {
		fmt.Println("step is too high")
		os.Exit(1)
	}

	fmt.Println(fib.Calc(*step))
}
