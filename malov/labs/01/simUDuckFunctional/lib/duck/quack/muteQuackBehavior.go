package quack

func MuteQuackBehavior() IQuackBehavior {
	return func() {}
}
