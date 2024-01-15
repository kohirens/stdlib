package str

import (
	"bytes"
	"fmt"
	"regexp"
)

// InArray Returns true if the string is found in the array of strings.
func InArray(currFile string, files []string) bool {
	for _, aFile := range files {
		if aFile == currFile {
			return true
		}
	}

	return false
}

// ToCamelCase convert a dash/underscore separated string to camel case. Works great on pretty/vanity URLs.
//
// camel case - a typographical convention in which an initial capital is
// used for the first letter of a word forming the second element of a
// closed compound, e.g. iPhone, eBay, PayPal.
//
// Pascal case -- or PascalCase - is a programming naming convention where
// the first letter of each compound word in a variable is capitalized.
func ToCamelCase(subject, separate string, pascal bool) (string, error) {
	if len(subject) < 1 {
		return subject, fmt.Errorf("there is nothing to do, the string is empty %q", subject)
	}

	// Look for the separator, when found return it and the following char if any.
	re, err1 := regexp.Compile("(?:^|" + separate + ")(.?)")

	if err1 != nil {
		return "", err1
	}

	sb := []byte(separate)
	rVal := re.ReplaceAllFunc([]byte(subject), func(match []byte) []byte {
		if len(match) == 2 { // match the separator plus 1 letter.
			b := bytes.ToUpper(match)[1]
			return []byte{b}
		}

		if sb[0] != match[0] { // match the first letter
			if pascal {
				return bytes.ToUpper(match)
			} else {
				return bytes.ToLower(match)
			}
		}

		return []byte{}
	})

	return string(rVal), nil
}

// StringMap use for usage messages, templates, and template vars.
type StringMap = map[string]string
