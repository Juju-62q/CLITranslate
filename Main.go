package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var index = arrayContains(os.Args, "-s")
	if(index < 0){
		fmt.Println("option not found")
	}else{
		fmt.Println("option found at " + strconv.Itoa(index))
	}
	fmt.Println(getAccessToken())
}

func arrayContains(arr []string, s string) int{
	for i, v := range arr{
		if v == s{
			return i
		}
	}
	return -1
}