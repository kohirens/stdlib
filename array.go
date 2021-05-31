package stdlib

// Prepend Put an element at the beginning of an array.
func Prepend(item interface{}, a ...interface{}) (s []interface{}) {
	// switch v := item.(type) {
	// case int:
	// 	s = append([]interface{}{item}, a...)
	// default:
	// }

	s = append([]interface{}{item}, a...)

	return
}
