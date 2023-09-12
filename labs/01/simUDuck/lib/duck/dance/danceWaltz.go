package dance

import (
	"fmt"
)

type DanceWaltz struct{}

func (dw *DanceWaltz) Dance() {
	fmt.Println("I'm dancing a waltz")
}
