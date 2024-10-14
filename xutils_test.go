package xutils

import (
	"testing"
)

func TestStringToInt64(t *testing.T) {

	for _, v := range []struct {
		Value    string
		Expected int64
	}{
		{
			Value:    "1",
			Expected: 1,
		},
		{
			Value:    "123-4-8978-4566-22",
			Expected: 12348978456622,
		},
		{
			Value:    "a2db3",
			Expected: 23,
		},
	} {
		result, _ := StringToInt64(v.Value)
		if result != v.Expected {
			t.Error("wrong result. expected=", v.Expected, ", got=", v.Expected)
		}
	}
}
