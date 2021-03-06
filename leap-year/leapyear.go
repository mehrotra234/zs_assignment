package leapyear

// CheckLeapYear : Checks if the year is leap year or not
func CheckLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
