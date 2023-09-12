package dance

import (
	"fmt"
)

func DanceBehavior() IDanceBehavior {
	return func() {
		fmt.Println("I'm Dancing")
	}
}
