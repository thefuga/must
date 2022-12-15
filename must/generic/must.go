package must

import (
	"github.com/thefuga/must/internal/reflect"
)

// Must is the generic implementation of Must. It is specially usefull to populate
// a slice of targets of the same type. E.g.: Must[int](FuncReturningInts(), intTargetSlice...)
// The results matching T will be populated into the targets slice as they appear.
// The targets slice must be previously allocated.
// Must returns after the capacity of the targets slic is fullfilled. After this point,
// all (if any) remaining result values will be discarded.
func Must[T any](results []any, targets []T) {
	if results == nil {
		return
	}

	if err := reflect.CastErr(results[len(results)-1]); err != nil {
		panic(err)
	}

	for i := range targets {
		v, ok := results[i].(T)
		if !ok {
			continue
		}

		if len(targets) <= i {
			return
		}

		targets[i] = v
	}
}
