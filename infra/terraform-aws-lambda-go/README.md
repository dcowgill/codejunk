# example-terraform-aws-lambda-go

## Building

Install dependencies:
```
go get github.com/aws/aws-lambda-go/lambda
```

Build the deployment artifact:
```
make
```

## Installing

Initial setup:
```
terraform init
```

To install or update the Lambda function:
```
terraform plan -out plan.txt
```

Verify that the plan is sensible, then:
```
terraform apply plan.txt
```
