Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function
and prints success or failure.

package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
    myMap := make(map[string]int)
    w := strings.Fields(s)   // Strings.Fields split the string using space delimeter.
	fmt.Println(w)
	for  _, v := range w {
	    myMap[v]++
	}
	 return (myMap)
}

func main() {
	wc.Test(WordCount)
}
