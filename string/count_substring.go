package main

import (
	"fmt"
	"strings"
)

func main(){

	var aString string = "this is a string.    "
	var s string = "s"

	fmt.Println(strings.Index(aString, s))
	// Index Substring
	fmt.Println(strings.Count(aString, s))
	// Count Substring
	fmt.Println(strings.Replace(aString, s, "S", -1))
	// Replace
	fmt.Println(aString)
	// Println
	fmt.Println(strings.ToLower(aString))
	// string ToLower
	fmt.Println(strings.ToUpper(aString))
	// string ToUpper

	fmt.Println(strings.TrimSpace(aString))
	// TrimSpace
	fmt.Println(strings.Trim(aString, "."))
	// Trim

	str := "The quick brown fox jumps over the lazy dog"
    sl := strings.Fields(str)
    fmt.Printf("Splitted in slice: %v\n", sl)
    for a, val := range sl {
		fmt.Print(a)
        fmt.Printf("%s - ", val)
	}
	fmt.Println()
	fmt.Println(strings.Join(sl, "|"))
	// Join func

	
}