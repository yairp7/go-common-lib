package fn

type ConditionalReduceAction[T any] struct {
	condition func(acc T) bool
	action    func(acc T) T
}

func ConditionalReduce[T any](initialValue T, crds ...ConditionalReduceAction[T]) T {
	acc := initialValue
	for _, crd := range crds {
		if crd.condition(acc) {
			acc = crd.action(acc)
		}
	}
	return acc
}
