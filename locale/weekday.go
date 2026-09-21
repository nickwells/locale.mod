package locale

import (
	"time"
)

// WeekdayByName returns the weekday corresponding to the given weekday
// name. It uses English names. If the name is not recognised a non-nil error
// is returned.
//
// See [WeekdayByNameByCountry] and [WeekdayByNameByLanguage] if you want to
// use alternative languages.
func WeekdayByName(name string) (time.Weekday, error) {
	return TimeEnglish.ToWeekday(name)
}

// WeekdayByNameByCountry returns the weekday corresponding to the given
// weekday name for the given country.
//
// For how the name to weekday map is chosen see [TimeByName]
//
// If the weekday name or country code is not recognised a non-nil error is
// returned.
//
// See also [WeekdayByNameByLanguage].
func WeekdayByNameByCountry(name, country string) (time.Weekday, error) {
	t, err := TimeByName(country)
	if err != nil {
		return 0, err
	}

	return t.ToWeekday(name)
}

// WeekdayByNameByLanguage returns the weekday corresponding to the given
// weekday name for the given language.
//
// It is an alias for [WeekdayByNameByCountry].
func WeekdayByNameByLanguage(name, language string) (time.Weekday, error) {
	return WeekdayByNameByCountry(name, language)
}
