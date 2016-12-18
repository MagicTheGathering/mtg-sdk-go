package mtg

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func ShouldBeOn(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return "Invalid parameters for ShouldBeOn"
	}
	var actualTime time.Time
	date, firstOk := actual.(Date)
	if firstOk {
		actualTime = time.Time(date)
	} else {
		actualTime, firstOk = actual.(time.Time)
	}

	expectedTime, secondOk := expected[0].(time.Time)

	if !firstOk || !secondOk {
		return "ShouldBeOn should be used on date / time values"
	}

	if actualTime.Before(expectedTime) || actualTime.After(expectedTime) {
		return fmt.Sprintf("Expected to happen on %v but happened on %v", expectedTime, actualTime)
	}

	return ""
}

func Test_Date(t *testing.T) {
	Convey("json date decoding", t, func() {
		var date Date

		Convey("with only a year", func() {
			err := json.Unmarshal([]byte(`"2000"`), &date)

			So(err, ShouldBeNil)
			So(date, ShouldBeOn, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
		})

		Convey("with a year and a month", func() {
			err := json.Unmarshal([]byte(`"2001-12"`), &date)
			So(err, ShouldBeNil)
			So(date, ShouldBeOn, time.Date(2001, 12, 1, 0, 0, 0, 0, time.UTC))
		})

		Convey("with a full date", func() {
			err := json.Unmarshal([]byte(`"2010-03-12"`), &date)
			So(err, ShouldBeNil)
			So(date, ShouldBeOn, time.Date(2010, 3, 12, 0, 0, 0, 0, time.UTC))
		})

		Convey("with some invalid input", func() {
			Convey("empty string", func() {
				err := json.Unmarshal([]byte(`""`), &date)
				So(err, ShouldNotBeNil)
			})

			Convey("a number", func() {
				err := json.Unmarshal([]byte(`2000`), &date)
				So(err, ShouldNotBeNil)
			})

			Convey("null", func() {
				err := json.Unmarshal([]byte(`null`), &date)
				So(err, ShouldNotBeNil)
			})

			Convey("an invalid date", func() {
				err := json.Unmarshal([]byte(`"2001-02-30"`), &date)
				So(err, ShouldNotBeNil)
			})

			Convey("a date with time", func() {
				err := json.Unmarshal([]byte(`"2001-01-01 10:11"`), &date)
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func Test_BoosterContent(t *testing.T) {
	Convey("json BoosterContent decoding", t, func() {
		var bc BoosterContent

		Convey("a single card type", func() {
			err := json.Unmarshal([]byte(`"Common"`), &bc)
			So(err, ShouldBeNil)
			So(bc, ShouldResemble, BoosterContent{"Common"})
		})

		Convey("two or more different card types", func() {
			err := json.Unmarshal([]byte(`["Common", "Rare"]`), &bc)
			So(err, ShouldBeNil)
			So(bc, ShouldResemble, BoosterContent{"Common", "Rare"})

			err = json.Unmarshal([]byte(`["Common", "Uncommon", "Rare"]`), &bc)
			So(err, ShouldBeNil)
			So(bc, ShouldResemble, BoosterContent{"Common", "Uncommon", "Rare"})
		})

		Convey("other values should return an error", func() {
			err := json.Unmarshal([]byte(`["Common", 123, "Rare"]`), &bc)
			So(err, ShouldNotBeNil)

			err = json.Unmarshal([]byte(`false`), &bc)
			So(err, ShouldNotBeNil)
		})
	})
}
