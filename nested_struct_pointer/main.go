package main

func main() {
	bar := Bar{baz: &Baz{}}
	foo := Foo{bar: &bar}
	println(&foo)         //0xc000046738
	println(foo.bar)      //0xc000046740
	println(&foo.bar)     //0xc000046738
	println(&(foo.bar))   //0xc000046738
	println(&foo.bar.baz) //0xc000046740
}

type Foo struct {
	bar *Bar
}

type Bar struct {
	baz *Baz
}

type Baz struct {
}
