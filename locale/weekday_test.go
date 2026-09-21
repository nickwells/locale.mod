package locale_test

import (
	"testing"
	"time"

	"github.com/nickwells/locale.mod/locale"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestWeekdayByName(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		name       string
		expWeekday time.Weekday
	}{
		{
			ID:     testhelper.MkID("no name"),
			ExpErr: testhelper.MkExpErr(`unknown day of week: ""`),
		},
		{
			ID:     testhelper.MkID("nonesuch"),
			ExpErr: testhelper.MkExpErr(`unknown day of week: "nonesuch"`),
			name:   "nonesuch",
		},
		{
			ID:         testhelper.MkID("Monday"),
			name:       "Monday",
			expWeekday: time.Monday,
		},
		{
			ID:         testhelper.MkID("Mon"),
			name:       "Mon",
			expWeekday: time.Monday,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			wd, err := locale.WeekdayByName(tc.name)
			if !testhelper.CheckExpErr(t, err, tc) {
				return
			}

			if err != nil {
				return
			}

			testhelper.DiffInt(t, tc.IDStr(), "weekday", wd, tc.expWeekday)
		})
	}
}

func TestWeekdayByCountry(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		name        string
		countryName string
		expWeekday  time.Weekday
	}{
		{
			ID:     testhelper.MkID("no name, no country"),
			ExpErr: testhelper.MkExpErr(`unknown language: ""`),
		},
		{
			ID:          testhelper.MkID("no name, nonesuch"),
			ExpErr:      testhelper.MkExpErr(`unknown language: "nonesuch"`),
			countryName: "nonesuch",
		},
		{
			ID:          testhelper.MkID("no name, English"),
			ExpErr:      testhelper.MkExpErr(`unknown day of week: ""`),
			countryName: "English",
		},
		{
			ID:          testhelper.MkID("nonesuch, English"),
			ExpErr:      testhelper.MkExpErr(`unknown day of week: "nonesuch"`),
			name:        "nonesuch",
			countryName: "English",
		},
		{
			ID:          testhelper.MkID("Monday, English"),
			name:        "Monday",
			countryName: "English",
			expWeekday:  time.Monday,
		},
		{
			ID:          testhelper.MkID("Mon, English"),
			name:        "Mon",
			countryName: "English",
			expWeekday:  time.Monday,
		},
		{
			ID:          testhelper.MkID("Mon, French"),
			ExpErr:      testhelper.MkExpErr(`unknown day of week: "Mon"`),
			name:        "Mon",
			countryName: "French",
		},
		{
			ID:          testhelper.MkID("Lundi, FR"),
			name:        "Lundi",
			countryName: "FR",
			expWeekday:  time.Monday,
		},
		{
			ID:          testhelper.MkID("Lundi, Français"),
			name:        "Lundi",
			countryName: "Français",
			expWeekday:  time.Monday,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			wd, err := locale.WeekdayByNameByCountry(
				tc.name, tc.countryName)
			if !testhelper.CheckExpErr(t, err, tc) {
				return
			}

			if err != nil {
				return
			}

			testhelper.DiffInt(t, tc.IDStr(), "weekday", wd, tc.expWeekday)
		})
	}
}
