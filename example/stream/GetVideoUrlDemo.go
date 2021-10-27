package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/stream"
	"github.com/volcengine/volc-sdk-golang/service/stream/stream"
)

func main() {
	streamService := stream_s.NewStreamService()

	streamService.Client.SetAccessKey("AKLTY2IwYmI4NWI2NGE2NDU0MDgwNDkxN2ZlYzRjYjZkMDQ")
	streamService.Client.SetSecretKey("TldZd09XVmhOMkl5TUdJNU5HRXdPV0kyWTJOak1HSmhPR1UwTXpjd1lqTQ==")

	var req stream.GetVideoUrlRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "3016420494751676625350562829b71f"
	req.Partner = "mi_mcc_api_2019"
	req.VideoId = "v02004g10000c5h9rhbc77u5t3n7pdt0_7017353858951316006_2_20"

	response,err := streamService.GetVideoUrl(req)
	if err!=nil {
		fmt.Println("Error")
		return 
	}
	mar,_ := json.Marshal(response)
	
	fmt.Println(string(mar))
}