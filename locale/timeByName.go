package locale

import (
	"fmt"
	"os"
	"strings"
)

var namedTimes = map[string]Time{
	"en":      TimeEnglish,
	"English": TimeEnglish,
	"english": TimeEnglish,
	"GB":      TimeEnglish, // Great Britain
	"en-GB":   TimeEnglish, // Great Britain
	"en_GB":   TimeEnglish, // Great Britain
	"US":      TimeEnglish, // United States of America
	"en-US":   TimeEnglish, // United States of America
	"en_US":   TimeEnglish, // United States of America
	"CA":      TimeEnglish, // Canada
	"en-CA":   TimeEnglish, // Canada
	"en_CA":   TimeEnglish, // Canada
	"AU":      TimeEnglish, // Australia
	"en-AU":   TimeEnglish, // Australia
	"en_AU":   TimeEnglish, // Australia
	"NZ":      TimeEnglish, // New Zealand
	"en-NZ":   TimeEnglish, // New Zealand
	"en_NZ":   TimeEnglish, // New Zealand
	"NG":      TimeEnglish, // Nigeria
	"en-NG":   TimeEnglish, // Nigeria
	"en_NG":   TimeEnglish, // Nigeria
	"SG":      TimeEnglish, // Singapore
	"en-SG":   TimeEnglish, // Singapore
	"en_SG":   TimeEnglish, // Singapore
	"en-PH":   TimeEnglish, // English - Philippines
	"en_PH":   TimeEnglish, // English - Philippines
	"KE":      TimeEnglish, // Kenya
	"en-KE":   TimeEnglish, // Kenya
	"en_KE":   TimeEnglish, // Kenya
	"UG":      TimeEnglish, // Uganda
	"en-UG":   TimeEnglish, // Uganda
	"en_UG":   TimeEnglish, // Uganda
	"GH":      TimeEnglish, // Ghana
	"en-GH":   TimeEnglish, // Ghana
	"en_GH":   TimeEnglish, // Ghana
	"MW":      TimeEnglish, // Malawi
	"en-MW":   TimeEnglish, // Malawi
	"en_MW":   TimeEnglish, // Malawi
	"ZM":      TimeEnglish, // Zambia
	"en-ZM":   TimeEnglish, // Zambia
	"en_ZM":   TimeEnglish, // Zambia
	"ZW":      TimeEnglish, // Zimbabwe
	"en-ZW":   TimeEnglish, // Zimbabwe
	"en_ZW":   TimeEnglish, // Zimbabwe
	"en-SD":   TimeEnglish, // English - Sudan
	"en_SD":   TimeEnglish, // English - Sudan
	"SS":      TimeEnglish, // South Sudan
	"en-SS":   TimeEnglish, // South Sudan
	"en_SS":   TimeEnglish, // South Sudan
	"ZA":      TimeEnglish, // South Africa
	"en-ZA":   TimeEnglish, // South Africa
	"en_ZA":   TimeEnglish, // South Africa

	"fr":       TimeFrench,
	"French":   TimeFrench,
	"Francais": TimeFrench,
	"Français": TimeFrench,
	"french":   TimeFrench,
	"francais": TimeFrench,
	"français": TimeFrench,
	"FR":       TimeFrench, // France
	"fr-FR":    TimeFrench, // France
	"fr_FR":    TimeFrench, // France
	"fr-CA":    TimeFrench, // French - Canada
	"fr_CA":    TimeFrench, // French - Canada
	"fr-CH":    TimeFrench, // French - Switzerland
	"fr_CH":    TimeFrench, // French - Switzerland
	"HT":       TimeFrench, // Haiti
	"fr-HT":    TimeFrench, // Haiti
	"fr_HT":    TimeFrench, // Haiti

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
	"es-AR":      TimeSpanish, // Argentina
	"es_AR":      TimeSpanish, // Argentina
	"BO":         TimeSpanish, // Bolivia
	"es-BO":      TimeSpanish, // Bolivia
	"es_BO":      TimeSpanish, // Bolivia
	"IC":         TimeSpanish, // Canary Islands
	"es-IC":      TimeSpanish, // Canary Islands
	"es_IC":      TimeSpanish, // Canary Islands
	"CL":         TimeSpanish, // Chile
	"es-CL":      TimeSpanish, // Chile
	"es_CL":      TimeSpanish, // Chile
	"CO":         TimeSpanish, // Colombia
	"es-CO":      TimeSpanish, // Colombia
	"es_CO":      TimeSpanish, // Colombia
	"CR":         TimeSpanish, // Costa Rica
	"es-CR":      TimeSpanish, // Costa Rica
	"es_CR":      TimeSpanish, // Costa Rica
	"CU":         TimeSpanish, // Cuba
	"es-CU":      TimeSpanish, // Cuba
	"es_CU":      TimeSpanish, // Cuba
	"DO":         TimeSpanish, // Dominican Republic
	"es-DO":      TimeSpanish, // Dominican Republic
	"es_DO":      TimeSpanish, // Dominican Republic
	"EC":         TimeSpanish, // Ecuador
	"es-EC":      TimeSpanish, // Ecuador
	"es_EC":      TimeSpanish, // Ecuador
	"SV":         TimeSpanish, // El Salvador
	"es-SV":      TimeSpanish, // El Salvador
	"es_SV":      TimeSpanish, // El Salvador
	"GT":         TimeSpanish, // Guatemala
	"es-GT":      TimeSpanish, // Guatemala
	"es_GT":      TimeSpanish, // Guatemala
	"HN":         TimeSpanish, // Honduras
	"es-HN":      TimeSpanish, // Honduras
	"es_HN":      TimeSpanish, // Honduras
	"MX":         TimeSpanish, // Mexico
	"es-MX":      TimeSpanish, // Mexico
	"es_MX":      TimeSpanish, // Mexico
	"NI":         TimeSpanish, // Nicaragua
	"es-NI":      TimeSpanish, // Nicaragua
	"es_NI":      TimeSpanish, // Nicaragua
	"PA":         TimeSpanish, // Panama
	"es-PA":      TimeSpanish, // Panama
	"es_PA":      TimeSpanish, // Panama
	"PY":         TimeSpanish, // Paraguay
	"es-PY":      TimeSpanish, // Paraguay
	"es_PY":      TimeSpanish, // Paraguay
	"PE":         TimeSpanish, // Peru
	"es-PE":      TimeSpanish, // Peru
	"es_PE":      TimeSpanish, // Peru
	"es-PH":      TimeSpanish, // Spanish - Philippines
	"es_PH":      TimeSpanish, // Spanish - Philippines
	"PR":         TimeSpanish, // Puerto Rico
	"es-PR":      TimeSpanish, // Puerto Rico
	"es_PR":      TimeSpanish, // Puerto Rico
	"es-US":      TimeSpanish, // Spanish - USA
	"es_US":      TimeSpanish, // Spanish - USA
	"UY":         TimeSpanish, // Uruguay
	"es-UY":      TimeSpanish, // Uruguay
	"es_UY":      TimeSpanish, // Uruguay
	"VE":         TimeSpanish, // Venezuela
	"es-VE":      TimeSpanish, // Venezuela
	"es_VE":      TimeSpanish, // Venezuela
	"ES":         TimeSpanish, // Spain
	"es-ES":      TimeSpanish, // Spain
	"es_ES":      TimeSpanish, // Spain

	"pt":         TimePortuguese,
	"Portuguese": TimePortuguese,
	"Portugues":  TimePortuguese, //nolint:misspell
	"portuguese": TimePortuguese,
	"portugues":  TimePortuguese, //nolint:misspell
	"PT":         TimePortuguese, // Portugal
	"pt-PT":      TimePortuguese, // Portugal
	"pt_PT":      TimePortuguese, // Portugal
	"BR":         TimePortuguese, // Brazil
	"pt-BR":      TimePortuguese, // Brazil
	"pt_BR":      TimePortuguese, // Brazil
	"pt-AO":      TimePortuguese, // Portuguese - Angola
	"pt_AO":      TimePortuguese, // Portuguese - Angola
	"pt-CV":      TimePortuguese, // Portuguese - Cape Verde
	"pt_CV":      TimePortuguese, // Portuguese - Cape Verde
	"pt-GQ":      TimePortuguese, // Portuguese - Equatorial Guinea
	"pt_GQ":      TimePortuguese, // Portuguese - Equatorial Guinea
	"pt-GW":      TimePortuguese, // Portuguese - Guinea-Bissau
	"pt_GW":      TimePortuguese, // Portuguese - Guinea-Bissau
	"pt-MO":      TimePortuguese, // Portuguese - Macao
	"pt_MO":      TimePortuguese, // Portuguese - Macao
	"pt-MZ":      TimePortuguese, // Portuguese - Mozambique
	"pt_MZ":      TimePortuguese, // Portuguese - Mozambique

	"de":      TimeGerman,
	"German":  TimeGerman,
	"Deutsch": TimeGerman,
	"german":  TimeGerman,
	"deutsch": TimeGerman,
	"AT":      TimeGerman, // Austria
	"de-AT":   TimeGerman, // Austria
	"de_AT":   TimeGerman, // Austria
	"DE":      TimeGerman, // Germany
	"de-DE":   TimeGerman, // Germany
	"de_DE":   TimeGerman, // Germany
	"de-CH":   TimeGerman, // German - Switzerland
	"de_CH":   TimeGerman, // German - Switzerland

	"hi":     TimeHindi,
	"हिन्दी": TimeHindi,
	"Hindi":  TimeHindi,
	"hindi":  TimeHindi,
	"IN":     TimeHindi, // India
	"hi-IN":  TimeHindi, // India
	"hi_IN":  TimeHindi, // India

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
	"zh-CH":    TimeChinese, // People's Republic of China
	"zh_CH":    TimeChinese, // People's Republic of China
	"TW":       TimeChinese, // Republic of China (Taiwan)
	"zh-TW":    TimeChinese, // Republic of China (Taiwan)
	"zh_TW":    TimeChinese, // Republic of China (Taiwan)
	"zh-SG":    TimeChinese, // Chinese - Singapore
	"zh_SG":    TimeChinese, // Chinese - Singapore

	"ja":       TimeJapanese,
	"日本語":      TimeJapanese,
	"Nihongo":  TimeJapanese,
	"nihongo":  TimeJapanese,
	"Japanese": TimeJapanese,
	"japanese": TimeJapanese,
	"JP":       TimeJapanese, // Japan
	"ja-JP":    TimeJapanese, // Japan
	"ja_JP":    TimeJapanese, // Japan

	"ms":    TimeMalay,
	"Malay": TimeMalay,
	"malay": TimeMalay,
	"MY":    TimeMalay, // Malaysia
	"ms-MY": TimeMalay, // Malaysia
	"ms_MY": TimeMalay, // Malaysia
	"ID":    TimeMalay, // Indonesia
	"ms-ID": TimeMalay, // Indonesia
	"ms_ID": TimeMalay, // Indonesia
	"ms-SG": TimeMalay, // Malay - Singapore
	"ms_SG": TimeMalay, // Malay - Singapore

	"ar":              TimeArabic,
	"اَلْعَرَبِيَّةُ": TimeArabic,
	"al-ʻArabīyah":    TimeArabic,
	"al-ʻArabīyyah":   TimeArabic,
	"al-Arabiyah":     TimeArabic,
	"al-Arabiyyah":    TimeArabic,
	"Arabic":          TimeArabic,
	"arabic":          TimeArabic,
	"ar-DZ":           TimeArabic, // Arabic - Algeria
	"ar_DZ":           TimeArabic, // Arabic - Algeria
	"BH":              TimeArabic, // Bahrain
	"ar-BH":           TimeArabic, // Bahrain
	"ar_BH":           TimeArabic, // Bahrain
	"ar-TD":           TimeArabic, // Arabic - Chad
	"ar_TD":           TimeArabic, // Arabic - Chad
	"ar-KM":           TimeArabic, // Arabic - Comoros
	"ar_KM":           TimeArabic, // Arabic - Comoros
	"ar-DJ":           TimeArabic, // Arabic - Djibouti
	"ar_DJ":           TimeArabic, // Arabic - Djibouti
	"EG":              TimeArabic, // Egypt
	"ar-EG":           TimeArabic, // Egypt
	"ar_EG":           TimeArabic, // Egypt
	"ar-IQ":           TimeArabic, // Arabic - Iraq
	"ar_IQ":           TimeArabic, // Arabic - Iraq
	"JO":              TimeArabic, // Jordan
	"ar-JO":           TimeArabic, // Jordan
	"ar_JO":           TimeArabic, // Jordan
	"KW":              TimeArabic, // Kuwait
	"ar-KW":           TimeArabic, // Kuwait
	"ar_KW":           TimeArabic, // Kuwait
	"LB":              TimeArabic, // Lebanon
	"ar-LB":           TimeArabic, // Lebanon
	"ar_LB":           TimeArabic, // Lebanon
	"LY":              TimeArabic, // Libya
	"ar-LY":           TimeArabic, // Libya
	"ar_LY":           TimeArabic, // Libya
	"ar-ML":           TimeArabic, // Arabic - Mali
	"ar_ML":           TimeArabic, // Arabic - Mali
	"MR":              TimeArabic, // Mauritania
	"ar-MR":           TimeArabic, // Mauritania
	"ar_MR":           TimeArabic, // Mauritania
	"ar-MA":           TimeArabic, // Arabic - Morocco
	"ar_MA":           TimeArabic, // Arabic - Morocco
	"OM":              TimeArabic, // Oman
	"ar-OM":           TimeArabic, // Oman
	"ar_OM":           TimeArabic, // Oman
	"PS":              TimeArabic, // Palestine
	"ar-PS":           TimeArabic, // Palestine
	"ar_PS":           TimeArabic, // Palestine
	"QA":              TimeArabic, // Qatar
	"ar-QA":           TimeArabic, // Qatar
	"ar_QA":           TimeArabic, // Qatar
	"SA":              TimeArabic, // Saudi Arabia
	"ar-SA":           TimeArabic, // Saudi Arabia
	"ar_SA":           TimeArabic, // Saudi Arabia
	"ar-SO":           TimeArabic, // Arabic - Somalia
	"ar_SO":           TimeArabic, // Arabic - Somalia
	"ar-SD":           TimeArabic, // Arabic - Sudan
	"ar_SD":           TimeArabic, // Arabic - Sudan
	"SY":              TimeArabic, // Syria
	"ar-SY":           TimeArabic, // Syria
	"ar_SY":           TimeArabic, // Syria
	"TN":              TimeArabic, // Tunisia
	"ar-TN":           TimeArabic, // Tunisia
	"ar_TN":           TimeArabic, // Tunisia
	"AE":              TimeArabic, // United Arab Emirates
	"ar-AE":           TimeArabic, // United Arab Emirates
	"ar_AE":           TimeArabic, // United Arab Emirates
	"YE":              TimeArabic, // Yemen
	"ar-YE":           TimeArabic, // Yemen
	"ar_YE":           TimeArabic, // Yemen
}

// TimeByName provides a mapping between names and the associated Time
// conversion tables.
//
// Times are identified by a range of different names. The ISO 3166-2 code
// for a country will return the name-to-time-parts mapping for the majority
// language in that country. Note that for some countries having multiple
// languages the other languages can be selected by prefixing the country
// code with the two-letter language code (as defined by ISO 639) and a dash
// or underscore (both variants should be present); for instance, to select
// French in Canada 'fr-CA' may be used. Additionally, the common, English
// names of the languages may be used or just the ISO 639 codes.
//
// Note that this is NOT a complete set of languages.
//
// If the country code is not recognised a non-nil error is returned.
func TimeByName(name string) (Time, error) {
	t, ok := namedTimes[name]
	if !ok {
		return t, fmt.Errorf("unknown language: %q", name)
	}

	return t, nil
}

// TimeByEnv finds a Time conversion table based on the values in the
// environment variables LANGUAGE and LANG. The LANGUAGE variable holds a
// colon-separated list of languages - both in the form of an ISO 639 code
// followed by an underscore and the ISO 3166-2 code and in the form of just
// the ISO 639 code. These will be searched for in order and the first match
// will be returned. If none of these give a value then the value in the LANG
// variable will be used (any trailing part after a dot will be
// discarded). If neither of these results in a Time value then the default
// value of TimeEnglish will be used.
func TimeByEnv() Time {
	languages := []string{}

	languageEnvvar := os.Getenv("LANGUAGE")
	if languageEnvvar != "" {
		languages = strings.Split(languageEnvvar, ":")
	}

	langEnvvar := os.Getenv("LANG")

	lang, _, _ := strings.Cut(langEnvvar, ".")
	if lang != "" {
		languages = append(languages, lang)
	}

	for _, name := range languages {
		if t, ok := namedTimes[name]; ok {
			return t
		}
	}

	return TimeEnglish
}
