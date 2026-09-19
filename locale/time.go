package locale

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/nickwells/tempus.mod/v2/tempus"
)

// Time holds conversion tables for mapping language-specific names to the
// tokens used in the standard time package and vice versa. Various languages
// are pre-populated and are available either through named variables or the
// TimeByName func.
type Time struct {
	toWeekday   map[string]time.Weekday
	toMonth     map[string]time.Month
	fromWeekday map[time.Weekday]string
	fromMonth   map[time.Month]string
}

// WeekdayNames returns the names of recognised weekdays. They are sorted by
// day of week and name.
func (l Time) WeekdayNames() []string {
	weekdays := maps.Keys(l.toWeekday)

	return slices.SortedFunc(weekdays, func(a, b string) int {
		rval := cmp.Compare(l.toWeekday[a], l.toWeekday[b])
		if rval != 0 {
			return rval
		}

		return cmp.Compare(a, b)
	})
}

// MonthNames returns the names of recognised months. They are sorted by
// month and name.
func (l Time) MonthNames() []string {
	months := maps.Keys(l.toMonth)

	return slices.SortedFunc(months, func(a, b string) int {
		rval := cmp.Compare(l.toMonth[a], l.toMonth[b])
		if rval != 0 {
			return rval
		}

		return cmp.Compare(a, b)
	})
}

// ToWeekday returns the day of week corresponding to the given name. If the
// name cannot be found a non-nil error is returned.
func (l Time) ToWeekday(name string) (time.Weekday, error) {
	w, ok := l.toWeekday[name]
	if !ok {
		return w, fmt.Errorf("unknown day of week: %q", name)
	}

	return w, nil
}

// ToMonth returns the month corresponding to the given name. If the name
// cannot be found it will attempt to parse the name as a number with,
// possibly, a leading zero. Numbers 1-12 corresponding to months January to
// December. If neither interpretation yields a valid month a non-nil error
// is returned.
func (l Time) ToMonth(name string) (time.Month, error) {
	m, ok := l.toMonth[name]
	if !ok {
		asNumber := strings.Trim(name, "0")
		if monthNum, err := strconv.ParseInt(asNumber, 10, 64); err == nil {
			m = time.Month(monthNum)
			if m >= time.January || m <= time.December {
				return m, nil
			}
		}

		return m, fmt.Errorf("unknown month: %q", name)
	}

	return m, nil
}

// FromWeekday returns the name corresponding to the given day of week. if the
// weekday cannot be found a non-nil error is returned.
func (l Time) FromWeekday(w time.Weekday) (string, error) {
	name, ok := l.fromWeekday[w]
	if !ok {
		return name, fmt.Errorf("unknown day of week: %q", w)
	}

	return name, nil
}

// FromMonth returns the name corresponding to the given month. if the
// month cannot be found a non-nil error is returned.
func (l Time) FromMonth(m time.Month) (string, error) {
	name, ok := l.fromMonth[m]
	if !ok {
		return name, fmt.Errorf("unknown month: %q", m)
	}

	return name, nil
}

// keyType is the interface which findMissingKeys expects a map key to
// satisfy.
type keyType interface {
	fmt.Stringer
	comparable
}

// findMissingKeys checks that every value in the slice 'exp' yields a map
// entry in 'm' with value true. The string value of the key for missing
// entries is added to the return message. If there are no missing entries
// the return value is a string saying that all expected 'name's are present.
func findMissingKeys[K keyType, V any](exp []K, m map[K]V, name string) string {
	var missing []string

	for _, k := range exp {
		if _, ok := m[k]; !ok {
			missing = append(missing, k.String())
		}
	}

	if len(missing) == 0 {
		return "all expected " + name + " are present"
	}

	return "missing: " + strings.Join(missing, ", ")
}

// checkTimeToWeekdayMap returns a non-nil error if the toWeekday map does
// not have some mapping to each day of the week or if any of the mappings is
// to an invalid weekday.
func checkTimeToWeekdayMap(toWeekday map[string]time.Weekday) error {
	weekdayMap := map[time.Weekday]bool{}

	for name, wd := range toWeekday {
		if wd < time.Sunday || wd > time.Saturday {
			return fmt.Errorf(
				"toWeekday maps %q to %s (which is not a valid weekday)",
				name, wd)
		}

		weekdayMap[wd] = true
	}

	if len(weekdayMap) != tempus.DaysPerWeek {
		return fmt.Errorf("toWeekday only has mappings to %d days; %s",
			len(weekdayMap),
			findMissingKeys(tempus.AllWeekdays, weekdayMap, "weekdays"))
	}

	return nil
}

// checkTimeToMonthMap returns a non-nil error if the toMonth map does not
// have some mapping to each month or if any of the mappings is to an invalid
// month.
func checkTimeToMonthMap(toMonth map[string]time.Month) error {
	monthMap := map[time.Month]bool{}

	for name, m := range toMonth {
		if m < time.January || m > time.December {
			return fmt.Errorf(
				"toMonth maps %q to %s (which is not a valid month)",
				name, m)
		}

		monthMap[m] = true
	}

	if len(monthMap) != tempus.MonthsPerYear {
		return fmt.Errorf("toMonth only has mappings to %d months; %s",
			len(monthMap),
			findMissingKeys(tempus.AllMonths, monthMap, "months"))
	}

	return nil
}

// checkTimeFromWeekdayMap returns a non-nil error if the fromWeekday map
// does not have a mapping from each day of the week, if any of the weekdays
// is invalid or if any of the mappings are to duplicate strings.
func checkTimeFromWeekdayMap(fromWeekday map[time.Weekday]string) error {
	dups := map[string]time.Weekday{}

	for wd, name := range fromWeekday {
		if wd < time.Sunday || wd > time.Saturday {
			return fmt.Errorf(
				"fromWeekday maps %s (which is not a valid weekday) to %q",
				wd, name)
		}

		if existingWD, ok := dups[name]; ok {
			if wd < existingWD { // to make tests reproducible
				wd, existingWD = existingWD, wd
			}

			return fmt.Errorf(
				"duplicate detected; fromWeekday maps both %s and %s to %q",
				existingWD, wd, name)
		}

		dups[name] = wd
	}

	if len(fromWeekday) != tempus.DaysPerWeek {
		return fmt.Errorf("fromWeekday only has mappings to %d weekdays; %s",
			len(fromWeekday),
			findMissingKeys(tempus.AllWeekdays, fromWeekday, "weekdays"))
	}

	return nil
}

// checkTimeFromMonthMap returns a non-nil error if the fromMonth map does
// not have a mapping from each month, if any of the months is invalid or if
// any of the mappings are to duplicate strings.
func checkTimeFromMonthMap(fromMonth map[time.Month]string) error {
	dups := map[string]time.Month{}

	for m, name := range fromMonth {
		if m < time.January || m > time.December {
			return fmt.Errorf(
				"fromMonth maps %s (which is not a valid month) to %q",
				m, name)
		}

		if existingMonth, ok := dups[name]; ok {
			if m < existingMonth { // to make tests reproducible
				m, existingMonth = existingMonth, m
			}

			return fmt.Errorf(
				"duplicate detected; fromMonth maps both %s and %s to %q",
				existingMonth, m, name)
		}

		dups[name] = m
	}

	if len(fromMonth) != tempus.MonthsPerYear {
		return fmt.Errorf("fromMonth only has mappings to %d months; %s",
			len(fromMonth),
			findMissingKeys(tempus.AllMonths, fromMonth, "months"))
	}

	return nil
}

// checkTimeRoundtrip returns a non-nil error if any entry in the
// from map generates a name which does not take you back to the same
// Month or Weekday when mapping through the to map.
func checkTimeRoundtrip[T time.Weekday | time.Month](
	from map[T]string, to map[string]T,
	all []T,
) error {
	for _, start := range all {
		name, ok := from[start]
		if !ok {
			return fmt.Errorf("bad %T maps: from[time.%s] is not found",
				start, start)
		}

		end, ok := to[name]
		if !ok {
			return fmt.Errorf(
				"bad %T maps: from[time.%s] gives %q but to[%q] is not found",
				start,
				start, name, name)
		}

		if start != end {
			return fmt.Errorf(
				"bad %T maps: from[time.%s] gives %q"+
					" but to[%q] gives time.%s not time.%s",
				start,
				start, name,
				name, end, start)
		}
	}

	return nil
}

// MakeTime constructs a Time (a mapping between names and time values for
// Month and Weekday). It also checks the supplied parameters for validity so
// that the resultant Time has coherent and comprehensive maps. If the
// parameters are all valid then it will return the constructed Time and a
// nil error. If any of the parameters is invalid it will return an empty
// Time and a non-nil error.
func MakeTime(
	toWeekday map[string]time.Weekday,
	toMonth map[string]time.Month,
	fromWeekday map[time.Weekday]string,
	fromMonth map[time.Month]string,
) (Time, error) {
	l := Time{}

	if err := checkTimeToWeekdayMap(toWeekday); err != nil {
		return l, err
	}

	if err := checkTimeToMonthMap(toMonth); err != nil {
		return l, err
	}

	if err := checkTimeFromWeekdayMap(fromWeekday); err != nil {
		return l, err
	}

	if err := checkTimeFromMonthMap(fromMonth); err != nil {
		return l, err
	}

	if err := checkTimeRoundtrip(fromMonth, toMonth, tempus.AllMonths); err != nil {
		return l, err
	}

	if err := checkTimeRoundtrip(fromWeekday, toWeekday, tempus.AllWeekdays); err != nil {
		return l, err
	}

	l.toWeekday = toWeekday
	l.toMonth = toMonth
	l.fromWeekday = fromWeekday
	l.fromMonth = fromMonth

	return l, nil
}

// MakeTimeOrPanic constructs a Time just like MakeTime but if an error is
// detected it will panic.
func MakeTimeOrPanic(
	toWeekday map[string]time.Weekday,
	toMonth map[string]time.Month,
	fromWeekday map[time.Weekday]string,
	fromMonth map[time.Month]string,
) Time {
	l, err := MakeTime(toWeekday, toMonth, fromWeekday, fromMonth)
	if err != nil {
		panic(err)
	}

	return l
}
