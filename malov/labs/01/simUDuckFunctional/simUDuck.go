package simUDuck

import (
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuckFunctional/lib/duck/fly"
)

func Main() {
	mallardDuck := duck.NewMallardDuck()
	lib.PlayWithDuck(mallardDuck)
	mallardDuck.PerformFly()

	redheadDuck := duck.NewRedHeadDuck()
	lib.PlayWithDuck(redheadDuck)

	rubberDuck := duck.NewRubberDuck()
	lib.PlayWithDuck(rubberDuck)
	rubberDuck.PerformFly()

	decoyDuck := duck.NewDecoyDuck()
	lib.PlayWithDuck(decoyDuck)

	modelDuck := duck.NewModelDuck()
	lib.PlayWithDuck(modelDuck)

	modelDuck.PerformFly = fly.FlyWithWings()
	lib.PlayWithDuck(modelDuck)
}
