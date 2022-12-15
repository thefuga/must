package must

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/thefuga/must/functions"
)

func Test_Must(t *testing.T) {
	type structArg struct {
		unexported int
		Exported   string
		pointer    *int
	}

	errFunc := func(a, b int, err error) (int, int, error) {
		return a, b, err
	}

	testCases := []struct {
		description string
		resultA     int
		resultB     int
		err         error
	}{
		{
			description: "no error",
			resultA:     1,
			resultB:     2,
			err:         nil,
		},
		{
			description: "func returning error",
			err:         fmt.Errorf("function returned an unexpectedError"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			defer func() {
				if err := recover(); err != tc.err {
					t.Errorf("expected panic with '%v', got '%v'", tc.err, err)
				}
			}()

			var targets = make([]int, 2)

			Must(
				Values(errFunc(tc.resultA, tc.resultB, tc.err)),
				targets,
			)

			if !reflect.DeepEqual(tc.resultA, targets[0]) {
				t.Errorf("expected '%d', got '%d'", tc.resultA, targets[0])
			}

			if !reflect.DeepEqual(tc.resultB, targets[1]) {
				t.Errorf("expected '%d', got '%d'", tc.resultB, targets[1])
			}
		})
	}
}
