package wasabis3manager

import "github.com/nikola43/ecoapigorm/utils"

var WasabiS3Client = utils.New(
	utils.GetEnvVariable("S3_ACCESS_KEY"),
	utils.GetEnvVariable("S3_SECRET_KEY"),
	utils.GetEnvVariable("S3_ENDPOINT"),
	utils.GetEnvVariable("S3_BUCKET_REGION"))
