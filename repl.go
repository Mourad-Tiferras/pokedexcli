package main


func cleanInput(text string) []string {
	var result []string
	word := ""

	for _, ch := range text {
		if ch == ' ' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}