package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	var endpoint = "https://www.examle.com"

	url, _ := url.Parse(endpoint)

	url.Path = path.Join(url.Path, "test")

	query := url.Query()
	query.Set("firstParam", "1")
	query.Set("SecondParam", "two")
	url.RawQuery = query.Encode()

	// https://www.examle.com/test?SecondParam=two&firstParam=1
	fmt.Println(url.String())
}
