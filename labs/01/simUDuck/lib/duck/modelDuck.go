package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/quack"
)

type ModelDuck struct {
	Duck
}

func InitModelDuck() *ModelDuck {
	return &ModelDuck{*InitDuck(&fly.FlyNoWay{}, &quack.QuackBehavior{}, &dance.DanceNoWay{})}
}

func (md *ModelDuck) Display() {
	fmt.Println("I'm model duck")
}
