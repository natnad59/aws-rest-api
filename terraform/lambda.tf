resource "aws_lambda_function" "_" {
  function_name = "rest-api-example"
  role          = aws_iam_role._.arn
  s3_bucket     = "backend-dev-artifacts"
  s3_key        = "examples/rest-api-example.zip"
  runtime       = "provided.al2023"
  handler       = "api"
}

resource "aws_iam_role" "_" {
  name                = "rest-api-example-role"
  managed_policy_arns = ["arn:${data.aws_partition.current.partition}:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"]

  depends_on = [ aws_dynamodb_table._ ]
  assume_role_policy = jsonencode(
    {
      Version = "2012-10-17"
      Statement = [
        {
          Effect    = "Allow"
          Principal = { Service = "lambda.amazonaws.com" }
          Action    = ["sts:AssumeRole"]
        }
      ]
    }
  )

  inline_policy {
    name = "AllowDynamoDb"
    policy = jsonencode(
      {
        Version = "2012-10-17"
        Statement = [
          {
            Effect = "Allow"
            Action = [
              "dynamodb:GetItem",
              "dynamodb:PutItem",
              "dynamodb:UpdateItem",
              "dynamodb:DeleteItem"
            ]
            Resource = [
              aws_dynamodb_table._.arn,
              "${aws_dynamodb_table._.arn}/*"
            ]
          }
        ]
      }
    )
  }
}
