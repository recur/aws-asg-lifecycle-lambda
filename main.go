package main

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var as = autoscaling.New(session.New())

func handler(ctx context.Context, snsEvent events.SNSEvent) (string, error) {
	var msg map[string]string

	if err := json.Unmarshal([]byte(snsEvent.Records[0].SNS.Message), &msg); err != nil {
		return "json unmarshall error", err
	}

	fmt.Printf("%v\n", msg)

	completion := &autoscaling.CompleteLifecycleActionInput{
		AutoScalingGroupName:  aws.String(msg["AutoScalingGroupName"]),
		LifecycleActionToken:  aws.String(msg["LifecycleActionToken"]),
		LifecycleHookName:     aws.String(msg["LifecycleHookName"]),
		LifecycleActionResult: aws.String("CONTINUE"),
	}

	if _, err := as.CompleteLifecycleAction(completion); err != nil {
		return "couldn't complete lifecycle action", err
	}

	return msg["LifecycleActionToken"], nil
}

func main() {
	lambda.Start(handler)
}
