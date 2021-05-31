package slice

// PrependInt Put an element at the beginning of a slice.
func PrependInt(item int, ary []int) (s []int) {
	return append([]int{item}, ary...)
}

// PrependByte Put an element at the beginning of a byte slice.
func PrependByte(item byte, ary []byte) (s []byte) {
	str := string(item)
	return append([]byte(str), ary...)
}
