package a

func nums(s string) string {
	if string(s[len(s)-1]) == "s" {
		return s + "es"
	} else {
		return s + "s"
	}
}
