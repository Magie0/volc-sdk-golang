package main
/*
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

	var req stream.RefreshTipsRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "4016360944571705153676044295679a"
	req.Partner = "vivoliulanqi"
	req.Ouid = "7008339270066766349"
	req.Cursor = 4

	response,err := streamService.RefreshTips(req)
	if err!=nil {
		fmt.Println("Error")
		return 
	}
	mar,_ := json.Marshal(response)
	
	fmt.Println(string(mar))
}
*/