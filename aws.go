package claws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	RegionNV = "us-east-1"
	RegionSG = "ap-southeast-1"
	RegionTH = "ap-southeast-7"
)

// Applies region to default configuration
// This is for local use, production code should have its own config object
func ConfigWithRegion(region string) aws.Config {
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	// Swallowing errors here
	return cfg
}

// Prints the current user's identity
// For testing purposes only
func WhoAmI() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := sts.NewFromConfig(cfg)

	caller, err := client.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Account ID:", aws.ToString(caller.Account))
	fmt.Println("ARN:", aws.ToString(caller.Arn))
	fmt.Println("User ID:", aws.ToString(caller.UserId))
}
