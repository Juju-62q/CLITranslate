package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("not enough argment")
		return
	}
	var index = arrayContains(os.Args, "-f")
	if(index < 0){
		fmt.Println(translateString(os.Args[1], "ja"))
	}else{
		if len(os.Args) <= 3 {
			fmt.Println("not enough argment")
			return
		}else {
			fmt.Println(translateString(os.Args[1], os.Args[index + 1]))
		}
	}
}

func arrayContains(arr []string, s string) int{
	for i, v := range arr{
		if v == s{
			return i
		}
	}
	return -1
}