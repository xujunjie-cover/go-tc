package tc

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRsvp(t *testing.T) {
	tests := map[string]struct {
		val  Rsvp
		err1 error
		err2 error
	}{
		"empty":  {err1: fmt.Errorf("Rsvp options are missing")},
		"simple": {val: Rsvp{ClassID: 42, Police: &Police{AvRate: 1337, Result: 12}}},
	}

	for name, testcase := range tests {
		t.Run(name, func(t *testing.T) {
			data, err1 := marshalRsvp(&testcase.val)
			if err1 != nil {
				if testcase.err1 != nil && testcase.err1.Error() == err1.Error() {
					return
				}
				t.Fatalf("Unexpected error: %v", err1)
			}
			val := Rsvp{}
			err2 := unmarshalRsvp(data, &val)
			if err2 != nil {
				if testcase.err2 != nil && testcase.err2.Error() == err2.Error() {
					return
				}
				t.Fatalf("Unexpected error: %v", err2)

			}
			if diff := cmp.Diff(val, testcase.val); diff != "" {
				t.Fatalf("Rsvp missmatch (want +got):\n%s", diff)
			}
		})
	}
}