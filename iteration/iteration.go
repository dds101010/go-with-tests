package iteration

import "strings"

func Repeat(text string, itr int) string {
	var retVal string
	for i := 0; i < itr; i++ {
		retVal += text
	}
	return retVal
}

// Optimized version of Repeat function that uses lib implementation
func RepeatOpt(text string, itr int) string {
	return strings.Repeat(text, itr)
}
