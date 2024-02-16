package slice

// PrependInt Put an element at the beginning of a slice.
// Deprecated see Prepend
func PrependInt(item int, ary []int) (s []int) {
	return append([]int{item}, ary...)
}

// PrependByte Put an element at the beginning of a byte slice.
// Deprecated see Prepend
func PrependByte(item byte, ary []byte) (s []byte) {
	str := string(item)
	return append([]byte(str), ary...)
}

// Prepend Put an element at the beginning of a byte slice.
//
//	NOTE: This is inefficient as it touches every element of the array, and is
//	only meant for small arrays.
func Prepend[V comparable](ary []V, item V) []V {
	newAry := []V{item}
	return append(newAry, ary...)
}
