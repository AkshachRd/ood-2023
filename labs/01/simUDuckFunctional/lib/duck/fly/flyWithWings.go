package fly

import (
	"fmt"
)

func FlyWithWings() IFlyBehavior {
	var flyCount uint

	return func() {
		flyCount++
		fmt.Println("I'm flying with wings for the ", flyCount, " time!!")
	}
}
