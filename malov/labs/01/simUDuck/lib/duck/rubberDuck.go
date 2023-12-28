package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/quack"
)

type RubberDuck struct {
	Duck
}

func InitRubberDuck() *RubberDuck {
	return &RubberDuck{*InitDuck(&fly.FlyNoWay{}, &quack.SqueakBehavior{}, &dance.DanceNoWay{})}
}

func (rd *RubberDuck) Display() {
	fmt.Println("I'm rubber duck")
}
