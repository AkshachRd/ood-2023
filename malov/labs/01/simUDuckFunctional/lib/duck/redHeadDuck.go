package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/quack"
)

func NewRedHeadDuck() *Duck {
	return NewDuck(fly.FlyWithWings(), quack.QuackBehavior(), dance.DanceMinuet(), func() {
		fmt.Println("I'm redhead duck")
	})
}
