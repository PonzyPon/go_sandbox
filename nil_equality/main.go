package main

func main() {
	var nilStr1 *string
	var nilStr2 *string
	println(nilStr1 == nil)     // true
	println(nilStr1 == nilStr2) // true

	var foo1 *foo
	var foo2 *foo
	println(foo1 == nil)  // true
	println(foo1 == foo2) // true

	var i1 interface{}
	var i2 interface{}
	println(i1 == nil) // true
	println(i1 == i2)  // true
	// println(nilStr1 == foo1) // compilation error
	println(i1 == nilStr1) // false。どちらも == nilはtrueだが、比較するとfalse。
}

type foo struct{}
