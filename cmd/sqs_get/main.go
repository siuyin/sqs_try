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
	if err != nil {
		log.Fatalf("could not create session: %v", err)
	}

	// service client
	svc := sqs.New(sess)

	// Get an SQS URL given a queue name.
	q, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("junk"),
	})
	if err != nil {
		log.Fatalf("could not get queue URL: %v", err)
	}
	fmt.Printf("junk queue URL: %s\n", *q.QueueUrl)

	// Receive a message from the queue.
	m, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:        q.QueueUrl,
		WaitTimeSeconds: aws.Int64(10),
	})
	if err != nil {
		log.Fatalf("could not retrieve message(s): %v", err)
	}
	fmt.Printf("Received %d messages.\n", len(m.Messages))
	if len(m.Messages) > 0 {
		for i := range m.Messages {
			fmt.Println(m.Messages[i])
		}
	}

	// Delete message from queue.
	for i := range m.Messages {
		d, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
			QueueUrl:      q.QueueUrl,
			ReceiptHandle: m.Messages[i].ReceiptHandle,
		})
		if err != nil {
			fmt.Println("Delete Error", err)
			return
		}
		fmt.Println("Message Deleted", d)
	}

}
