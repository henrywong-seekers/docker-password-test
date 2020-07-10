package main

import (
	"context"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Region = "ap-southeast-1"

	svc := ecr.New(cfg)

	req := svc.GetAuthorizationTokenRequest(&ecr.GetAuthorizationTokenInput{})

	result, err := req.Send(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("docker-password", []byte(result.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
