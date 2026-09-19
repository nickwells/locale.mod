package locale_test

import (
	"testing"
	"time"

	"github.com/nickwells/locale.mod/locale"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestMonthByName(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		name     string
		expMonth time.Month
	}{
		{
			ID:     testhelper.MkID("no name"),
			ExpErr: testhelper.MkExpErr(`unknown month: ""`),
		},
		{
			ID:     testhelper.MkID("nonesuch"),
			ExpErr: testhelper.MkExpErr(`unknown month: "nonesuch"`),
			name:   "nonesuch",
		},
		{
			ID:       testhelper.MkID("January"),
			name:     "January",
			expMonth: time.January,
		},
		{
			ID:       testhelper.MkID("Jan"),
			name:     "Jan",
			expMonth: time.January,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			month, err := locale.MonthByName(tc.name)
			if !testhelper.CheckExpErr(t, err, tc) {
				return
			}

			if err != nil {
				return
			}

			testhelper.DiffInt(t, tc.IDStr(), "month", month, tc.expMonth)
		})
	}
}

func TestMonthByCountry(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		name        string
		countryName string
		expMonth    time.Month
	}{
		{
			ID:     testhelper.MkID("no name, no country"),
			ExpErr: testhelper.MkExpErr(`unknown country code: ""`),
		},
		{
			ID:          testhelper.MkID("no name, nonesuch"),
			ExpErr:      testhelper.MkExpErr(`unknown country code: "nonesuch"`),
			countryName: "nonesuch",
		},
		{
			ID:          testhelper.MkID("no name, English"),
			ExpErr:      testhelper.MkExpErr(`unknown month: ""`),
			countryName: "English",
		},
		{
			ID:          testhelper.MkID("nonesuch, English"),
			ExpErr:      testhelper.MkExpErr(`unknown month: "nonesuch"`),
			name:        "nonesuch",
			countryName: "English",
		},
		{
			ID:          testhelper.MkID("January, English"),
			name:        "January",
			countryName: "English",
			expMonth:    time.January,
		},
		{
			ID:          testhelper.MkID("Jan, English"),
			name:        "Jan",
			countryName: "English",
			expMonth:    time.January,
		},
		{
			ID:          testhelper.MkID("Feb, French"),
			ExpErr:      testhelper.MkExpErr(`unknown month: "Feb"`),
			name:        "Feb",
			countryName: "French",
		},
		{
			ID:          testhelper.MkID("Janvier, FR"),
			name:        "Janvier",
			countryName: "FR",
			expMonth:    time.January,
		},
		{
			ID:          testhelper.MkID("Janvier, Français"),
			name:        "Janvier",
			countryName: "Français",
			expMonth:    time.January,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			month, err := locale.MonthByNameByCountry(
				tc.name, tc.countryName)
			if !testhelper.CheckExpErr(t, err, tc) {
				return
			}

			if err != nil {
				return
			}

			testhelper.DiffInt(t, tc.IDStr(), "month", month, tc.expMonth)
		})
	}
}
