package awsmanager

import "github.com/nikola43/ecoapigorm/utils"

var AwsManager = utils.New(
	utils.GetEnvVariable("AWS_ACCESS_KEY"),
	utils.GetEnvVariable("AWS_SECRET_KEY"),
	utils.GetEnvVariable("AWS_ENDPOINT"),
	utils.GetEnvVariable("AWS_BUCKET_NAME"),
	utils.GetEnvVariable("AWS_BUCKET_REGION"))
