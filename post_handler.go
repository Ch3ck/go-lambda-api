package main

import (  
"fmt"

"github.com/aws/aws-lambda-go/events"
"github.com/aws/aws-lambda-go/lambda"
)

func Handler(event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {

    fmt.Printf("User Attributes: %#v \n", event)

    return event, nil
}

func main() {  
    lambda.Start(Handler)
}
