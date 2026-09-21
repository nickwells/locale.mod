package locale

import "time"

// TimeMalay provides the Time conversion tables for the Malay language. Only
// Latin (Rumi) script forms are given; there are no entries for the Arab
// Melayu (Jawi) script forms as the Latim script is now official in
// Malaysia, Indonesia and Singapore.
var TimeMalay = MakeTimeOrPanic(
	map[string]time.Weekday{
		"Ahad":   time.Sunday,
		"Minggu": time.Sunday,
		"Isnin":  time.Monday,
		"Senin":  time.Monday,
		"Selasa": time.Tuesday,
		"Rabu":   time.Wednesday,
		"Kamis":  time.Thursday,
		"Khamis": time.Thursday,
		"Jumat":  time.Friday,
		"Jumaat": time.Friday,
		"Sabtu":  time.Saturday,
	},
	map[string]time.Month{
		"Januari":   time.January,
		"Februari":  time.February,
		"Mac":       time.March,
		"April":     time.April,
		"Mei":       time.May,
		"Jun":       time.June,
		"Julai":     time.July,
		"Ogos":      time.August,
		"September": time.September,
		"Oktober":   time.October,
		"November":  time.November,
		"Disember":  time.December,
	},
	map[time.Weekday]string{
		time.Monday:    "Isnin",
		time.Tuesday:   "Selasa",
		time.Wednesday: "Rabu",
		time.Thursday:  "Kamis",
		time.Friday:    "Jumat",
		time.Saturday:  "Sabtu",
		time.Sunday:    "Ahad",
	},
	map[time.Month]string{
		time.January:   "Januari",
		time.February:  "Februari",
		time.March:     "Mac",
		time.April:     "April",
		time.May:       "Mei",
		time.June:      "Jun",
		time.July:      "Julai",
		time.August:    "Ogos",
		time.September: "September",
		time.October:   "Oktober",
		time.November:  "November",
		time.December:  "Disember",
	},
)
