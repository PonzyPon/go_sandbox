package main

func main() {
	incrementer := generateIncrementer(10)
	println(incrementer()) // 11
	println(incrementer()) // 12
	println(incrementer()) // 13
	println(incrementer()) // 14
}

func generateIncrementer(initial int) func() int {
	value := initial
	return func() int {
		value++
		return value
	}
}
