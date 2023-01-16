package kronos

var allCountries *map[string][]string

func init() {
	x := map[string][]string{
		// English month and day
		"EnglishMonth": {
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
		"EnglishDay": {
			"Sunday",
			"Monday",
			"Tuesday",
			"Wednesday",
			"Thursday",
			"Friday",
			"Saturday",
		},
		// Turkish month and day
		"TurkishMonth": {
			"Ocak",
			"Şubat",
			"Mart",
			"Nisan",
			"Mayis",
			"Haziran",
			"Temmuz",
			"Ağustos",
			"Eylül",
			"Ekim",
			"Kasim",
			"Aralik",
		},
		"TurkishDay": {
			"Pazar",
			"Pazartesi",
			"Salı",
			"Çarşamba",
			"Perşembe",
			"Cuma",
			"Cumartesi",
		},
	}
	allCountries = &x

}
