package quack

import (
	"fmt"
)

func SqueakBehavior() IQuackBehavior {
	return func() {
		fmt.Println("Squeek!!!")
	}
}
