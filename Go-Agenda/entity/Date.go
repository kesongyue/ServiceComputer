package entity

import "strconv"

type Date struct {
	Year, Month, Day, Hour, Minute int
}

func StringToDate(dateString string) Date {
	trash := Date{0, 0, 0, 0, 0}
	if len(dateString) != 16 {
		return trash
	}
	if dateString[4] != '-' || dateString[7] != '-' || dateString[10] != '/' || dateString[13] != ':' {
		return trash
	}
	yearStr := dateString[0:4]
	monthStr := dateString[5:7]
	dayStr := dateString[8:10]
	hourStr := dateString[11:13]
	minStr := dateString[14:16]

	year, error1 := strconv.Atoi(yearStr)
	month, error2 := strconv.Atoi(monthStr)
	day, error3 := strconv.Atoi(dayStr)
	hour, error4 := strconv.Atoi(hourStr)
	min, error5 := strconv.Atoi(minStr)

	if error1 != nil || error2 != nil || error3 != nil || error4 != nil || error5 != nil {
		return trash
	} else {
		newDate := Date{year, month, day, hour, min}
		return newDate
	}
}

func IsDateValid(date Date) bool {
	var mon1 = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var mon2 = [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if date.Year < 1000 || date.Year > 9999 {
		return false
	}
	if date.Year%400 == 0 || (date.Year%100 != 0 && date.Year%4 == 0) {
		if date.Month > 0 && date.Month < 13 && date.Day <= mon2[date.Month-1] &&
			date.Day > 0 && date.Hour >= 0 && date.Hour < 24 && date.Minute <= 59 && date.Minute >= 0 {
			return true
		} else {
			return false
		}
	} else {
		if date.Month > 0 && date.Month < 13 && date.Day <= mon1[date.Month-1] &&
			date.Day > 0 && date.Hour >= 0 && date.Hour < 24 && date.Minute <= 59 && date.Minute >= 0 {
			return true
		} else {
			return false
		}
	}
}

func IsEndBigthanStart(start, end Date) bool {
	if end.Year > start.Year {
		return true
	}
	if end.Year == start.Year && end.Month > start.Month {
		return true
	}
	if end.Year == start.Year && end.Month == start.Month && end.Day > start.Day {
		return true
	}
	if end.Year == start.Year && end.Month == start.Month && end.Day == start.Day && end.Hour > start.Hour {
		return true
	}
	if end.Year == start.Year && end.Month == start.Month && end.Day == start.Day && end.Hour == start.Hour && end.Minute > start.Minute {
		return true
	}
	return false
}

func (date Date) DateToString() string {
	year := strconv.Itoa(date.Year)
	month := strconv.Itoa(date.Month)
	day := strconv.Itoa(date.Day)
	hour := strconv.Itoa(date.Hour)
	min := strconv.Itoa(date.Minute)

	dateString := year + "-" + month + "-" + day + "/" + hour + ":" + min
	return dateString
}
