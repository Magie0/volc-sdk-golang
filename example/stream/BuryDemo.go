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

	var req stream.WapRegisterRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.Uuid = "1695657902873656"
	req.Ouid = "12138"
	req.Partner = "tmxw_api"

	response,err := streamService.WapRegister(req)
	if err!=nil {
		fmt.Println("Error")
		return 
	}
	mar,_ := json.Marshal(response)
	
	fmt.Println(string(mar))
}
