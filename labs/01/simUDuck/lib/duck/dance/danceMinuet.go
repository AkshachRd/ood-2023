package dance

import (
	"fmt"
)

type DanceMinuet struct{}

func (dm *DanceMinuet) Dance() {
	fmt.Println("I'm dancing a minuet")
}
