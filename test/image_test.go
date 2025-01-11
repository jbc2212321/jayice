package test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var accessToken = "24.c7906b238118d61128c56b219adc3567.2592000.1650544920.282335-25820164"

func TestImage(t *testing.T) {
	var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()

	filebytes, err := ioutil.ReadFile("../image/7782a9ae-8dba-4a18-bcc0-40016ad7182a.png")
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}

func TestImageWithPosition(t *testing.T) {
	var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()

	filebytes, err := ioutil.ReadFile("../image/7782a9ae-8dba-4a18-bcc0-40016ad7182a.png")
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}

func TestTicket(t *testing.T) {

	//通用票据识别 https://cloud.baidu.com/apiexplorer/index.html?Product=GWSE-DJAQ8YwekkQ&Api=GWAI-ocX4vUosqfg
	var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/receipt"

	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()

	filebytes, err := ioutil.ReadFile("../image/Y9R}WGN]HY}M[7B8Q{3GXCI.jpg")
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}

//增值税发票
func TestInvoice(t *testing.T) {
	var host = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"
	uri, err := url.Parse(host)
	if err != nil {
		fmt.Println(err)
	}
	query := uri.Query()
	query.Set("access_token", accessToken)
	uri.RawQuery = query.Encode()

	filebytes, err := ioutil.ReadFile("../image/2e3631ccbfd2850a668ece07ff63811d.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	sendBody := http.Request{}
	sendBody.ParseForm()
	sendBody.Form.Add("image", image)
	sendData := sendBody.Form.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", uri.String(), strings.NewReader(sendData))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(result))

}

func TestParseForm(t *testing.T) {
	filebytes, err := ioutil.ReadFile("C:\\Users\\78240\\go\\ticket\\image\\vat_invoice.jpeg")
	if err != nil {
		fmt.Println(err)
	}

	image := base64.StdEncoding.EncodeToString(filebytes)
	fmt.Println(image)
}
