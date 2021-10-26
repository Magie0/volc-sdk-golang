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

	var req stream.DiggRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "4016360944571705153676044295679a"
	req.Partner = "vivoliulanqi"
	req.GroupId = "7008338242986623518"

	response,err := streamService.UnDigg(req)
	if err!=nil {
		fmt.Println("Error")
		return 
	}
	mar,_ := json.Marshal(response)
	
	fmt.Println(string(mar))
}
*/