package domain

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ConnectAws() (connection *session.Session, err error) {
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	endpoint := os.Getenv("AWS_ENDPOINT")
	connection, err = session.NewSession(
		&aws.Config{
			Endpoint: aws.String(endpoint),
			Region:   aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				accessKeyID,
				secretAccessKey,
				"", // a token will be created when the session is used.
			),
		})
	if err != nil {
		return nil, err
	}

	return connection, nil
}
