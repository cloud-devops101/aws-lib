package evb

import (
	"encoding/json"

	"github.com/cloud-devops101/aws-lib/services/config"

	"github.com/aws/aws-lambda-go/events"
)

type EVBService struct {
	lambdaConf config.LambdaConfig
}

type S3Event struct {
	events.CloudWatchEvent
	Detail events.S3Event `json:"detail"`
}

func GetS3EventFromEVB(event events.CloudWatchEvent) (events.S3Event, error) {
	s3Event := events.S3Event{}
	err := json.Unmarshal(event.Detail, s3Event)
	if err != nil {
		return s3Event, err
	}
	return s3Event, nil
}
