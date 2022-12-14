package must

import (
	"github.com/thefuga/must/internal/reflect"
)

// Return is an alias for Must, which, when importing the package, may read more naturally.
// e.g.: must.Return(must.Values(foo()), target...)
func Return[T any](results []any, targets ...any) {
	Must(results, targets...)
}

// Must is the generic implementation of Must. It is specially usefull to populate
// a slice of targets of the same type. E.g.: Must[int](FuncReturningInts(), intTargetSlice...)
func Must[T any](results []any, targets ...T) {
	if results == nil {
		return
	}

	if err := reflect.CastErr(results[len(results)-1]); err != nil {
		panic(err)
	}

	for i := range targets {
		reflect.Set(targets[i], results[i])
	}
}
