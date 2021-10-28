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

	var req log.FollowLogRequest
	req.Timestamp = int64(time.Now().Local().Unix())
	req.AccessToken = "4016360944571705153676044295679a"
	req.Partner = "vivoliulanqi"

	var bodyList []log.FollowLogRequestBody

	var body log.FollowLogRequestBody
	body.GroupId = "7008339270066766349"
	body.ToUserId = "104792520555"
	body.CategoryName = "fhh_app_default_content"
	body.Source = "article_detail"
	bodyList = append(bodyList, body)

	var body1 log.FollowLogRequestBody
	body1.GroupId = "7008339270066766349"
	body1.ToUserId = "104792520555"
	body1.CategoryName = "fhh_app_default_content"
	body1.Source = "article_detail"
	bodyList = append(bodyList, body1)

	req.Body = body

	response, err := logService.UnfollowLog(req)
	if err != nil {
		fmt.Println("Error")
		return
	}
	mar, _ := json.Marshal(response)

	fmt.Println(string(mar))
}
*/