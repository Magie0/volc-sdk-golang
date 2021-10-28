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

	var req log.DislikeLogRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "4016360944571705153676044295679a"
	req.Partner = "vivoliulanqi"
	var bodyList  []*log.DislikeLogRequestBody

	var body log.DislikeLogRequestBody
	body.GroupId = "7008339270066766349"
	body.Timestamp = int64(time.Now().Local().Unix())
	body.Category = "fhh_app_default_content"

	var filterWords []string
	filterWords = append(filterWords, "id1")
	filterWords = append(filterWords, "id2")
	body.FilterWords = filterWords
	bodyList = append(bodyList, &body)
	
	var body1 log.DislikeLogRequestBody

	body1.GroupId = "7008339270066766349"
	body1.Timestamp = int64(time.Now().Local().Unix())
	body1.Category = "fhh_app_default_content"
	var filterWords1 []string
	filterWords1 = append(filterWords1, "id1")
	filterWords1 = append(filterWords1, "id2")
	body1.FilterWords = filterWords1
	bodyList = append(bodyList, &body1)

	req.Body = bodyList

	response, err := logService.DislikeLog(req)
	if err != nil {
		fmt.Println("Error")
		return
	}
	mar, _ := json.Marshal(response)

	fmt.Println(string(mar))
}
*/