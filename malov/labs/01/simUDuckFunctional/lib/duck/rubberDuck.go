package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/quack"
)

func NewRubberDuck() *Duck {
	return NewDuck(fly.FlyNoWay(), quack.SqueakBehavior(), dance.DanceNoWay(), func() {
		fmt.Println("I'm rubber duck")
	})
}
