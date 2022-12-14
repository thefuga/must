package must

import "github.com/thefuga/must/internal/reflect"

// Return is an alias for Must, which, when importing the package, may read more naturally.
// e.g.: must.Return(must.Values(foo()), target...)
func Return(results []any, targets ...any) {
	Must(results, targets...)
}

// Must receives a slice representing the return values of a function and a list of
// targets to set the results to. Should the last element of the results slice be
// an error, Must panics with the error value.
// All result values are set to their targets. In case there are less targets than
// results, only the results matching a target will be set. Any missmatch between
// a result and a target is ignored.
func Must(results []any, targets ...any) {
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
