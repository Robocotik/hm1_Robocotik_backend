package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){

	
	counter := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	nums := strings.Split(line, " ")
	for i := 0; i < len(nums); i++{
		counter[nums[i]]++
	}
	fmt.Println(len(counter))
	
}
