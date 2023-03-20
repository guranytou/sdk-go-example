package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	listsBucket(client)
	createBucket("test-bucket-guranytou-2", client)

	listsBucket(client)
}

func createBucket(name string, client *s3.Client) {
	_, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(name),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintApNortheast1,
		},
		ACL: types.BucketCannedACLPrivate,
	})
	if err != nil {
		log.Printf("Couldn't create bucket %v, messages: %v", name, err)
	}
}

func listsBucket(client *s3.Client) {
	reseult, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		log.Printf("error: %v", err)
	}
	fmt.Println("Bucket:")

	for _, b := range reseult.Buckets {
		fmt.Printf("* %s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
