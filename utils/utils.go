package utils

//Revers get string and return revers of it
func ReverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func RemoveRightZeros(s string) string {
	for len(s) != 0 && s[0] == '0' {
		s = s[1:]
	}
	if len(s) == 0 {
		s = "0"
	}
	return s
}
