package iteration

func Repeat(text string, itr int) string {
	var retVal string
	for i := 0; i < itr; i++ {
		retVal += text
	}
	return retVal
}
