package commonType

// "None 10", "One 10.9", "Two 10.98"
type DecimalPlaces int8

const (
	DecimalPlacesNone DecimalPlaces = iota + 1
	DecimalPlacesOne
	DecimalPlacesTwo
)

func AllDecimalPlaces() []DecimalPlaces {
	return []DecimalPlaces{
		DecimalPlacesNone,
		DecimalPlacesOne,
		DecimalPlacesTwo,
	}
}

func (t DecimalPlaces) String() string {
	names := [...]string{"None 10", "One 10.9", "Two 10.98"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t DecimalPlaces) IsValid() bool {
	return t >= DecimalPlacesNone && t <= DecimalPlacesTwo
}

// "2,10,343.20", "2.10.343,20", "2 10 343,20", "2 10 343.20"
type NumberFormat int8

const (
	NumberFormatFirst NumberFormat = iota + 1
	NumberFormatSecond
	NumberFormatThird
	NumberFormatFourth
)

func AllNumberFormats() []NumberFormat {
	return []NumberFormat{
		NumberFormatFirst,
		NumberFormatSecond,
		NumberFormatThird,
		NumberFormatFourth,
	}
}

func (t NumberFormat) String() string {
	names := [...]string{"", "First 2,10,343.20", "Second 2.10.343,20", "Third 2 10 343,20", "Fourth 2 10 343.20"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t NumberFormat) IsValid() bool {
	return t >= NumberFormatFirst && t <= NumberFormatFourth
}
