package dance

import (
	"fmt"
)

func DanceWaltz() IDanceBehavior {
	return func() {
		fmt.Println("I'm dancing a waltz")
	}
}
