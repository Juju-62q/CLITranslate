package main

import (
	"fmt"
	"os"
)

func main() {
	// オプション関係の設定を読まないようにするための配列
	isNotNeedRead := make([]bool, len(os.Args), len(os.Args))
	isNotNeedRead[0] = true

	// 引数がない場合にはエラーを返す
	if len(os.Args) <= 1 {
		fmt.Println("not enough argment")
		return
	}

	// オプションの処理
	sourceLang := ""
	var index = arrayContains(os.Args, "-f")
	// 元言語についての処理
	if(index < 0) {
		 sourceLang = "ja"
	}else{
		// オプションあるのに引数少ない
		if len(os.Args) <= 3 {
			fmt.Println("not enough argment")
			return
		}else{
			sourceLang = os.Args[index + 1]
			isNotNeedRead[index] = true
			isNotNeedRead[index + 1] = true
		}
	}

	// 翻訳の実行
	for i, v := range(isNotNeedRead){
		if(!v){
			fmt.Println(translateString(os.Args[i], sourceLang))
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

