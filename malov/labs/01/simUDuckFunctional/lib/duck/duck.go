package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/quack"
)

type Duck struct {
	PerformFly   fly.IFlyBehavior
	PerformQuack quack.IQuackBehavior
	PerformDance dance.IDanceBehavior
	Swim         func()
	Display      func()
}

func NewDuck(fb fly.IFlyBehavior, qb quack.IQuackBehavior, db dance.IDanceBehavior, d func()) *Duck {
	return &Duck{
		PerformFly:   fb,
		PerformQuack: qb,
		PerformDance: db,
		Swim: func() {
			fmt.Println("I'm swimming")
		},
		Display: d,
	}
}
