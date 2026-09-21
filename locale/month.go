package locale

import "time"

// MonthByName returns the month corresponding to the given month name. It
// uses English names. If the name is not recognised a non-nil error is
// returned.
func MonthByName(name string) (time.Month, error) {
	return TimeEnglish.ToMonth(name)
}

// MonthByNameByCountry returns the month corresponding to the given month
// name for the given country.
//
// For how the name to month map is chosen see [TimeByName]
//
// If the month name or country code is not recognised a non-nil error is
// returned.
func MonthByNameByCountry(name, country string) (time.Month, error) {
	l, err := TimeByName(country)
	if err != nil {
		return 0, err
	}

	return l.ToMonth(name)
}

// MonthByNameByLanguage returns the month corresponding to the given
// month name for the given language.
//
// It is an alias for [MonthByNameByCountry].
func MonthByNameByLanguage(name, language string) (time.Month, error) {
	return MonthByNameByCountry(name, language)
}
