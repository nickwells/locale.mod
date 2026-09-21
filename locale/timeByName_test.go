package locale_test

import (
	"testing"

	"github.com/nickwells/locale.mod/locale"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestTimeByName(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		testhelper.ExpErr
		name    string
		expTime locale.Time
	}{
		{
			ID:     testhelper.MkID("bad - unknown name"),
			ExpErr: testhelper.MkExpErr(`unknown language: "nonesuch"`),
			name:   "nonesuch",
		},
		{
			ID:      testhelper.MkID("good - TimeEnglish"),
			name:    "en",
			expTime: locale.TimeEnglish,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			actTime, err := locale.TimeByName(tc.name)
			testhelper.CheckExpErr(t, err, tc)
			testhelper.DiffValsReport(t,
				tc.IDStr(), "Time",
				actTime, tc.expTime)
		})
	}
}

func TestTimeByEnv(t *testing.T) {
	const (
		languageEV = "LANGUAGE"
		langEV     = "LANG"
	)

	testCases := []struct {
		testhelper.ID
		language string
		lang     string
		expTime  locale.Time
	}{
		{
			ID:       testhelper.MkID("unknown LANGUAGE - TimeEnglish"),
			language: "nonesuch",
			expTime:  locale.TimeEnglish,
		},
		{
			ID:      testhelper.MkID("unknown LANG - TimeEnglish"),
			lang:    "nonesuch",
			expTime: locale.TimeEnglish,
		},
		{
			ID:      testhelper.MkID("no LANGUAGE or LANG - TimeEnglish"),
			expTime: locale.TimeEnglish,
		},
		{
			ID:       testhelper.MkID("good - LANGUAGE - TimeFrench"),
			language: "nonesuch:fr",
			expTime:  locale.TimeFrench,
		},
		{
			ID:      testhelper.MkID("good - LANG - TimeGerman"),
			lang:    "de_DE.UTF-8",
			expTime: locale.TimeGerman,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			var ec testhelper.EnvCache

			ec.Clearenv()
			defer ec.ResetEnv()

			if tc.language != "" {
				err := ec.Setenv(testhelper.EnvEntry{
					Key:   languageEV,
					Value: tc.language,
				})
				if err != nil {
					t.Logf("Env : %q=%q", languageEV, tc.language)
					t.Fatalf("Couldn't populate the environment")
				}
			}

			if tc.lang != "" {
				err := ec.Setenv(testhelper.EnvEntry{
					Key:   langEV,
					Value: tc.lang,
				})
				if err != nil {
					t.Logf("Env : %q=%q", langEV, tc.lang)
					t.Fatalf("Couldn't populate the environment")
				}
			}

			actTime := locale.TimeByEnv()
			testhelper.DiffValsReport(t,
				tc.IDStr(), "Time",
				actTime, tc.expTime)
		})
	}
}
