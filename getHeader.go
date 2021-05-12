// WAP to get the URL and return the value of Content type response HTTP header

package main

import (
	"fmt"
	"net/http"
)

func getURLContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
	  return "", err
	}
	defer response.Body.Close()
	
	content := response.Header.Get("Content-Type")
	if content == "" {
	return "",fmt.Errorf("content not found")
	}
	return content,nil
}

func main() {
	url1 := "https://google.com"
	resp, err := getURLContent(url1)
	if err != nil {
		fmt.Printf("ERR = %s \n", err)
	} else {
		fmt.Println(resp)
	}

}
