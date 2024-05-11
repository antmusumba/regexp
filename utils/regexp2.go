package utils

import (
	"fmt"
	"regexp"
)
func SimpleMatchFalse(){
	str := "antony musumba gopher"
	secondMatch, err := regexp.MatchString("ful-stack",str)
	fmt.Println("Match :", secondMatch, "Error :", err)
}