package utils

import (
	"fmt"
	"regexp"
)
func SimpleMatch(){
	str := "antony musumba gopher"

    // returns true if the pattern is present 
    // in the string, else returns false 
    // err is nil if the regexp is valid 
	firstMatch , err := regexp.MatchString("gopher",str)
	fmt.Println("Match :",firstMatch,"Error :", err)
}