package locale

import "time"

// TimeChinese provides the Time conversion tables for the Chinese
// language. Note that both simplified and traditional Chinese (and pinyin,
// both accented and unaccented) are used in the maps to days of weeks and
// months. Simplified Chinese is used as the result in the maps from weekdays
// and months. This is in deference to the greater number of people using
// simplified Chinese.
var TimeChinese = MakeTimeOrPanic(
	map[string]time.Weekday{
		"日":          time.Sunday,
		"rì":         time.Sunday,
		"ri":         time.Sunday,
		"星期日":        time.Sunday,
		"xīngqīrì":   time.Sunday,
		"xingqirì":   time.Sunday,
		"星期天":        time.Sunday,
		"xīngqītiān": time.Sunday,
		"xingqitian": time.Sunday,
		"週日":         time.Sunday, // traditional Chinese
		"周日":         time.Sunday, // simplified Chinese
		"zhōurì":     time.Sunday,
		"zhouri":     time.Sunday,
		"週天":         time.Sunday, // traditional Chinese
		"周天":         time.Sunday, // simplified Chinese
		"zhōutiān":   time.Sunday,
		"zhoutian":   time.Sunday,
		"禮拜天":        time.Sunday, // traditional Chinese
		"礼拜天":        time.Sunday, // simplified Chinese
		"lǐbàitiān":  time.Sunday,
		"libaitian":  time.Sunday,
		"禮拜日":        time.Sunday, // traditional Chinese
		"礼拜日":        time.Sunday, // simplified Chinese
		"lǐbàirì":    time.Sunday,
		"libairi":    time.Sunday,
		"一":          time.Monday,
		"yī":         time.Monday,
		"yi":         time.Monday,
		"星期一":        time.Monday,
		"xīngqīyī":   time.Monday,
		"xingqiyi":   time.Monday,
		"週一":         time.Monday, // traditional Chinese
		"周一":         time.Monday, // simplified Chinese
		"zhōuyī":     time.Monday,
		"zhouyi":     time.Monday,
		"禮拜一":        time.Monday, // traditional Chinese
		"礼拜一":        time.Monday, // simplified Chinese
		"lǐbàiyī":    time.Monday,
		"libaiyi":    time.Monday,
		"二":          time.Tuesday,
		"èr":         time.Tuesday,
		"er":         time.Tuesday,
		"星期二":        time.Tuesday,
		"xīngqīèr":   time.Tuesday,
		"xingqier":   time.Tuesday,
		"週二":         time.Tuesday, // traditional Chinese
		"周二":         time.Tuesday, // simplified Chinese
		"zhōuèr":     time.Tuesday,
		"zhouer":     time.Tuesday,
		"禮拜二":        time.Tuesday, // traditional Chinese
		"礼拜二":        time.Tuesday, // simplified Chinese
		"lǐbàièr":    time.Tuesday,
		"libaier":    time.Tuesday,
		"三":          time.Wednesday,
		"sān":        time.Wednesday,
		"san":        time.Wednesday,
		"星期三":        time.Wednesday,
		"xīngqīsān":  time.Wednesday,
		"xingqisan":  time.Wednesday,
		"週三":         time.Wednesday, // traditional Chinese
		"周三":         time.Wednesday, // simplified Chinese
		"zhōusān":    time.Wednesday,
		"zhousan":    time.Wednesday,
		"禮拜三":        time.Wednesday, // traditional Chinese
		"礼拜三":        time.Wednesday, // simplified Chinese
		"lǐbàisān":   time.Wednesday,
		"libaisan":   time.Wednesday,
		"四":          time.Thursday,
		"sì":         time.Thursday,
		"si":         time.Thursday,
		"星期四":        time.Thursday,
		"xīngqīsì":   time.Thursday,
		"xingqisi":   time.Thursday,
		"週四":         time.Thursday, // traditional Chinese
		"周四":         time.Thursday, // simplified Chinese
		"zhōusì":     time.Thursday,
		"zhousi":     time.Thursday,
		"禮拜四":        time.Thursday, // traditional Chinese
		"礼拜四":        time.Thursday, // simplified Chinese
		"lǐbàisì":    time.Thursday,
		"libaisi":    time.Thursday,
		"五":          time.Friday,
		"wǔ":         time.Friday,
		"wu":         time.Friday,
		"星期五":        time.Friday,
		"xīngqīwǔ":   time.Friday,
		"xingqiwu":   time.Friday,
		"週五":         time.Friday, // traditional Chinese
		"周五":         time.Friday, // simplified Chinese
		"zhōuwǔ":     time.Friday,
		"zhouwu":     time.Friday,
		"禮拜五":        time.Friday, // traditional Chinese
		"礼拜五":        time.Friday, // simplified Chinese
		"lǐbàiwǔ":    time.Friday,
		"libaiwu":    time.Friday,
		"六":          time.Saturday,
		"liù":        time.Saturday,
		"liu":        time.Saturday,
		"星期六":        time.Saturday,
		"xīngqīliù":  time.Saturday,
		"xingqiliu":  time.Saturday,
		"週六":         time.Saturday, // traditional Chinese
		"周六":         time.Saturday, // simplified Chinese
		"zhōuliù":    time.Saturday,
		"zhouliu":    time.Saturday,
		"禮拜六":        time.Saturday, // traditional Chinese
		"礼拜六":        time.Saturday, // simplified Chinese
		"lǐbàiliù":   time.Saturday,
		"libailiu":   time.Saturday,
	},
	map[string]time.Month{
		"一月":         time.January,
		"yī yuè":     time.January,
		"yi yue":     time.January,
		"二月":         time.February,
		"èr yuè":     time.February,
		"er yue":     time.February,
		"三月":         time.March,
		"sān yuè":    time.March,
		"san yue":    time.March,
		"四月":         time.April,
		"sì yuè":     time.April,
		"si yue":     time.April,
		"五月":         time.May,
		"wǔ yuè":     time.May,
		"wu yue":     time.May,
		"六月":         time.June,
		"liù yuè":    time.June,
		"liu yue":    time.June,
		"七月":         time.July,
		"qī yuè":     time.July,
		"qi yue":     time.July,
		"八月":         time.August,
		"bā yuè":     time.August,
		"ba yue":     time.August,
		"九月":         time.September,
		"jiǔ yuè":    time.September,
		"jiu yue":    time.September,
		"十月":         time.October,
		"shí yuè":    time.October,
		"shi yue":    time.October,
		"十一月":        time.November,
		"shí yī yuè": time.November,
		"shi yi yue": time.November,
		"十二月":        time.December,
		"shí èr yuè": time.December,
		"shi er yue": time.December,
	},
	map[time.Weekday]string{
		time.Monday:    "星期一",
		time.Tuesday:   "星期二",
		time.Wednesday: "星期三",
		time.Thursday:  "星期四",
		time.Friday:    "星期五",
		time.Saturday:  "星期六",
		time.Sunday:    "星期日",
	},
	map[time.Month]string{
		time.January:   "一月",
		time.February:  "二月",
		time.March:     "三月",
		time.April:     "四月",
		time.May:       "五月",
		time.June:      "六月",
		time.July:      "七月",
		time.August:    "八月",
		time.September: "九月",
		time.October:   "十月",
		time.November:  "十一月",
		time.December:  "十二月",
	},
)
