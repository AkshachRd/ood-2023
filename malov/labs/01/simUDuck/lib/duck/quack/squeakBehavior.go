package quack

import (
	"fmt"
)

type SqueakBehavior struct{}

func (sb *SqueakBehavior) Quack() {
	fmt.Println("Squeek!!!")
}
