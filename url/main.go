package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	endpoint := "https://www.examle.com?param=value"

	generatedURL, _ := url.Parse(endpoint)

	generatedURL.Path = path.Join(generatedURL.Path, "test")

	query := generatedURL.Query()
	query.Set("firstParam", "1")
	query.Set("SecondParam", "two")
	generatedURL.RawQuery = query.Encode()

	// https://www.examle.com/test?SecondParam=two&firstParam=1&param=value
	fmt.Println(generatedURL.String())
}
