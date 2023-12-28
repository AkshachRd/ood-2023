package lib

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck"
)

func DrawDuck(duck *duck.Duck) {
	duck.Display()
}
func PlayWithDuck(duck *duck.Duck) {
	DrawDuck(duck)
	duck.Quack()
	duck.Fly()
	duck.Dance()
	fmt.Println()
}
