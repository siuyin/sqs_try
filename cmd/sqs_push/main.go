package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	// The Environment variables for credentials will have precedence over shared config even if SharedConfig is enabled.
	fmt.Println("AWS Sessions")
	sess, err := session.NewSession() // Retrive region from environment.
	//sess, err := session.NewSession(&aws.Config{
	//	Region: aws.String("us-east-1")},
	//)
	if err != nil {
		log.Fatalf("could not create session: %v", err)
	}

	// service client
	svc := sqs.New(sess)

	// list urls
	res, err := svc.ListQueues(nil)
	if err != nil {
		log.Fatalf("could not list queues: %v", err)
	}

	for i, urls := range res.QueueUrls {
		if urls == nil {
			continue
		}
		fmt.Printf("%d: %s\n", i, *urls)
	}

	// Get an SQS URL given a queue name.
	q, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("junk"),
	})
	if err != nil {
		log.Fatalf("could not get queue URL: %v", err)
	}
	fmt.Printf("junk queue URL: %s\n", *q.QueueUrl)

	// Send a message into the queue.
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(1),
		QueueUrl:     q.QueueUrl,
		MessageBody:  aws.String("The quick brown fox"),
	})
	if err != nil {
		log.Fatalf("could not send message: %v", err)
	}
	fmt.Printf("message id: %v\n", *result.MessageId)

}
