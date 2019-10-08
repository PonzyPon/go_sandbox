package main

import (
	"encoding/json"
	"fmt"
)

// Key is awesome
type Key struct {
	ID      string `json:"id"`
	Version string `json:"version"`
}

func main() {
	key1 := Key{
		"id1",
		"version1",
	}
	keyJSON, _ := json.Marshal(key1)
	fmt.Println(string(keyJSON))

	key2 := Key{
		"id2",
		"version2",
	}
	key3 := Key{
		"id3",
		"version3",
	}

	keys := []Key{key1, key2, key3}
	keySliceJSON, _ := json.Marshal(keys)
	fmt.Println(string(keySliceJSON))
}
