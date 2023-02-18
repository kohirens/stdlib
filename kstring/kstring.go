package kstring

// InArray Returns true if the string is found in the array of strings.
func InArray(currFile string, files []string) bool {
	for _, aFile := range files {
		if aFile == currFile {
			return true
		}
	}

	return false
}
