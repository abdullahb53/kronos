package kronos

import (
	"time"
)

type callingTypes interface {
	FindMonth(string, string) string
	FindDay(string, string) string
	FindMonthInt(int, string) string
}

type countries interface {
	Turkish()
	English()
	Deutsch()
}

type getKronosInstance struct {
	callingTypes
}

type kronosInstance struct {
	countries
}

// It gives you translated "Month" string that what you choose language.
//
// (EnglishMonth, Language) -> Translated String
func (c *getKronosInstance) FindMonth(EnglishString_MonthName string, language string) (string, error) {

	objectToBeFound := "English" + "Month"
	var FoundedIndex uint8 = 255
	for i := 0; i < len((*allCountries)[objectToBeFound]); i++ {
		if (*allCountries)[objectToBeFound][i] == EnglishString_MonthName {
			FoundedIndex = uint8(i)
		}
	}
	if FoundedIndex != 255 {
		objectToBeFound = language + "Month"
		return (*allCountries)[objectToBeFound][FoundedIndex], nil
	} else {
		return "", &time.ParseError{Message: "Check your 'Month' spelling."}
	}

}

// It gives you translated "Day" string that what you choose language.
//
// (EnglishDay, Language) -> Translated String
func (c *getKronosInstance) FindDay(EnglishString_DayName string, language string) (string, error) {
	objectToBeFound := "English" + "Day"
	var FoundedIndex uint8 = 255
	for i := 0; i < len((*allCountries)[objectToBeFound]); i++ {
		if (*allCountries)[objectToBeFound][i] == EnglishString_DayName {
			FoundedIndex = uint8(i)
		}
	}
	if FoundedIndex != 255 {
		objectToBeFound = language + "Day"
		return (*allCountries)[objectToBeFound][FoundedIndex], nil
	} else {
		return "", &time.ParseError{Message: "Not found! check your 'Day' spelling."}
	}

}

// It gives you translated "Month" string that what you find with "Month" number.
//
// (MonthNumber, Language) -> Translated String
func (c *getKronosInstance) FindMonthInt(MonthNumber int, language string) (string, error) {
	objectToBeFound := language + "Month"

	_, errorLanguage := (*allCountries)[objectToBeFound]

	if !errorLanguage {
		if MonthNumber < 12 && MonthNumber > -1 {
			return (*allCountries)[objectToBeFound][MonthNumber], nil
		} else {
			return "", &time.ParseError{Message: "Check your 'month number' spelling."}
		}
	} else {
		return "", &time.ParseError{Message: "Check your 'language' spelling."}
	}

}

func (c *kronosInstance) English(t time.Time) (string, string, error) {

	Month := t.Month().String()
	Day := t.Weekday().String()

	resultMonthName, resultDayName, err := getOwnLangFormat(Month, Day, "English")
	if err != nil {
		return "", "", err
	}

	return resultMonthName, resultDayName, nil
}

func (c *kronosInstance) Turkish(t time.Time) (string, string, error) {

	Month := t.Month().String()
	Day := t.Weekday().String()

	resultMonthName, resultDayName, err := getOwnLangFormat(Month, Day, "Turkish")
	if err != nil {
		return "", "", err
	}

	return resultMonthName, resultDayName, nil
}

func (c *kronosInstance) Deutsch(t time.Time) (string, string, error) {

	Month := t.Month().String()
	Day := t.Weekday().String()

	resultMonthName, resultDayName, err := getOwnLangFormat(Month, Day, "Deutsch")
	if err != nil {
		return "", "", err
	}

	return resultMonthName, resultDayName, nil
}

func getOwnLangFormat(Month string, Day string, whichCountry string) (string, string, error) {

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

	errMsg := ""
	if resultDayName == "none" {
		errMsg += "dayName not founded. Check your spelling. \n"
	}
	if resultMonthName == "none" {
		errMsg += "dayName not founded. Check your spelling. \n"
	}
	if errMsg != "" {
		return "", "", &time.ParseError{Message: errMsg}
	}

	return resultMonthName, resultDayName, nil
}

var myNewInstance *kronosInstance
var myGetInstance *getKronosInstance

func init() {

	myGetInstance = &getKronosInstance{
		callingTypes: nil,
	}

	myNewInstance = &kronosInstance{
		countries: nil,
	}
}

func Now() *kronosInstance {
	return myNewInstance
}

func Get() *getKronosInstance {
	return myGetInstance
}
