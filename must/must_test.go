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

	errFunc := func(a int, b string, c structArg, err error) (int, string, structArg, error) {
		return a, b, c, err
	}

	testCases := []struct {
		description string
		intInput    int
		stringInput string
		structInput structArg
		err         error
	}{
		{
			description: "no error",
			intInput:    1,
			stringInput: "foo",
			structInput: structArg{
				unexported: 2,
				Exported:   "bar",
				pointer:    new(int),
			},
			err: nil,
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

			var (
				intResult    int
				stringResult string
				structResult structArg
			)

			Must(
				Values(errFunc(tc.intInput, tc.stringInput, tc.structInput, tc.err)),
				&intResult, &stringResult, &structResult,
			)

			if !reflect.DeepEqual(tc.intInput, intResult) {
				t.Errorf("expected '%d', got '%d'", tc.intInput, intResult)
			}

			if !reflect.DeepEqual(tc.stringInput, stringResult) {
				t.Errorf("expected '%s', got '%s'", tc.stringInput, stringResult)
			}

			if !reflect.DeepEqual(tc.structInput, structResult) {
				t.Errorf("expected '%v', got '%v'", tc.structInput, structResult)
			}
		})
	}
}
