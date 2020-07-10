package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatal(err)
	}

	cfg.Region = "ap-southeast-1"

	svc0 := sts.New(cfg)

	req0 := svc0.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})

	result0, err := req0.Send(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result0)

	svc := ecr.New(cfg)

	req := svc.GetAuthorizationTokenRequest(&ecr.GetAuthorizationTokenInput{})

	result, err := req.Send(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.String())

	err = ioutil.WriteFile("docker-password", []byte(result.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
