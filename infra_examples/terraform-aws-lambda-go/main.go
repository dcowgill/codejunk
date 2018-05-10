// TBD
package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.SetFlags(0) // AWS prepends timestamps to log messages
	lambda.Start(LambdaHandler)
}

func LambdaHandler(ctx context.Context) error {
	log.Printf("(version two) executed at %s", time.Now().Format(time.UnixDate))
	return nil
}
