package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/quack"
)

func NewMallardDuck() *Duck {
	return NewDuck(fly.FlyWithWings(), quack.QuackBehavior(), dance.DanceWaltz(), func() {
		fmt.Println("I'm mallard duck")
	})
}
