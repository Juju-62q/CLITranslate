package main

import(
	"os"
	"net/http"
	"bufio"
	"io/ioutil"
)

const(
	translateTextAPI = "https://api.cognitive.microsoft.com/sts/v1.0/issueToken"
)

func getAccessToken() (string, error){
	translateAccessKey := getSubscriptionKey()
	req, err := http.NewRequest("POST", translateTextAPI,nil)
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