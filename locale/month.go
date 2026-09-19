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
// Countries are identified by the ISO 3166 code. Note that for some
// countries having multiple languages the other languages can be selected by
// prefixing the country code with the two-letter language code (as defined
// by ISO 639) and a dash; for instance, to select French Canadian 'fr-CA'
// may be used. Additionally, the common, english names of the languages may
// be used or just the ISO 639 codes.
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
