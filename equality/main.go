package main

type myString struct{ Value string }
type myInt struct{ Value int64 }
type myBool struct{ Value bool }

func main() {
	str1 := myString{Value: "value"}
	str2 := myString{Value: "value"}
	println(str1 == str2)   // true
	println(&str1 == &str2) // false

	myMap := map[myString]string{}
	myMap[str1] = "str1_value"
	println(myMap[str1]) // str1_value
	println(myMap[str2]) // str1_value
}
