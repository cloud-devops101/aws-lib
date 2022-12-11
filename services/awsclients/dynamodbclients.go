package awsclients

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Stub of dynamodb.Client
type DynamodbClient interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

// mocks

type MockDynamodbClient struct {
	GetItemInputs    []dynamodb.GetItemInput
	GetItemOutputs   []dynamodb.GetItemOutput
	GetItemErrors    []error
	PutItemInputs    []dynamodb.PutItemInput
	PutItemOutputs   []dynamodb.PutItemOutput
	PutItemErrors    []error
	QueryInputs      []dynamodb.QueryInput
	QueryOutputs     []dynamodb.QueryOutput
	QueryErrors      []error
	ScanInputs       []dynamodb.ScanInput
	ScanOutputs      []dynamodb.ScanOutput
	ScanErrors       []error
	UpdateItemInputs []dynamodb.UpdateItemInput
	UpdateItemOutput []dynamodb.UpdateItemOutput
	UpdateItemErrors []error
}

func (client *MockDynamodbClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	client.PutItemInputs = append(client.PutItemInputs, *params)

	index := len(client.PutItemInputs) - 1

	return &client.PutItemOutputs[index], client.PutItemErrors[index]
}

func (client *MockDynamodbClient) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	client.GetItemInputs = append(client.GetItemInputs, *params)

	index := len(client.GetItemInputs) - 1

	return &client.GetItemOutputs[index], client.GetItemErrors[index]
}

func (client *MockDynamodbClient) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	client.QueryInputs = append(client.QueryInputs, *params)

	index := len(client.QueryInputs) - 1

	return &client.QueryOutputs[index], client.QueryErrors[index]
}

func (client *MockDynamodbClient) Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error) {
	client.ScanInputs = append(client.ScanInputs, *params)

	index := len(client.ScanInputs) - 1

	return &client.ScanOutputs[index], client.ScanErrors[index]
}

func (client *MockDynamodbClient) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	client.UpdateItemInputs = append(client.UpdateItemInputs, *params)

	index := len(client.UpdateItemInputs) - 1

	return &client.UpdateItemOutput[index], client.UpdateItemErrors[index]
}

func (client *MockDynamodbClient) AddUpdateResponse(response dynamodb.UpdateItemOutput, err error) {
	client.UpdateItemOutput = append(client.UpdateItemOutput, response)
	client.UpdateItemErrors = append(client.UpdateItemErrors, err)
}
