package locale_test

import (
	"testing"
	"time"

	"github.com/nickwells/locale.mod/locale"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

var goodToWeekday = map[string]time.Weekday{
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
	"Sunday":    time.Sunday,
}

var badToWeekdayBadDay = map[string]time.Weekday{
	"Monday":    time.Weekday(99),
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
	"Sunday":    time.Sunday,
}

var badToWeekdayMissingMondayFriday = map[string]time.Weekday{
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Saturday":  time.Saturday,
	"Sunday":    time.Sunday,
}

var badToWeekdayNoMonday = map[string]time.Weekday{
	"Mon":       time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,
	"Sunday":    time.Sunday,
}

var goodToMonth = map[string]time.Month{
	"January":   time.January,
	"February":  time.February,
	"March":     time.March,
	"April":     time.April,
	"May":       time.May,
	"June":      time.June,
	"July":      time.July,
	"August":    time.August,
	"September": time.September,
	"October":   time.October,
	"November":  time.November,
	"December":  time.December,
}

var badToMonthBadMonth = map[string]time.Month{
	"January":   time.Month(99),
	"February":  time.February,
	"March":     time.March,
	"April":     time.April,
	"May":       time.May,
	"June":      time.June,
	"July":      time.July,
	"August":    time.August,
	"September": time.September,
	"October":   time.October,
	"November":  time.November,
	"December":  time.December,
}

var badToMonthMissingMayJuly = map[string]time.Month{
	"January":   time.January,
	"February":  time.February,
	"March":     time.March,
	"April":     time.April,
	"June":      time.June,
	"August":    time.August,
	"September": time.September,
	"October":   time.October,
	"November":  time.November,
	"December":  time.December,
}

var badToMonthNoJanuary = map[string]time.Month{
	"Jan":       time.January,
	"February":  time.February,
	"March":     time.March,
	"April":     time.April,
	"May":       time.May,
	"June":      time.June,
	"July":      time.July,
	"August":    time.August,
	"September": time.September,
	"October":   time.October,
	"November":  time.November,
	"December":  time.December,
}

var goodFromWeekday = map[time.Weekday]string{
	time.Monday:    "Monday",
	time.Tuesday:   "Tuesday",
	time.Wednesday: "Wednesday",
	time.Thursday:  "Thursday",
	time.Friday:    "Friday",
	time.Saturday:  "Saturday",
	time.Sunday:    "Sunday",
}

var badFromWeekdayBadDay = map[time.Weekday]string{
	time.Weekday(99): "Monday",
	time.Tuesday:     "Tuesday",
	time.Wednesday:   "Wednesday",
	time.Thursday:    "Thursday",
	time.Friday:      "Friday",
	time.Saturday:    "Saturday",
	time.Sunday:      "Sunday",
}

var badFromWeekdayDupDay = map[time.Weekday]string{
	time.Monday:    "Monday",
	time.Tuesday:   "Monday",
	time.Wednesday: "Wednesday",
	time.Thursday:  "Thursday",
	time.Friday:    "Friday",
	time.Saturday:  "Saturday",
	time.Sunday:    "Sunday",
}

var badFromWeekdayMissingMondayFriday = map[time.Weekday]string{
	time.Tuesday:   "Monday",
	time.Wednesday: "Wednesday",
	time.Thursday:  "Thursday",
	time.Saturday:  "Saturday",
	time.Sunday:    "Sunday",
}

var badFromWeekdayMonTueSwapped = map[time.Weekday]string{
	time.Monday:    "Tuesday",
	time.Tuesday:   "Monday",
	time.Wednesday: "Wednesday",
	time.Thursday:  "Thursday",
	time.Friday:    "Friday",
	time.Saturday:  "Saturday",
	time.Sunday:    "Sunday",
}

var goodFromMonth = map[time.Month]string{
	time.January:   "January",
	time.February:  "February",
	time.March:     "March",
	time.April:     "April",
	time.May:       "May",
	time.June:      "June",
	time.July:      "July",
	time.August:    "August",
	time.September: "September",
	time.October:   "October",
	time.November:  "November",
	time.December:  "December",
}

var badFromMonthBadMonth = map[time.Month]string{
	time.Month(99): "January",
	time.February:  "February",
	time.March:     "March",
	time.April:     "April",
	time.May:       "May",
	time.June:      "June",
	time.July:      "July",
	time.August:    "August",
	time.September: "September",
	time.October:   "October",
	time.November:  "November",
	time.December:  "December",
}

var badFromMonthDupMonth = map[time.Month]string{
	time.January:   "January",
	time.February:  "January",
	time.March:     "March",
	time.April:     "April",
	time.May:       "May",
	time.June:      "June",
	time.July:      "July",
	time.August:    "August",
	time.September: "September",
	time.October:   "October",
	time.November:  "November",
	time.December:  "December",
}

var badFromMonthMissingMayJuly = map[time.Month]string{
	time.January:   "January",
	time.February:  "February",
	time.March:     "March",
	time.April:     "April",
	time.June:      "June",
	time.August:    "August",
	time.September: "September",
	time.October:   "October",
	time.November:  "November",
	time.December:  "December",
}

var badFromMonthJanFebSwapped = map[time.Month]string{
	time.January:   "February",
	time.February:  "January",
	time.March:     "March",
	time.April:     "April",
	time.May:       "May",
	time.June:      "June",
	time.July:      "July",
	time.August:    "August",
	time.September: "September",
	time.October:   "October",
	time.November:  "November",
	time.December:  "December",
}

var locTest1 = locale.MakeTimeOrPanic(
	goodToWeekday,
	goodToMonth,
	goodFromWeekday,
	goodFromMonth,
)

func TestMonthNames(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		ct       locale.Time
		expNames []string
	}{
		{
			ID: testhelper.MkID(""),
			ct: locTest1,
			expNames: []string{
				"January",
				"February",
				"March",
				"April",
				"May",
				"June",
				"July",
				"August",
				"September",
				"October",
				"November",
				"December",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			names := tc.ct.MonthNames()
			testhelper.DiffSlice(t, tc.IDStr(), "month", names, tc.expNames)
		})
	}
}

func TestWeekdayNames(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		ct       locale.Time
		expNames []string
	}{
		{
			ID: testhelper.MkID(""),
			ct: locTest1,
			expNames: []string{
				"Sunday",
				"Monday",
				"Tuesday",
				"Wednesday",
				"Thursday",
				"Friday",
				"Saturday",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			names := tc.ct.WeekdayNames()
			testhelper.DiffSlice(t, tc.IDStr(), "weekday", names, tc.expNames)
		})
	}
}

func TestToMonth(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		name     string
		l        locale.Time
		expMonth time.Month
	}{
		{
			ID:     testhelper.MkID("no name, English conversion table"),
			ExpErr: testhelper.MkExpErr(`unknown month: ""`),
			l:      locale.TimeEnglish,
		},
		{
			ID:     testhelper.MkID("nonesuch, English conversion table"),
			ExpErr: testhelper.MkExpErr(`unknown month: "nonesuch"`),
			name:   "nonesuch",
			l:      locale.TimeEnglish,
		},
		{
			ID:       testhelper.MkID("January, English conversion table"),
			name:     "January",
			l:        locale.TimeEnglish,
			expMonth: time.January,
		},
		{
			ID:       testhelper.MkID("Jan, English conversion table"),
			name:     "Jan",
			l:        locale.TimeEnglish,
			expMonth: time.January,
		},
		{
			ID:       testhelper.MkID("1, English conversion table"),
			name:     "1",
			l:        locale.TimeEnglish,
			expMonth: time.January,
		},
		{
			ID:       testhelper.MkID("01, English conversion table"),
			name:     "01",
			l:        locale.TimeEnglish,
			expMonth: time.January,
		},
		{
			ID:     testhelper.MkID("Feb, French conversion table"),
			ExpErr: testhelper.MkExpErr(`unknown month: "Feb"`),
			name:   "Feb",
			l:      locale.TimeFrench,
		},
		{
			ID:       testhelper.MkID("Janvier, French conversion table"),
			name:     "Janvier",
			l:        locale.TimeFrench,
			expMonth: time.January,
		},
		{
			ID:       testhelper.MkID("1, French conversion table"),
			name:     "1",
			l:        locale.TimeFrench,
			expMonth: time.January,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			m, err := tc.l.ToMonth(tc.name)
			if !testhelper.CheckExpErr(t, err, tc) {
				return
			}

			if err != nil {
				return
			}

			testhelper.DiffString(t, tc.IDStr(), "month",
				m.String(), tc.expMonth.String())
		})
	}
}

func TestToWeekday(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		name       string
		l          locale.Time
		expWeekday time.Weekday
	}{
		{
			ID:     testhelper.MkID("no name, English conversion table"),
			ExpErr: testhelper.MkExpErr(`unknown day of week: ""`),
			l:      locale.TimeEnglish,
		},
		{
			ID:     testhelper.MkID("nonesuch, English conversion table"),
			ExpErr: testhelper.MkExpErr(`unknown day of week: "nonesuch"`),
			name:   "nonesuch",
			l:      locale.TimeEnglish,
		},
		{
			ID:         testhelper.MkID("Monday, English conversion table"),
			name:       "Monday",
			l:          locale.TimeEnglish,
			expWeekday: time.Monday,
		},
		{
			ID:         testhelper.MkID("Mon, English conversion table"),
			name:       "Mon",
			l:          locale.TimeEnglish,
			expWeekday: time.Monday,
		},
		{
			ID:     testhelper.MkID("Mon, French conversion table"),
			ExpErr: testhelper.MkExpErr(`unknown day of week: "Mon"`),
			name:   "Mon",
			l:      locale.TimeFrench,
		},
		{
			ID:         testhelper.MkID("Lundi, French conversion table"),
			name:       "Lundi",
			l:          locale.TimeFrench,
			expWeekday: time.Monday,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			w, err := tc.l.ToWeekday(tc.name)
			if !testhelper.CheckExpErr(t, err, tc) {
				return
			}

			if err != nil {
				return
			}

			testhelper.DiffString(t, tc.IDStr(), "weekday",
				w.String(), tc.expWeekday.String())
		})
	}
}

func TestMakeLocaleOrPanic(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpPanic
		toWeekday   map[string]time.Weekday
		toMonth     map[string]time.Month
		fromWeekday map[time.Weekday]string
		fromMonth   map[time.Month]string
	}{
		{
			ID:          testhelper.MkID("all good"),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad toWeekday - bad day"),
			ExpPanic: testhelper.MkExpPanic(
				`toWeekday maps "Monday"` +
					" to %!Weekday(99) (which is not a valid weekday)"),
			toWeekday:   badToWeekdayBadDay,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad toWeekday - missing Monday and Friday"),
			ExpPanic: testhelper.MkExpPanic(
				"toWeekday only has mappings to 5 days;" +
					" missing: Monday, Friday"),
			toWeekday:   badToWeekdayMissingMondayFriday,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad toMonth - bad month"),
			ExpPanic: testhelper.MkExpPanic(
				`toMonth maps "January"` +
					" to %!Month(99) (which is not a valid month)"),
			toWeekday:   goodToWeekday,
			toMonth:     badToMonthBadMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad toMonth - missing May and July"),
			ExpPanic: testhelper.MkExpPanic(
				"toMonth only has mappings to 10 months;" +
					" missing: May, July"),
			toWeekday:   goodToWeekday,
			toMonth:     badToMonthMissingMayJuly,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad fromWeekday - bad day"),
			ExpPanic: testhelper.MkExpPanic(
				"fromWeekday maps %!Weekday(99)" +
					" (which is not a valid weekday)" +
					` to "Monday"`),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: badFromWeekdayBadDay,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad fromWeekday - duplicate day"),
			ExpPanic: testhelper.MkExpPanic(
				"duplicate detected; fromWeekday maps both" +
					" Monday and Tuesday" +
					` to "Monday"`),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: badFromWeekdayDupDay,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad fromWeekday - missing Monday and Friday"),
			ExpPanic: testhelper.MkExpPanic(
				"fromWeekday only has mappings to 5 weekdays;" +
					" missing: Monday, Friday"),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: badFromWeekdayMissingMondayFriday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad fromMonth - bad month"),
			ExpPanic: testhelper.MkExpPanic(
				"fromMonth maps %!Month(99)" +
					" (which is not a valid month)" +
					` to "January"`),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   badFromMonthBadMonth,
		},
		{
			ID: testhelper.MkID("bad fromMonth - duplicate month"),
			ExpPanic: testhelper.MkExpPanic(
				"duplicate detected; fromMonth maps both" +
					" January and February" +
					` to "January"`),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   badFromMonthDupMonth,
		},
		{
			ID: testhelper.MkID("bad fromMonth - missing May and July"),
			ExpPanic: testhelper.MkExpPanic(
				"fromMonth only has mappings to 10 months;" +
					" missing: May, July"),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   badFromMonthMissingMayJuly,
		},
		{
			ID: testhelper.MkID("bad toMonth - has Jan but no January"),
			ExpPanic: testhelper.MkExpPanic(
				`bad time.Month maps: from[time.January] gives "January"` +
					` but to["January"] is not found`),
			toWeekday:   goodToWeekday,
			toMonth:     badToMonthNoJanuary,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad toMonth - has Jan but no January"),
			ExpPanic: testhelper.MkExpPanic(
				`bad time.Month maps: from[time.January] gives "January"` +
					` but to["January"] is not found`),
			toWeekday:   goodToWeekday,
			toMonth:     badToMonthNoJanuary,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad toWeekday - has Mon but no Monday"),
			ExpPanic: testhelper.MkExpPanic(
				`bad time.Weekday maps: from[time.Monday] gives "Monday"` +
					` but to["Monday"] is not found`),
			toWeekday:   badToWeekdayNoMonday,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   goodFromMonth,
		},
		{
			ID: testhelper.MkID("bad fromMonth - January gives February"),
			ExpPanic: testhelper.MkExpPanic(
				`bad time.Month maps: from[time.January] gives "February"` +
					` but to["February"] gives time.February not time.January`),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: goodFromWeekday,
			fromMonth:   badFromMonthJanFebSwapped,
		},
		{
			ID: testhelper.MkID("bad fromWeekday - Monday gives Tuesday"),
			ExpPanic: testhelper.MkExpPanic(
				`bad time.Weekday maps: from[time.Monday] gives "Tuesday"` +
					` but to["Tuesday"] gives time.Tuesday not time.Monday`),
			toWeekday:   goodToWeekday,
			toMonth:     goodToMonth,
			fromWeekday: badFromWeekdayMonTueSwapped,
			fromMonth:   goodFromMonth,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			panicked, panicVal := testhelper.PanicSafe(
				func() {
					locale.MakeTimeOrPanic(
						tc.toWeekday, tc.toMonth,
						tc.fromWeekday, tc.fromMonth)
				},
			)

			testhelper.CheckExpPanicError(t, panicked, panicVal, tc)
		})
	}
}
