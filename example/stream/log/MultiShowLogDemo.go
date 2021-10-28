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

	var req log.MultiShowLogRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "4016360944571705153676044295679a"
	req.Partner = "vivoliulanqi"

	var bodyList []log.MultiShowLogRequestBody

	var body log.MultiShowLogRequestBody
	body.GroupId = "7008339270066766349"
	body.Category = "fhh_app_default_content"
	body.EventTime = int64(time.Now().Local().Unix())
	body.FromGid = "6938426748178530823"
	body.Dt = "iphone6s"
	body.Os = "iOS"
	body.OsVersion = "1.0.0"
	body.ClientVersion = "7.3.25"
	body.DeviceBrand = "Apple"
	bodyList = append(bodyList, body)

	var body1 log.MultiShowLogRequestBody
	body1.GroupId = "7008339270066766349"
	body1.Category = "fhh_app_default_content"
	body1.EventTime = 1626861782
	body1.FromGid = "6938426748178530823"
	body1.Dt = "iphone6s"
	body1.Os = "iOS"
	body1.OsVersion = "1.0.0"
	body1.ClientVersion = "7.3.25"
	body1.DeviceBrand = "Apple"
	bodyList = append(bodyList, body1)

	req.Body = bodyList

	response, err := logService.MultiShowLog(req)
	if err != nil {
		fmt.Println("Error")
		return
	}
	mar, _ := json.Marshal(response)

	fmt.Println(string(mar))
}
*/