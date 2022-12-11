package config

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type LambdaConfig struct {
	Ctx    context.Context
	Logger *log.Logger
	Cfg    aws.Config
}

func GetLambdaConfig() (LambdaConfig, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("Error loading AWS config: %s", err.Error())
		return LambdaConfig{}, err
	}
	lambdaConf := LambdaConfig{
		Ctx:    context.TODO(),
		Logger: log.New(os.Stdout, "", log.LstdFlags),
		Cfg:    cfg,
	}
	return lambdaConf, nil
}

//Assumes new role when roleArn , region, sts client is provided
func AssumeRole(roleArn string, Region string, client sts.Client) aws.Config {
	cfg := aws.Config{}

	creds := stscreds.NewAssumeRoleProvider(&client, roleArn, func(o *stscreds.AssumeRoleOptions) {
		o.RoleARN = roleArn
	})

	cfg.Credentials = aws.NewCredentialsCache(creds)
	cfg.Region = Region

	return cfg
}
