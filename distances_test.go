package main

import (
	"metrics/distances"
	"testing"
)


func TestIncorrectDistance(t *testing.T) {
	_, err := distances.New(-1, -1, -1, -1 , -1)
	if err == nil {
		t.Error("distances.New()")
	}
}

func TestDistance(t *testing.T) {
	value, err := distances.New(0, 2, 3, 0 ,8)
	if err != nil {
		t.Fatal("Initialization error")
	}
	expectedValue := 2308
	if value.Value() != expectedValue {
		t.Errorf(".Value() error, expected=%d, result=%d", expectedValue, value.Value())
	}

	expectedString := "0 км 2 м 3 дм 0 см 8 мм"
	expectedNoZeroString := "2 м 3 дм 8 мм"
	approximation := distances.Downward

	if value.String() != expectedString {
		t.Errorf("value.String() == %s, expected %s", value, expectedString)
	}

	if value.StringNoZero(approximation) != expectedNoZeroString {
		t.Errorf("value.StringNoZero(distances.Downward) == %s, expected %s", value.StringNoZero(approximation), expectedNoZeroString)
	}
}

func TestApproximation(t *testing.T) {
	expectedDownward := 10
	expectedUpward := 11
	expectedRouting := 11

	downwardApproximation := distances.Downward
	upwardApproximation := distances.Upward
	routingApproximation := distances.Rounding

	value, err := distances.New(0,10, 5, 0, 0)
	if err != nil {
		t.Fatal("Initialization error")
	}

	if value.Meters(downwardApproximation) != expectedDownward {
		t.Errorf("value.Meters(DOWNARD) == %d expected %d", value.Meters(downwardApproximation), expectedDownward)
	}

	if value.Meters(upwardApproximation) != expectedUpward {
		t.Errorf("value.Meters(Upward) == %d expected %d", value.Meters(upwardApproximation), expectedUpward)
	}

	if value.Meters(routingApproximation) != expectedRouting {
		t.Errorf("value.Meters(Rounding) == %d expected %d", value.Meters(routingApproximation), expectedRouting)
	}
}
