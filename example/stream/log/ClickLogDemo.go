package main

import (
	"encoding/json"
	"fmt"
	"time"
	
	stream_s "github.com/volcengine/volc-sdk-golang/service/stream"
	"github.com/volcengine/volc-sdk-golang/service/stream/log/log"
)

func main() {
	streamService := stream_s.NewStreamService()

	streamService.Client.SetAccessKey("AKLTY2IwYmI4NWI2NGE2NDU0MDgwNDkxN2ZlYzRjYjZkMDQ")
	streamService.Client.SetSecretKey("TldZd09XVmhOMkl5TUdJNU5HRXdPV0kyWTJOak1HSmhPR1UwTXpjd1lqTQ==")

	var req log.ClickLogRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "4016360944571705153676044295679a"
	req.GroupId = "6938426748178530823"
	req.Partner = "vivoliulanqi"
	req.Category = ""
	req.EventTime = "1626861782"
	req.Dt = "iphone6s"
	req.Os = "iOS"
	req.OsVersion = "12"
	req.ClientVersion = "7.3.25"
	req.DeviceBrand = "Apple"

	response, err := log_s.LogService
	if err != nil {
		fmt.Println("Error")
		return
	}
	mar, _ := json.Marshal(response)

	fmt.Println(string(mar))
}
