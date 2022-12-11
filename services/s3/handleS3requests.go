package s3

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/cloud-devops101/aws-lib/services/awsclients"
	"github.com/cloud-devops101/aws-lib/services/config"
)

type S3Service struct {
	lambdaConf config.LambdaConfig
	S3Client   awsclients.S3Client
}

func CreateS3Client(lambdaConf config.LambdaConfig) S3Service {
	return S3Service{
		lambdaConf: lambdaConf,
		S3Client:   s3.NewFromConfig(lambdaConf.Cfg),
	}
}

func (svc *S3Service) GetS3Object(bucketName string, keyName string) (s3.GetObjectOutput, error) {

	output, err := svc.S3Client.GetObject(
		svc.lambdaConf.Ctx,
		&s3.GetObjectInput{
			Bucket: &bucketName,
			Key:    &keyName,
		},
	)
	if err != nil {
		svc.lambdaConf.Logger.Printf("Unable to get the object from %s/%s", bucketName, keyName)
		return s3.GetObjectOutput{}, err
	}
	defer output.Body.Close()

	return *output, nil
}

func (svc *S3Service) GetS3ObjectWithFileUnzip(bucketName string, keyName string) (io.Reader, error) {

	output, err := svc.S3Client.GetObject(
		svc.lambdaConf.Ctx,
		&s3.GetObjectInput{
			Bucket: &bucketName,
			Key:    &keyName,
		},
	)
	if err != nil {
		svc.lambdaConf.Logger.Printf("Unable to get the object from %s/%s", bucketName, keyName)
		return nil, err
	}

	gz_reader, err := gzip.NewReader(output.Body)
	if err != nil {
		svc.lambdaConf.Logger.Printf("Unable to unzip s3 object : %s/%s", bucketName, keyName)
		return nil, err
	}
	ioutilOutput, err := ioutil.ReadAll(gz_reader)
	if err != nil {
		svc.lambdaConf.Logger.Printf("Unable to read data from gz Reader for s3 object %s/%s", bucketName, keyName)
		return nil, err
	}
	stringIoReader := strings.NewReader(string(ioutilOutput))

	return stringIoReader, nil
}
