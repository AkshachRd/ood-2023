package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/quack"
)

type RedHeadDuck struct {
	Duck
}

func InitRedHeadDuck() *RedHeadDuck {
	return &RedHeadDuck{*InitDuck(&fly.FlyWithWings{}, &quack.QuackBehavior{}, &dance.DanceMinuet{})}
}

func (rhd *RedHeadDuck) Display() {
	fmt.Println("I'm redhead duck")
}
