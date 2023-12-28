package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/quack"
)

type MallardDuck struct {
	Duck
}

func InitMallardDuck() *MallardDuck {
	return &MallardDuck{*InitDuck(&fly.FlyWithWings{}, &quack.QuackBehavior{}, &dance.DanceWaltz{})}
}

func (md *MallardDuck) Display() {
	fmt.Println("I'm mallard duck")
}
