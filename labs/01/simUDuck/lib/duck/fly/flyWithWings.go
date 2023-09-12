package fly

import (
	"fmt"
)

type FlyWithWings struct {
	flyCount uint
}

func (fww *FlyWithWings) Fly() {
	fww.flyCount++
	fmt.Println("I'm flying with wings for the ", fww.flyCount, " time!!")
}
