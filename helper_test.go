package mtg

import (
	"fmt"
	"time"

	"gopkg.in/jarcoal/httpmock.v1"
)

func NewStringResponderWithHeader(status int, body string, header map[string]string) httpmock.Responder {
	resp := httpmock.NewStringResponse(status, body)
	for k, v := range header {
		resp.Header.Set(k, v)
	}
	return httpmock.ResponderFromResponse(resp)
}

func ShouldContainCard(actual interface{}, expected ...interface{}) string {
	if len(expected) != 1 {
		return "Invalid parameters for ShouldContainCard"
	}

	cards, firstOk := actual.([]*Card)
	cardName, secondOk := expected[0].(string)

	if !firstOk || !secondOk {
		return "Invalid Arguments for ShouldContainCard"
	}

	for _, card := range cards {
		if card.Name == cardName {
			return ""
		}
	}

	return fmt.Sprintf("Card %q was not found in %q", cardName, cards)
}

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
