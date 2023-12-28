package simUDuck

import (
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/fly"
)

func Main() {
	var duckInstance *duck.Duck

	mallardDuck := duck.InitMallardDuck()
	duckInstance = &mallardDuck.Duck
	lib.PlayWithDuck(duckInstance)
	mallardDuck.Fly()

	redheadDuck := duck.InitRedHeadDuck()
	duckInstance = &redheadDuck.Duck
	lib.PlayWithDuck(duckInstance)

	rubberDuck := duck.InitRubberDuck()
	duckInstance = &rubberDuck.Duck
	lib.PlayWithDuck(duckInstance)
	rubberDuck.Fly()

	decoyDuck := duck.InitDecoyDuck()
	duckInstance = &decoyDuck.Duck
	lib.PlayWithDuck(duckInstance)

	modelDuck := duck.InitModelDuck()
	duckInstance = &modelDuck.Duck
	lib.PlayWithDuck(duckInstance)

	modelDuck.SetFlyBehavior(&fly.FlyWithWings{})
	duckInstance = &modelDuck.Duck
	lib.PlayWithDuck(duckInstance)
}
