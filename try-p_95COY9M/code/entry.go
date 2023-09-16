package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	ShelterTemplate int    `json:"shelterTemplate"`
	KennelCardsUrl  string `json:"kennelCardsUrl"`
}

func main() {
	url := "https://www.shelterluv.com/api/v3/kennel_cards/get-shelter-template"
	method := "POST"
	payload := strings.NewReader("_csrfToken=mf3apQIxOahZ_r3UFqPeaiRxPDTJIev5yhmm0Y8z9Q4&animal_ids=74325299")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add headers here
	req.Header.Add("sec-ch-ua", "\"Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"115\", \"Chromium\";v=\"115\"")
	req.Header.Add("Accept", "text/plain, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("host", "www.shelterluv.com")
	req.Header.Add("X-Api-Key", "61a29cf6-cc4d-42d0-892c-a585dd30e3f8")
	req.Header.Add("Cookie", "__hssrc=1; __hstc=148675048.12b3b651a03d87361bbfbe1281c22cfe.1694806288679.1694806288679.1694806288679.1; __zlcmid=1Hlm6ro5ZU0CLJ9; _fbp=fb.1.1694806288148.107900607; _ga=GA1.2.1824892527.1692831979; _ga_4L6NCQ6CSL=GS1.1.1694806234.7.1.1694806285.0.0.0; _gid=GA1.2.2093128975.1694696472; _hjSessionUser_2611551=eyJpZCI6IjJlNTJjMjI4LThiMjEtNTA3Yy1iNzliLTgyMGIyNDczMmNlMSIsImNyZWF0ZWQiOjE2OTQyNzU5NzY0NDMsImV4aXN0aW5nIjp0cnVlfQ==; hubspotutk=12b3b651a03d87361bbfbe1281c22cfe; SESSb4bb675ba22f5c3c57e8f975aa787735=221HQ4Ky2Jzigf2qTN2s1tvjAanQ10td02flojtQp34")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Received non-200 status code:", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var resBody Response
	if err := json.Unmarshal(body, &resBody); err != nil {
		fmt.Println("Error unmarshaling response:", err)
		return
	}

	// Speculative code to set a scoped variable in Pipedream
	// Replace `steps` and `setVariable` with the actual functions or methods provided by Pipedream's Go environment
	// steps.setVariable("kennelCardsUrl", resBody.KennelCardsUrl)

	fmt.Println("Kennel Cards URL:", resBody.KennelCardsUrl)
}
