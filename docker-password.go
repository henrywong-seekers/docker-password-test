package main

import (
	"context"
	"encoding/base64"
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

	authorizationToken := *result.AuthorizationData[0].AuthorizationToken

	data, err := base64.StdEncoding.DecodeString(authorizationToken)
	if err != nil {
		log.Fatal(err)
	}

	dockerPassword := data[4:]

	err = ioutil.WriteFile("docker-password", []byte(dockerPassword), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
