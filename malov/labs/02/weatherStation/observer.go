package weatherStation

import "sort"

type IObserver[T any] interface {
	Update(data *T)
}

type IObservable[T any] interface {
	RegisterObserver(observer *IObserver[T])
	NotifyObservers(data *T)
	RemoveObserver(observer *IObserver[T])
}

type Observable[T any] struct {
	observers  map[IObserver[T]]bool
	priorities map[IObserver[T]]int
}

func NewObservable[K any]() *Observable[K] {
	return &Observable[K]{
		observers:  make(map[IObserver[K]]bool),
		priorities: make(map[IObserver[K]]int),
	}
}

func (o *Observable[T]) RegisterObserver(observer IObserver[T], priority int) {
	if _, exist := o.observers[observer]; !exist {
		o.observers[observer] = true
		o.priorities[observer] = priority
	}
}

func (o *Observable[T]) NotifyObservers(data *T) {
	// Copying observers to preserve it from deleting during Update
	//observersCopy := make(map[IObserver[T]]bool)
	//for k, v := range o.observers {
	//	observersCopy[k] = v
	//}

	var observersWithPriority []IObserver[T]
	for observer := range o.observers {
		observersWithPriority = append(observersWithPriority, observer)
	}
	sort.Slice(observersWithPriority, func(i, j int) bool {
		return o.priorities[observersWithPriority[i]] < o.priorities[observersWithPriority[j]]
	})

	for _, observer := range observersWithPriority {
		observer.Update(data)
	}
}

func (o *Observable[T]) RemoveObserver(observer IObserver[T]) {
	delete(o.observers, observer)
}
