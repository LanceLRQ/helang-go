package plugins

import (
	"encoding/json"
	"fmt"
	"helang-go/helang/core"
	"io/ioutil"
	"net/http"
	"regexp"
)

var CyberRegions = []string {
	"UNITED STATES", "JAPAN",
}

func getRegion() (string, error) {
	resp, err := http.Get("https://pv.sohu.com/cityjson?ie=utf-8")
	if err != nil {
		return "", fmt.Errorf("%w: %s", core.CyberNetworkException, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w: %s", core.CyberNetworkException, err.Error())
	}
	data := struct {
		CName string `json:"cname"`
	}{}
	re, err := regexp.Compile("{.+}")
	if err != nil {
		return "", fmt.Errorf("%w: %s", core.CyberNetworkException, err.Error())
	}
	jsonStr := re.Find(body)
	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		return "", fmt.Errorf("%w: %s", core.CyberNetworkException, err.Error())
	}
	return data.CName, nil
}

func CheckCyberSpaces() error {
	fmt.Println("Getting your location...")
	region, err := getRegion()
	if err != nil {
		return err
	}

	fmt.Printf("Your location is %s.\n", region)

	for _, v := range CyberRegions {
		if v == region {
			fmt.Println("Congratulations! You are in the Cyber Spaces!")
			return nil
		}
	}
	fmt.Println("What a pity! It seems that you are not in the Cyber Spaces.")
	return nil
}