package kronos

import (
	"time"
)

type countries interface {
	Turkish(time.Time) (string, string)
	English()
}

type kronosInstance struct {
	countries
}

func (c *kronosInstance) English(t time.Time) (string, string) {

	Month := t.Month().String()
	Day := t.Weekday().String()

	resultMonthName, resultDayName := getOwnLangFormat(Month, Day, "English")

	return resultMonthName, resultDayName
}

func (c *kronosInstance) Turkish(t time.Time) (string, string) {

	Month := t.Month().String()
	Day := t.Weekday().String()

	resultMonthName, resultDayName := getOwnLangFormat(Month, Day, "Turkish")

	return resultMonthName, resultDayName
}

func getOwnLangFormat(Month string, Day string, whichCountry string) (string, string) {

	CountryMonth := whichCountry + "Month"
	CountryDay := whichCountry + "Day"

	var resultMonthName string = "none"
	var resultDayName string = "none"
	var breakFlag int = 2

	var monthRowIndx uint8 = 0
	var dayRowIndx uint8 = 0

	for i := 0; i < len((*allCountries)["EnglishMonth"]); i++ {

		if (*allCountries)["EnglishMonth"][i] == Month {
			resultMonthName = (*allCountries)["EnglishMonth"][i]
			monthRowIndx = uint8(i)
			breakFlag--
		}

		if i < 7 {
			if (*allCountries)["EnglishDay"][i] == Day {
				resultDayName = (*allCountries)["EnglishDay"][i]
				dayRowIndx = uint8(i)
				breakFlag--
			}

		}
		if breakFlag >= 2 {
			break
		}
	}

	if whichCountry != "English" {
		resultMonthName = (*allCountries)[CountryMonth][monthRowIndx]
		resultDayName = (*allCountries)[CountryDay][dayRowIndx]
	}

	return resultMonthName, resultDayName
}

var myNewInstance *kronosInstance

func init() {

	myNewInstance = &kronosInstance{
		countries: nil,
	}
}

func Now() *kronosInstance {
	return myNewInstance
}
