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
	"en-CA":   TimeEnglish, // English - Canada
	"AU":      TimeEnglish, // Australia
	"NZ":      TimeEnglish, // New Zealand
	"NG":      TimeEnglish, // Nigeria
	"SG":      TimeEnglish, // Singapore
	"en-SG":   TimeEnglish, // English - Singapore
	"en-PH":   TimeEnglish, // English - Philippines
	"KE":      TimeEnglish, // Kenya
	"UG":      TimeEnglish, // Uganda
	"GH":      TimeEnglish, // Ghana
	"MW":      TimeEnglish, // Malawi
	"ZM":      TimeEnglish, // Zambia
	"ZW":      TimeEnglish, // Zimbabwe
	"en-SD":   TimeEnglish, // English - Sudan
	"SS":      TimeEnglish, // South Sudan
	"ZA":      TimeEnglish, // South Africa

	"fr":       TimeFrench,
	"French":   TimeFrench,
	"Francais": TimeFrench,
	"Français": TimeFrench,
	"french":   TimeFrench,
	"francais": TimeFrench,
	"français": TimeFrench,
	"FR":       TimeFrench, // France
	"fr-CA":    TimeFrench, // French - Canada
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

	"hi":     TimeHindi,
	"हिन्दी": TimeHindi,
	"Hindi":  TimeHindi,
	"hindi":  TimeHindi,
	"IN":     TimeHindi, // India

	"汉语":       TimeChinese, // simplified Chinese
	"漢語":       TimeChinese, // traditional Chinese
	"Hànyǔ":    TimeChinese,
	"Hanyu":    TimeChinese,
	"hànyǔ":    TimeChinese,
	"hanyu":    TimeChinese,
	"中文":       TimeChinese,
	"Zhōngwén": TimeChinese,
	"Zhongwen": TimeChinese,
	"zhōngwén": TimeChinese,
	"zhongwen": TimeChinese,
	"zh":       TimeChinese,
	"Chinese":  TimeChinese,
	"chinese":  TimeChinese,
	"CH":       TimeChinese, // People's Republic of China
	"TW":       TimeChinese, // Republic of China (Taiwan)
	"zh-SG":    TimeChinese, // Chinese - Singapore

	"ja":       TimeJapanese,
	"日本語":      TimeJapanese,
	"Nihongo":  TimeJapanese,
	"nihongo":  TimeJapanese,
	"Japanese": TimeJapanese,
	"japanese": TimeJapanese,
	"JP":       TimeJapanese, // Japan

	"ms":    TimeMalay,
	"Malay": TimeMalay,
	"malay": TimeMalay,
	"MY":    TimeMalay, // Malaysia
	"ID":    TimeMalay, // Indonesia
	"ms-SG": TimeMalay, // Malay - Singapore

	"ar":              TimeArabic,
	"اَلْعَرَبِيَّةُ": TimeArabic,
	"al-ʻArabīyah":    TimeArabic,
	"al-ʻArabīyyah":   TimeArabic,
	"al-Arabiyah":     TimeArabic,
	"al-Arabiyyah":    TimeArabic,
	"Arabic":          TimeArabic,
	"arabic":          TimeArabic,
	"ar-DZ":           TimeArabic, // Arabic - Algeria
	"BH":              TimeArabic, // Bahrain
	"ar-TD":           TimeArabic, // Arabic - Chad
	"ar-KM":           TimeArabic, // Arabic - Comoros
	"ar-DJ":           TimeArabic, // Arabic - Djibouti
	"EG":              TimeArabic, // Egypt
	"ar-IQ":           TimeArabic, // Arabic - Iraq
	"JO":              TimeArabic, // Jordan
	"KW":              TimeArabic, // Kuwait
	"LB":              TimeArabic, // Lebanon
	"LY":              TimeArabic, // Libya
	"ar-ML":           TimeArabic, // Arabic - Mali
	"MR":              TimeArabic, // Mauritania
	"ar-MA":           TimeArabic, // Arabic - Morocco
	"OM":              TimeArabic, // Oman
	"PS":              TimeArabic, // Palestine
	"QA":              TimeArabic, // Qatar
	"SA":              TimeArabic, // Saudi Arabia
	"ar-SO":           TimeArabic, // Arabic - Somalia
	"ar-SD":           TimeArabic, // Arabic - Sudan
	"SY":              TimeArabic, // Syria
	"TN":              TimeArabic, // Tunisia
	"AE":              TimeArabic, // United Arab Emirates
	"YE":              TimeArabic, // Yemen
}

// TimeByName provides a mapping between names and the associated Time
// conversion tables.
//
// Times are identified by a range of different names. The ISO 3166-2 code for
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
