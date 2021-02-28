package distances

import (
	"errors"
	"fmt"
	"strings"
)

type Distances int

const (
	MillimetersToKilometers  Distances = 1000000
	MillimetersToMeters      Distances = 1000
	MillimetersToDecimeters  Distances = 100
	MillimetersToCentimeters Distances = 10
	Millimeters              Distances = 1
)

func (d Distances) Value() int {
	return int(d)
}

type Approximation int

const (
	Rounding Approximation = iota // Округление
	Downward                      // Приближение с недостатком
	Upward                        // Приближение с избытком
)

func checkValue(value, expected int, result, other string) string {
	if value == expected {
		return result
	} else {
		return other
	}
}

func nearestNumber(num, previous int, numDistance, previousDistance Distances) int {
	if previous * previousDistance.Value() >= numDistance.Value() / 2 {
		return num + 1
	} else {
		return num
	}
}

func convertDistance(num, previous int, numDistance, previousDistance Distances, approximation Approximation) int {
	switch approximation {
	case Downward:
		return num
	case Upward:
		return num + 1
	case Rounding:
		return nearestNumber(num, previous, numDistance, previousDistance)
	default:
		return num
	}
}

type Distance struct {
	millimeters int
}

func New(kilometers, meters, decimeters, centimeters, millimeters int) (*Distance, error)  {
	if kilometers < 0 || meters < 0 || centimeters < 0 || millimeters < 0 {
		return nil, errors.New("передан отрицательный аргумент")
	} else {
		return &Distance{
			millimeters: kilometers * MillimetersToKilometers.Value() + meters * MillimetersToMeters.Value() +
				decimeters * MillimetersToDecimeters.Value() + centimeters * MillimetersToCentimeters.Value() + millimeters,
		}, nil
	}
}

func (d Distance) Millimeters() int {
	return d.millimeters % MillimetersToCentimeters.Value()
}

func (d Distance) Centimeters(approximation Approximation) int {
	centimeters := (d.millimeters % MillimetersToDecimeters.Value()) / MillimetersToCentimeters.Value()
	return convertDistance(centimeters, d.Millimeters(), MillimetersToCentimeters, Millimeters, approximation)
}

func (d Distance) Decimeters(approximation Approximation) int {
	decimeters := (d.millimeters % MillimetersToMeters.Value()) / MillimetersToDecimeters.Value()
	return convertDistance(decimeters, d.Centimeters(Downward), MillimetersToDecimeters, MillimetersToCentimeters, approximation)
}

func (d Distance) Meters(approximation Approximation) int {
	meters := (d.millimeters % MillimetersToKilometers.Value()) / MillimetersToMeters.Value()
	return convertDistance(meters, d.Decimeters(Downward), MillimetersToMeters, MillimetersToDecimeters, approximation)
}

func (d Distance) Kilometers(approximation Approximation) int {
	kilometers := d.millimeters / MillimetersToKilometers.Value()
	return convertDistance(kilometers, d.Meters(Downward), MillimetersToKilometers, MillimetersToMeters, approximation)
}

func (d Distance) strMillimeters() string {
	return fmt.Sprintf("%d мм ", d.Millimeters())
}

func (d Distance) strCentimeters(approximation Approximation) string {
	return fmt.Sprintf("%d см ", d.Centimeters(approximation))
}

func (d Distance) strDecimeters(approximation Approximation) string {
	return fmt.Sprintf("%d дм ", d.Decimeters(approximation))
}

func (d Distance) strMeters(approximation Approximation) string {
	return fmt.Sprintf("%d м ", d.Meters(approximation))
}

func (d Distance) strKilometers(approximation Approximation) string {
	return fmt.Sprintf("%d км ", d.Kilometers(approximation))
}

func (d Distance) String() string {
	return strings.TrimRight(fmt.Sprintf("%s%s%s%s%s", d.strKilometers(Downward), d.strMeters(Downward), d.strDecimeters(Downward), d.strCentimeters(Downward), d.strMillimeters()), " ")
}

func (d Distance) Value() int {
	return d.millimeters
}

func (d Distance) StringNoZero(approximation Approximation) string {
	millimeters := checkValue(d.Millimeters(), 0, "", d.strMillimeters())
	centimeters := checkValue(d.Centimeters(approximation), 0, "", d.strCentimeters(approximation))
	decimeters := checkValue(d.Decimeters(approximation), 0, "", d.strDecimeters(approximation))
	meters := checkValue(d.Meters(approximation), 0, "", d.strMeters(approximation))
	kilometers := checkValue(d.Kilometers(approximation), 0, "", d.strKilometers(approximation))
	return strings.TrimRight(kilometers + meters + decimeters + centimeters + millimeters, " ")
}
