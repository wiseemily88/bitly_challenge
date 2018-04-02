package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

type Bitlyapi struct {
  Code int  `json:"status_code"`
  data struct {
    Days int `json:"days"`
    TotalClicks int `json:"total_clicks"`
    clicks struct {
      clicks int `json:"clicks"`
      dayStart int `json:"day_start"`
      }`json:"quotes"`
      Units int `json:"units"`
      UnitReferenceTs int `json:"unit_reference_ts"`
      UserClicks struct {
        Dt int `json:"dt"`
        clicks int `json:"clicks"`
        }`json:"quotes"`
        TimzoneOffset int `json:"tz_offset"`
        Unit string `json"unit"`
      }
      Status string `json:"status_txt"`
    }

func main() {

	// YOUR_ACCESS_KEY := 702def740f210a0891078cb3d7277f3c65385efd

	url := fmt.Sprintf("https://api-ssl.bitly.com/v3/user/clicks?access_token=702def740f210a0891078cb3d7277f3c65385efd")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}


client := &http.Client{}
resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
defer resp.Body.Close()

var record Bitlyapi
if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Unit = ", record.data.Unit)
	fmt.Println("data", record.data)
	fmt.Println("Status Code = ", record.Code)
	fmt.Println("Days = ", record.data.Days)
	fmt.Println("TotalClicks = ", record.data.TotalClicks)
	fmt.Println("Clicks = ", record.data.clicks.clicks)
	fmt.Println("Day Start = ", record.data.clicks.dayStart)
	fmt.Println("Units   = ", record.data.Units)
	fmt.Println("Unit   = ", record.data.Unit)
	fmt.Println("Timzone   = ", record.data.TimzoneOffset)
	fmt.Println("Days  = ", record.data.Days)
	fmt.Println("User Clicks   = ", record.data.UserClicks.clicks)
	fmt.Println("User Clicks   = ", record.data.UserClicks.Dt)

}
