package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Days    string    `json:"days"`
	Clicks []Click `json:"clicks"`
}

type Click struct {
	clicks int     `json:"clicks"`
	dayStart int   `json:"day_start"`
}




func main() {
	response, err := http.Get("https://api-ssl.bitly.com/v3/user/clicks?access_token=702def740f210a0891078cb3d7277f3c65385efd")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

var responseObject Response
json.Unmarshal(responseData, &responseObject)

fmt.Println(responseObject.Days)
fmt.Println(responseObject.Clicks)
for i := 0; i < len(responseObject.Clicks); i++ {
  fmt.Println(responseObject.Clicks[i].clicks)
}

}
