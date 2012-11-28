package main

import (
	PrefixLookup "./PrefixLookup"
	"fmt"
)

func main() {
	//Load the prefix database
	PrefixLookup.LoadPrefixes()

	//Get the country the callsign is from
	if PrefixLookup.CountryForCallsign("K2GXT") == "United States" {
		fmt.Println("Test 1 Passed")
	} else {
		fmt.Println("Test 1 FAILED")
	}
	if PrefixLookup.CountryForCallsign("SSS3FF") == "Sudan" {
		fmt.Println("Test 2 Passed")
	} else {
		fmt.Println("Test 2 FAILED")
	}
	if PrefixLookup.CountryForCallsign("hb3Yd") == "Liechtenstein" {
		fmt.Println("Test 3 Passed")
	} else {
		fmt.Println("Test 3 FAILED")
	}
	if PrefixLookup.CountryForCallsign("hb0xx") == "Liechtenstein" {
		fmt.Println("Test 4 Passed")
	} else {
		fmt.Println("Test 4 FAILED")
	}
	if PrefixLookup.CountryForCallsign("hb2g") == "Switzerland" {
		fmt.Println("Test 5 Passed")
	} else {
		fmt.Println("Test 5 FAILED")
	}

}
