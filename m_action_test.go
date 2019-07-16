package tc

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAction(t *testing.T) {
	tests := map[string]struct {
		val  Action
		err1 error
		err2 error
	}{
		"empty":  {err1: fmt.Errorf("Action options are missing")},
		"simple": {val: Action{Kind: "test", Index: 123}},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			data, err1 := marshalAction(&testcase.val)
			if err1 != nil {
				if testcase.err1 != nil && testcase.err1.Error() == err1.Error() {
					return
				}
				t.Fatalf("Unexpected error: %v", err1)
			}
			val := Action{}
			err2 := unmarshalAction(data, &val)
			if err2 != nil {
				if testcase.err2 != nil && testcase.err2.Error() == err2.Error() {
					return
				}
				t.Fatalf("Unexpected error: %v", err2)

			}
			if diff := cmp.Diff(val, testcase.val); diff != "" {
				t.Fatalf("Action missmatch (want +got):\n%s", diff)
			}
		})
	}
}