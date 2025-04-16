package main

import (
	"fmt"
	"sort"
)

func main() {
	intToRomanMap := map[int]string{
		1000: "M", 900: "CM", 500: "D", 400: "CD",
		100: "C", 90: "XC", 50: "L", 40: "XL",
		10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I",
	}

	numToFind := 177
	romanStr := IntToRoman(numToFind, intToRomanMap)

	fmt.Printf("Roman number is %s\n", romanStr)
}

func IntToRoman(num int, intRomMap map[int]string) string {
	romanNum := ""

	var keys []int
	for key := range intRomMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, key := range keys {
		for num >= key {
			romanNum += intRomMap[key]
			num -= key
		}
	}
	return romanNum
}
