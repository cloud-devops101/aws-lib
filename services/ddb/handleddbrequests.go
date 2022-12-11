package ddb

import (
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamodb_types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/cloud-devops101/aws-lib/services/awsclients"
	"github.com/cloud-devops101/aws-lib/services/config"
)

type DDBService struct {
	lambdaConf     config.LambdaConfig
	DynamodbClient awsclients.DynamodbClient
}

const MAX_ITEMS_SCAN_PER_OPERATION = 500

func CreateDDBClient(lambdaConf config.LambdaConfig) DDBService {
	return DDBService{
		lambdaConf:     lambdaConf,
		DynamodbClient: dynamodb.NewFromConfig(lambdaConf.Cfg),
	}
}

func (svc *DDBService) ScanDDBFilter(tableName string, exp string, expAttVals map[string]dynamodb_types.AttributeValue) ([]map[string]dynamodb_types.AttributeValue, error) {
	svc.lambdaConf.Logger.Printf("scanning started for %v on %s \n ", exp, tableName)

	allOutputItems := []map[string]dynamodb_types.AttributeValue{}
	pagination := map[string]dynamodb_types.AttributeValue{}

	output, err := svc.DynamodbClient.Scan(svc.lambdaConf.Ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(tableName),
		FilterExpression:          aws.String(exp),
		ExpressionAttributeValues: expAttVals,
		Limit:                     aws.Int32(MAX_ITEMS_SCAN_PER_OPERATION),
	})
	if err != nil {
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			svc.lambdaConf.Logger.Printf("DynamoDB Table scan failed with ConditionalCheckException: %s", err.Error())
			return output.Items, nil
		}
		svc.lambdaConf.Logger.Printf("Could not Scan Campaign Object %s", err.Error())
		return nil, err
	}

	pagination = output.LastEvaluatedKey
	allOutputItems = append(allOutputItems, output.Items...)
	for pagination != nil {
		next_scan_output, err := svc.DynamodbClient.Scan(svc.lambdaConf.Ctx, &dynamodb.ScanInput{
			TableName:                 aws.String(tableName),
			FilterExpression:          aws.String(exp),
			ExpressionAttributeValues: expAttVals,
			ExclusiveStartKey:         pagination,
			Limit:                     aws.Int32(MAX_ITEMS_SCAN_PER_OPERATION),
		})
		if err != nil {
			if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
				svc.lambdaConf.Logger.Printf("DynamoDB Table scan failed with ConditionalCheckException: %s", err.Error())
				return output.Items, nil
			}
			svc.lambdaConf.Logger.Printf("Could not Scan Campaign Object %s", err.Error())
			return nil, err
		}
		allOutputItems = append(allOutputItems, next_scan_output.Items...)
		pagination = next_scan_output.LastEvaluatedKey
	}
	return allOutputItems, nil
}

func (svc *DDBService) ScanDDB(tableName string) ([]map[string]dynamodb_types.AttributeValue, error) {
	svc.lambdaConf.Logger.Printf("scanning started on %s \n ", tableName)

	allOutputItems := []map[string]dynamodb_types.AttributeValue{}
	pagination := map[string]dynamodb_types.AttributeValue{}

	output, err := svc.DynamodbClient.Scan(svc.lambdaConf.Ctx, &dynamodb.ScanInput{
		TableName: aws.String(tableName),
		Limit:     aws.Int32(MAX_ITEMS_SCAN_PER_OPERATION),
	})
	if err != nil {
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			svc.lambdaConf.Logger.Printf("DynamoDB Table scan failed with ConditionalCheckException: %s", err.Error())
			return output.Items, nil
		}
		svc.lambdaConf.Logger.Printf("Could not Scan Campaign Object %s", err.Error())
		return nil, err
	}

	pagination = output.LastEvaluatedKey
	allOutputItems = append(allOutputItems, output.Items...)
	for pagination != nil {
		next_scan_output, err := svc.DynamodbClient.Scan(svc.lambdaConf.Ctx, &dynamodb.ScanInput{
			TableName:         aws.String(tableName),
			ExclusiveStartKey: pagination,
			Limit:             aws.Int32(MAX_ITEMS_SCAN_PER_OPERATION),
		})
		if err != nil {
			if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
				svc.lambdaConf.Logger.Printf("DynamoDB Table scan failed with ConditionalCheckException: %s", err.Error())
				return output.Items, nil
			}
			svc.lambdaConf.Logger.Printf("Could not Scan Campaign Object %s", err.Error())
			return nil, err
		}
		allOutputItems = append(allOutputItems, next_scan_output.Items...)
		pagination = next_scan_output.LastEvaluatedKey
	}
	return allOutputItems, nil
}
