package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsInfra(t *testing.T) {
	awsRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",

		EnvVars: map[string]string{
      "AWS_DEFAULT_REGION": awsRegion,
    },
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketId := terraform.Output(t, terraformOptions, "bucket_id")
	aws.AssertS3BucketExists(t, awsRegion, bucketId)

	test1 := aws.GetS3ObjectContents(t, awsRegion, bucketId, "test1.txt")
	test1Expected := terraform.Output(t, terraformOptions, "test1_content")
	assert.Equal(t, test1, test1Expected, "they should be equal")

	test2 := aws.GetS3ObjectContents(t, awsRegion, bucketId, "test2.txt")
	test2Expected := terraform.Output(t, terraformOptions, "test2_content")
	assert.Equal(t, test2, test2Expected, "they should be equal")	
}
