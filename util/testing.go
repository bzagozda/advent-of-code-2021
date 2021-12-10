package util

import (
	"reflect"
	"testing"
)

func Assert(t *testing.T, got, expected interface{}) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got: %v, Expected: %v", got, expected)
	}
}
