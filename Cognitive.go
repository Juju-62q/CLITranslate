package main

import(
	"os"
	"net/http"
	"bufio"
	"io/ioutil"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

const(
	getTokenURL = "https://api.cognitive.microsoft.com/sts/v1.0/issueToken"
	textTranslateURL = "https://api.microsofttranslator.com/V2/Http.svc/Translate?"
)

func translateString(s string, from string) string{
	token, err := getAccessToken()
	if err != nil{
		return "can't access to cognitive service"
	}
	appid := url.QueryEscape("Bearer "+ string(token))
	text := url.QueryEscape(s)

	url := textTranslateURL+"from="+from+"&to=en&text="+text+"&appid="+appid

	doc, err := goquery.NewDocument(url)

	if err != nil{
		return "can't get appropriate response"
	}

	response := doc.Find("string").Text()

	return response
}



func getAccessToken() (string, error){
	translateAccessKey := getSubscriptionKey()
	req, err := http.NewRequest("POST", getTokenURL,nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", translateAccessKey)

	if err != nil{
		return "", err
	}

	client := new(http.Client)
	response, _ := client.Do(req)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	return string(body), err
}

func getSubscriptionKey() string{
	var fp *os.File
	var err error
	fp, err = os.Open("./.SubscriptionKey")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)
	line, _, err := reader.ReadLine()

	return string(line)
}