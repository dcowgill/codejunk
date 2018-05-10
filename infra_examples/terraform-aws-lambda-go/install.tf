provider "aws" {
  region = "us-east-1"
}

# Create the Lambda function using the locally compiled artifact.
resource "aws_lambda_function" "demo_lambda" {
  function_name = "demo_lambda"
  handler = "main"
  role = "${aws_iam_role.demo_lambda_exec_role.arn}"
  runtime = "go1.x"

  source_code_hash = "${base64sha256(file("build/deployment.zip"))}"
  filename = "build/deployment.zip" # TODO: use s3_bucket/s3_key instead

  memory_size = 256
  timeout = 300

  environment {
    variables = {
      AWS_S3_BUCKET = "rds-error-log-backup"
    }
  }
}

# Create an IAM role policy specifying the permissions the function requires.
resource "aws_iam_role_policy" "demo_lambda_role_policy" {
  name = "demo_lambda_role_policy"
  role = "${aws_iam_role.demo_lambda_exec_role.id}"
  policy = <<EOF
{
    "Statement": [
        {
            "Action": [
                "s3:PutObject",
                "s3:GetObject",
                "rds:DownloadDBLogFilePortion",
                "s3:ListBucket",
                "rds:DescribeDBLogFiles"
            ],
            "Effect": "Allow",
            "Resource": [
                "arn:aws:s3:::rds-error-log-backup",
                "arn:aws:s3:::rds-error-log-backup/*",
                "arn:aws:rds:*:*:db:*"
            ]
        },
        {
            "Action": [
                "logs:CreateLogStream",
                "logs:DescribeLogStreams",
                "logs:PutLogEvents"
            ],
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Action": [
                "rds:DescribeDBInstances",
                "rds:DownloadCompleteDBLogFile"
            ],
            "Effect": "Allow",
            "Resource": "*"
        },
        {
            "Action": "logs:CreateLogGroup",
            "Effect": "Allow",
            "Resource": "*"
        }
    ],
    "Version": "2012-10-17"
}
EOF
}

# Create the IAM role that will execute the function.
resource "aws_iam_role" "demo_lambda_exec_role" {
  name = "demo_lambda_exec_role"
  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "lambda.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}

# Create a CloudWatch event to trigger the function on a schedule.
resource "aws_cloudwatch_event_rule" "demo_lambda_schedule" {
  name = "demo_lambda_schedule"
  depends_on = ["aws_lambda_function.demo_lambda"]
  schedule_expression = "rate(10 minutes)"
}

# Connect the event to the function.
resource "aws_cloudwatch_event_target" "run_demo_lambda_schedule" {
    rule = "${aws_cloudwatch_event_rule.demo_lambda_schedule.name}"
    target_id = "${aws_lambda_function.demo_lambda.id}"
    arn = "${aws_lambda_function.demo_lambda.arn}"
}

# Grant permission to the function to access the event.
resource "aws_lambda_permission" "allow_cloudwatch_to_run_demo_lambda_schedule" {
    statement_id = "allow_cloudwatch_to_run_demo_lambda_schedule"
    action = "lambda:InvokeFunction"
    function_name = "${aws_lambda_function.demo_lambda.function_name}"
    principal = "events.amazonaws.com"
    source_arn = "${aws_cloudwatch_event_rule.demo_lambda_schedule.arn}"
}
