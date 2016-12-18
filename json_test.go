package mtg

import (
	"encoding/json"
	"testing"
	"time"
)

func Test_Date(t *testing.T) {
	tests := map[string]time.Time{
		`"2000"`:       time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		`"2001-12"`:    time.Date(2001, 12, 1, 0, 0, 0, 0, time.UTC),
		`"2010-05-03"`: time.Date(2010, 5, 3, 0, 0, 0, 0, time.UTC),
	}

	for js, expected := range tests {
		var d Date
		err := json.Unmarshal([]byte(js), &d)
		if err != nil {
			t.Error(err)
		} else {
			val := time.Time(d)
			if val != expected {
				t.Errorf("Failed to decode %q. (Got %v)", js, val)
			}
		}
	}
}

func Test_DateErrors(t *testing.T) {
	tests := []string{
		`""`,
		`2000`,
		`null`,
		`2001-02-30`,
		`2001-01-01 10:11`,
	}
	for _, js := range tests {
		var d Date
		err := json.Unmarshal([]byte(js), &d)
		if err == nil {
			t.Errorf("%q was decoded without any error.", js)
		}
	}
}

func Test_BoosterContent(t *testing.T) {
	tests := map[string][]string{
		`"Common"`:             []string{"Common"},
		`["Uncommon", "Rare"]`: []string{"Uncommon", "Rare"},
	}

	equals := func(a, b []string) bool {
		if len(a) != len(b) {
			return false
		} else {
			for i := 0; i < len(a); i++ {
				if a[i] != b[i] {
					return false
				}
			}
		}
		return true
	}

	for js, expected := range tests {
		var bc BoosterContent
		err := json.Unmarshal([]byte(js), &bc)
		if err != nil {
			t.Error(err)
		} else {
			if !equals(expected, bc) {
				t.Errorf("Failed to decode %q. (Got %q)", js, bc)
			}
		}
	}
}
