package lib

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck"
)

func DrawDuck(duck *duck.Duck) {
	duck.Display()
}
func PlayWithDuck(duck *duck.Duck) {
	DrawDuck(duck)
	duck.PerformQuack()
	duck.PerformFly()
	duck.PerformDance()
	fmt.Println()
}
