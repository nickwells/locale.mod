package locale

import (
	"fmt"
)

var namedTimes = map[string]Time{
	"en":      TimeEnglish,
	"English": TimeEnglish,
	"english": TimeEnglish,
	"GB":      TimeEnglish, // Great Britain
	"US":      TimeEnglish, // United States of America
	"CA":      TimeEnglish, // Canada
	"AU":      TimeEnglish, // Australia
	"NZ":      TimeEnglish, // New Zealand

	"fr":       TimeFrench,
	"French":   TimeFrench,
	"Francais": TimeFrench,
	"Français": TimeFrench,
	"french":   TimeFrench,
	"francais": TimeFrench,
	"français": TimeFrench,
	"FR":       TimeFrench, // France
	"fr-CA":    TimeFrench, // French - Canadian
	"fr-CH":    TimeFrench, // French - Switzerland
	"HT":       TimeFrench, // Haiti

	"es":         TimeSpanish,
	"Spanish":    TimeSpanish,
	"Espanol":    TimeSpanish,
	"Español":    TimeSpanish,
	"Castilian":  TimeSpanish,
	"Castellano": TimeSpanish,
	"spanish":    TimeSpanish,
	"espanol":    TimeSpanish,
	"español":    TimeSpanish,
	"castilian":  TimeSpanish,
	"castellano": TimeSpanish,
	"AR":         TimeSpanish, // Argentina
	"BO":         TimeSpanish, // Bolivia
	"IC":         TimeSpanish, // Canary Islands
	"CL":         TimeSpanish, // Chile
	"CO":         TimeSpanish, // Colombia
	"CR":         TimeSpanish, // Costa Rica
	"CU":         TimeSpanish, // Cuba
	"DO":         TimeSpanish, // Dominican Republic
	"EC":         TimeSpanish, // Ecuador
	"SV":         TimeSpanish, // El Salvador
	"GT":         TimeSpanish, // Guatemala
	"HN":         TimeSpanish, // Honduras
	"MX":         TimeSpanish, // Mexico
	"NI":         TimeSpanish, // Nicaragua
	"PA":         TimeSpanish, // Panama
	"PY":         TimeSpanish, // Paraguay
	"PE":         TimeSpanish, // Peru
	"es-PH":      TimeSpanish, // Spanish - Philippines
	"PR":         TimeSpanish, // Puerto Rico
	"es-US":      TimeSpanish, // Spanish - USA
	"UY":         TimeSpanish, // Uruguay
	"VE":         TimeSpanish, // Venezuela
	"ES":         TimeSpanish, // Spain

	"pt":         TimePortuguese,
	"Portuguese": TimePortuguese,
	"Portugues":  TimePortuguese, //nolint:misspell
	"portuguese": TimePortuguese,
	"portugues":  TimePortuguese, //nolint:misspell
	"PT":         TimePortuguese, // Portugal
	"pt-PT":      TimePortuguese, // Portugal
	"BR":         TimePortuguese, // Brazil
	"pt-AO":      TimePortuguese, // Portuguese - Angola
	"pt-CV":      TimePortuguese, // Portuguese - Cape Verde
	"pt-GQ":      TimePortuguese, // Portuguese - Equatorial Guinea
	"pt-GW":      TimePortuguese, // Portuguese - Guinea-Bissau
	"pt-MO":      TimePortuguese, // Portuguese - Macao
	"pt-MZ":      TimePortuguese, // Portuguese - Mozambique

	"de":      TimeGerman,
	"German":  TimeGerman,
	"Deutsch": TimeGerman,
	"german":  TimeGerman,
	"deutsch": TimeGerman,
	"AT":      TimeGerman, // Austria
	"DE":      TimeGerman, // Germany
	"de-CH":   TimeGerman, // German - Switzerland
}

// TimeByName provides a mapping between names and the associated Time
// conversion tables.
//
// Times are identified by a range of different names. The ISO 3166 code for
// a country will return the name-to-time-parts mapping for the majority
// language in that country. Note that for some countries having multiple
// languages the other languages can be selected by prefixing the country
// code with the two-letter language code (as defined by ISO 639) and a dash;
// for instance, to select French in Canada 'fr-CA' may be
// used. Additionally, the common, English names of the languages may be used
// or just the ISO 639 codes.
//
// Note that this is NOT a complete set of languages.
//
// If the country code is not recognised a non-nil error is returned.
func TimeByName(name string) (Time, error) {
	l, ok := namedTimes[name]
	if !ok {
		return l, fmt.Errorf("unknown country code: %q", name)
	}

	return l, nil
}
