package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// 入力が引数か標準入力かどうかのフラグ
	var useStandardInput = false
	// オプションの数を示す値
	var optNum = 0

	// オプション関係の設定を読まないようにするための配列
	isNotNeedRead := make([]bool, len(os.Args), len(os.Args))
	isNotNeedRead[0] = true

	sourceLang := ""
	var fOptionPos = arrayContains(os.Args, "-f")
	// -fがなければ元言語は日本語
	if (fOptionPos < 0) {
		sourceLang = "ja"
	// -fがある場合にはオプションの数をインクリメント
	}else{
		optNum++
	}

	targetLang := ""
	var tOptionPos = arrayContains(os.Args, "-t")
	// -t がなければ翻訳先言語は英語
	if (tOptionPos < 0) {
		targetLang = "en"
	// -t がある場合にはオプションの数をインクリメント
	}else{
		optNum++
	}

	// 引数が不足する場合のエラー処理
	if len(os.Args) - optNum*2 <= 0 {
		fmt.Println("not enough argment")
		return
	// 翻訳する言葉が指定されていない場合には標準入力を利用
	}else if(len(os.Args) - optNum*2 <= 1) {
		useStandardInput = true
	}

	// 元言語についての処理
	if (fOptionPos >= 0){
		sourceLang = os.Args[fOptionPos + 1]
		isNotNeedRead[fOptionPos] = true
		isNotNeedRead[fOptionPos + 1] = true
	}

	// 翻訳先言語についての処理
	if (tOptionPos >= 0){
		targetLang = os.Args[tOptionPos + 1]
		isNotNeedRead[tOptionPos] = true
		isNotNeedRead[tOptionPos + 1] = true
		fmt.Println(targetLang + " detected")
	}

	// stdinを利用
	if(useStandardInput){
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan(){
			fmt.Println(translateString(stdin.Text(), sourceLang))
		}
	// 引数で指定されている
	}else{
		for i, v := range(isNotNeedRead){
			if(!v){
				fmt.Println(translateString(os.Args[i], sourceLang))
			}
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

