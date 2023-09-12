package dance

import (
	"fmt"
)

type DanceBehavior struct{}

func (db *DanceBehavior) Dance() {
	fmt.Println("I'm Dancing")
}
