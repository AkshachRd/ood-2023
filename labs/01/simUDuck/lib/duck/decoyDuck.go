package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/quack"
)

type DecoyDuck struct {
	Duck
}

func InitDecoyDuck() *DecoyDuck {
	return &DecoyDuck{*InitDuck(&fly.FlyNoWay{}, &quack.MuteQuackBehavior{}, &dance.DanceNoWay{})}
}

func (dd *DecoyDuck) Display() {
	fmt.Println("I'm decoy duck")
}
