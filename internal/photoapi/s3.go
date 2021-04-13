package photoapi

import (
	_ "image/jpeg"

	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3ImageSource struct {
	Region     string
	BucketName string
	Prefix     string
}

func (s S3ImageSource) GetImage(photoData *PhotoData) error {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String(s.Region)})

	objectOutput, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(strings.Join([]string{s.Prefix, photoData.ID}, "/")),
	})

	photoData.Reader = objectOutput.Body

	return err
}
