resource "aws_lambda_function" "_" {
  function_name = "rest-api-example"
  role          = aws_iam_role._.arn
  s3_bucket     = data.aws_ssm_parameter.artifacts_bucket.value
  s3_key        = "examples/rest-api-example.zip"
  runtime       = "provided.al2023"
  handler       = "bootstrap"
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
              "dynamodb:Scan",
              "dynamodb:PutItem",
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

resource "aws_lambda_permission" "_" {
  statement_id = "AllowExecutionFromAPIGateway"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function._.function_name
  principal = "apigateway.amazonaws.com"
  source_arn = "${aws_api_gateway_rest_api._.execution_arn}/*/*"
}

resource "aws_api_gateway_resource" "user_api" {
  rest_api_id = aws_api_gateway_rest_api._.id
  parent_id = aws_api_gateway_rest_api._.root_resource_id
  path_part = "users"
}

resource "aws_api_gateway_resource" "user_api_id" {
  rest_api_id = aws_api_gateway_rest_api._.id
  parent_id = aws_api_gateway_resource.user_api.id
  path_part = "{id}"
}

// GET /users/{id}
resource "aws_api_gateway_method" "get_user" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api_id.id
  http_method = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "get_user" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api_id.id
  http_method = aws_api_gateway_method.get_user.http_method
  type = "AWS_PROXY"
  integration_http_method = "POST"
  uri = aws_lambda_function._.invoke_arn
}

// DELETE /users/{id}
resource "aws_api_gateway_method" "delete_user" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api_id.id
  http_method = "DELETE"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "delete_user" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api_id.id
  http_method = aws_api_gateway_method.delete_user.http_method
  type = "AWS_PROXY"
  integration_http_method = "POST"
  uri = aws_lambda_function._.invoke_arn
}

// GET /users
resource "aws_api_gateway_method" "get_users" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api.id
  http_method = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "get_users" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api.id
  http_method = aws_api_gateway_method.get_users.http_method
  type = "AWS_PROXY"
  integration_http_method = "POST"
  uri = aws_lambda_function._.invoke_arn
}

// POST /users
resource "aws_api_gateway_method" "post_users" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api.id
  http_method = "POST"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "post_user" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource.user_api.id
  http_method = aws_api_gateway_method.post_users.http_method
  type = "AWS_PROXY"
  integration_http_method = "POST"
  uri = aws_lambda_function._.invoke_arn
}
