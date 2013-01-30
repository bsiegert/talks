package main

func main() {
	// BEGIN OMIT
	a := []int{1, 2, 3, 4} // len(a) = 4, cap(a) = 4
	b := a[2:4]            // b[0] = 3, b[1] = 4

	b = append(b, 5) // b[2] = 5
	_ = b[3]         // will panic
	// END OMIT
}
