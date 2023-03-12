package main

import "fmt"

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func main() {
	// b := [...]string{"Penn", "Teller"}
	// fmt.Println(b)

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b)

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // a slice referencing the storage of x

	fmt.Println(s)

	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	fmt.Println(p)
}
