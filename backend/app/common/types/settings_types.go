package commonType

// "None 10", "One 10.9", "Two 10.98"
type DecimalPlaces int

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
	return [...]string{"None 10", "One 10.9", "Two 10.98"}[t]
}

func (t DecimalPlaces) IsValid() bool {
	return t >= DecimalPlacesNone && t <= DecimalPlacesTwo
}

// "2,10,343.20", "2.10.343,20", "2 10 343,20", "2 10 343.20"
type NumberFormat int

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
	return [...]string{"First 2,10,343.20", "Second 2.10.343,20", "Third 2 10 343,20", "Fourth 2 10 343.20"}[t]
}

func (t NumberFormat) IsValid() bool {
	return t >= NumberFormatFirst && t <= NumberFormatFourth
}
