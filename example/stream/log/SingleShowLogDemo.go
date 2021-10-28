package main
/*
import (
	"encoding/json"
	"fmt"
	"time"

	log_s "github.com/volcengine/volc-sdk-golang/service/stream/log"
	"github.com/volcengine/volc-sdk-golang/service/stream/log/log"
)

func main() {
	logService := log_s.NewStreamService()

	logService.Client.SetAccessKey("AKLTY2IwYmI4NWI2NGE2NDU0MDgwNDkxN2ZlYzRjYjZkMDQ")
	logService.Client.SetSecretKey("TldZd09XVmhOMkl5TUdJNU5HRXdPV0kyWTJOak1HSmhPR1UwTXpjd1lqTQ==")

	var req log.SingleShowLogRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "4016360944571705153676044295679a"
	req.GroupId = "6938426748178530823"
	req.Partner = "vivoliulanqi"
	req.Category = "__ALL__"
	req.EventTime = "1626861782"
	req.Dt = "iphone6s"
	req.Os = "iOS"
	req.OsVersion = "12"
	req.ClientVersion = "7.3.25"
	req.DeviceBrand = "Apple"

	response, err := logService.SingleShowLog(req)
	if err != nil {
		fmt.Println("Error")
		return
	}
	mar, _ := json.Marshal(response)

	fmt.Println(string(mar))
}
*/