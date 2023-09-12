package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/quack"
)

func NewModelDuck() *Duck {
	return NewDuck(fly.FlyNoWay(), quack.QuackBehavior(), dance.DanceNoWay(), func() {
		fmt.Println("I'm model duck")
	})
}
