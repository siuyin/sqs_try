package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/siuyin/dflt"
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

	url := dflt.EnvString("QUEUE_URL", "REPLACE_ME_WITH_ACTUAL_URL")
	fmt.Printf("fifo queue URL: %s\n", url)

	// Receive a message from the queue.
	m, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:        &url,
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
	//for i := range m.Messages {
	//	d, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
	//		QueueUrl:      q.QueueUrl,
	//		ReceiptHandle: m.Messages[i].ReceiptHandle,
	//	})
	//	if err != nil {
	//		fmt.Println("Delete Error", err)
	//		return
	//	}
	//	fmt.Println("Message Deleted", d)
	//}

}
