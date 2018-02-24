package main

import(
	"os"
	"net/http"
	"bufio"
	"io/ioutil"
	"net/url"
	"github.com/PuerkitoBio/goquery"
	"path/filepath"
)

const(
	getTokenURL = "https://api.cognitive.microsoft.com/sts/v1.0/issueToken"
	textTranslateURL = "https://api.microsofttranslator.com/V2/Http.svc/Translate?"
)

// 翻訳用の関数
func translateString(s string, from string, to string) string{
	token, err := getAccessToken()
	// Cognitive にアクセスできない場合
	if err != nil{
		return "can't access to cognitive service"
	}

	// urlの成形
	appid := url.QueryEscape("Bearer "+ string(token))
	text := url.QueryEscape(s)
	textTranslateApiUrl := textTranslateURL+"from="+from+"&to="+to+"&text="+text+"&appid="+appid

	// getRequestの発行と結果の取得
	doc, err := goquery.NewDocument(textTranslateApiUrl)

	// 500等何らかのエラーが発生した
	if err != nil{
		return "can't get appropriate response"
	}

	// Stringタグのデータをスクレイピングする
	response := doc.Find("string").Text()

	return response
}


// アクセストークンの発行
func getAccessToken() (string, error){
	// アクセスキーの取得
	translateAccessKey := getSubscriptionKey()

	// トークン取得用のポストメソッド作成
	req, err := http.NewRequest("POST", getTokenURL,nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", translateAccessKey)

	// 何らかのエラーが発生
	if err != nil{
		return "", err
	}

	// Requestの実行
	client := new(http.Client)
	response, _ := client.Do(req)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	return string(body), err
}

// サブスクキーのファイル読み込み(実行ファイルと同一ディレクトリに存在)
func getSubscriptionKey() string{
	var fp *os.File
	var err error

	// 実行ファイルの配置されているディレクトリの取得
	exe, err := os.Executable()
	path := filepath.Dir(exe)

	// エラー処理
	if err != nil {
		panic(err)
	}

	// ファイル読み込み
	fp, err = os.Open(path + "/.SubscriptionKey")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
	line, _, err := reader.ReadLine()

	return string(line)
}