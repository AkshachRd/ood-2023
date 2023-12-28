package quack

import (
	"fmt"
)

type QuackBehavior struct{}

func (qb *QuackBehavior) Quack() {
	fmt.Println("Quack Quack!!!")
}
