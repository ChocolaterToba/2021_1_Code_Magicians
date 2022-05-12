package s3client

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	s3session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type S3ClientInterface interface {
	UploadFile(ctx context.Context, filename string, file *bytes.Buffer) (err error)
	DeleteFile(ctx context.Context, filename string) (err error)
}

type S3Client struct {
	session    *s3session.Session
	bucketName string
}

func NewS3Client(session *s3session.Session, bucketName string) *S3Client {
	return &S3Client{
		session:    session,
		bucketName: bucketName,
	}
}

func (client *S3Client) UploadFile(ctx context.Context, filename string, file *bytes.Buffer) (err error) {
	uploader := s3manager.NewUploader(client.session)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(client.bucketName),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		// TODO: add s3 error parsing (for example, name conflict)
		return status.Errorf(codes.Unknown, "cannot upload to s3: %v", err)
	}

	return nil
}

func (client *S3Client) DeleteFile(ctx context.Context, filename string) (err error) {
	svc := s3.New(client.session)

	input := &s3.DeleteObjectInput{
		Bucket: aws.String(client.bucketName),
		Key:    aws.String(filename),
	}

	_, err = svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			// TODO: parse object not existing error
			default:
				return err
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return err
		}
	}

	return nil
}
