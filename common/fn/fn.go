package fn

type ConditionalReduceAction[T any] struct {
	Condition func(acc T) bool
	Action    func(acc T) T
}

func ConditionalReduce[T any](initialValue T, crds ...ConditionalReduceAction[T]) T {
	acc := initialValue
	for _, crd := range crds {
		if crd.Condition(acc) {
			acc = crd.Action(acc)
		}
	}
	return acc
}
