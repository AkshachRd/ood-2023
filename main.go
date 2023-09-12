package main

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck"
	simUDuckFn "github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional"
)

func main() {
	fmt.Println("SimUDuck")
	simUDuck.Main()

	fmt.Println("SimUDuckFunctional")
	simUDuckFn.Main()
}
