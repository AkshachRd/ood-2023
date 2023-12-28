package dance

import (
	"fmt"
)

func DanceMinuet() IDanceBehavior {
	return func() {
		fmt.Println("I'm dancing a minuet")

	}
}
