package main

func deferTest(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}
func double(x int) int {
	return x + x
}
