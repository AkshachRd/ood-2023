package quack

import (
	"fmt"
)

func QuackBehavior() IQuackBehavior {
	return func() {
		fmt.Println("Quack Quack!!!")
	}
}
