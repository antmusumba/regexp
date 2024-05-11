package utils

import (
	"fmt"
	"regexp"
)
func RegObject(){
	// a regexp object that can be used later
	re,_:= regexp.Compile("software")
	// 
	str := "all software developers must co-exist"
	match:= re.FindStringIndex(str)
	fmt.Println(match)
}