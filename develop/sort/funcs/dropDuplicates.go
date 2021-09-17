package funcs

func dropDuplicates(words []string) []string {
	unique := []string{}

	for _, word := range words {
		// If we alredy have this word, skip.
		if contains(unique, word) {
			continue
		}
		unique = append(unique, word)
	}

	return unique
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
