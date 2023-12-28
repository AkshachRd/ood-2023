package duck

import (
	"fmt"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/dance"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/fly"
	"github.com/AkshachRd/ood-2023/labs/01/simUDuck/lib/duck/quack"
)

type Duck struct {
	flyBehavior   fly.IFlyBehavior
	quackBehavior quack.IQuackBehavior
	danceBehavior dance.IDanceBehavior
}

func InitDuck(fb fly.IFlyBehavior, qb quack.IQuackBehavior, db dance.IDanceBehavior) *Duck {
	return &Duck{
		flyBehavior:   fb,
		quackBehavior: qb,
		danceBehavior: db,
	}
}

func (d *Duck) SetFlyBehavior(fb fly.IFlyBehavior) {
	d.flyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb quack.IQuackBehavior) {
	d.quackBehavior = qb
}

func (d *Duck) SetDanceBehavior(db dance.IDanceBehavior) {
	d.danceBehavior = db
}

func (d *Duck) Quack() {
	d.quackBehavior.Quack()
}

func (d *Duck) Swim() {
	fmt.Println("I'm swimming")
}

func (d *Duck) Fly() {
	d.flyBehavior.Fly()
}

func (d *Duck) Dance() {
	d.danceBehavior.Dance()
}

func (d *Duck) Display() {
}
