package client

import (
	
	"github.com/omniful/go_commons/log"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	appConfig "github.com/omniful/go_commons/config"
)

type S3Client struct {
	Client *s3.Client
	Bucket string
}

var s3Client *s3.Client

func NewS3Client(ctx context.Context) (*S3Client, error) {
	region := appConfig.GetString(ctx, "aws.region")
	endpoint := appConfig.GetString(ctx, "aws.endpoint")
	bucket := appConfig.GetString(ctx, "aws.s3.bucket_name")

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				if service == s3.ServiceID {
					return aws.Endpoint{
						URL:               endpoint,
						HostnameImmutable: true,
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			},
		)),
	)
	if err != nil {
		log.Errorf("❌ AWS config load failed: %v", err)
		return nil, err
	}

	s3Client = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true // Needed for LocalStack
	})

	return &S3Client{
		Client: s3Client,
		Bucket: bucket,
	}, nil
}


func GetS3Client()*s3.Client{
	return s3Client
}

