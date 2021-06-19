package team

import "strconv"

func ToOrdinalNumber(number int) string {
	num := strconv.Itoa(number)
	lastChar := num[len(num)-1:]
	switch lastChar {
	case "1":
		return num + "st"
	case "2":
		return num + "nd"
	case "3":
		return num + "rd"
	default:
		return num + "th"
	}
}
