package main

func main() {
	value := getValue()
	println(&value) //0xc000046720
	pointer := getPointer()
	println(pointer) //0xc000046730
}

func getValue() string {
	str := "string"
	println(&str) // 0xc000046740
	return str
}

func getPointer() *string {
	str := "string"
	println(&str) //0xc000046730
	return &str
}
