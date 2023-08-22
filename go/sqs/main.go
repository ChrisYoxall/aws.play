package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// See documentation at https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sqs and an
// example at https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/gov2/sqs/ReceiveMessage

func main() {

	queueURL := "https://sqs.ap-southeast-2.amazonaws.com/940626918307/TestSqsQueue"
	timeout := int32(60)

	gMInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   timeout,
	}

	msgResult, err := ReceiveMessages(gMInput)
	if err != nil {
		fmt.Println("Error receiving messages:")
		fmt.Println(err)
		return
	}

	if msgResult.Messages != nil {
		fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
		fmt.Println("Message Body:   " + *msgResult.Messages[0].Body)
		fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle) // required to delete message
	} else {
		fmt.Println("No messages found")
	}

}
