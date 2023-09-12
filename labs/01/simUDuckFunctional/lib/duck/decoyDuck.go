package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/quack"
)

func NewDecoyDuck() *Duck {
	return NewDuck(fly.FlyNoWay(), quack.MuteQuackBehavior(), dance.DanceNoWay(), func() {
		fmt.Println("I'm decoy duck")
	})
}
