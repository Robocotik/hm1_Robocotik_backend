package main

import (
	"fmt"
	"strings"
)

func main(){
	var str string
	fmt.Scanf("%s", &str)
	str = strings.ReplaceAll(str,"1","one")
	fmt.Println(str)
}