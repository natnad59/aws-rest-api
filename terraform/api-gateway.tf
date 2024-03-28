resource "aws_api_gateway_rest_api" "_" {
  name = "rest-api-example"
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

resource "aws_api_gateway_rest_api_policy" "_" {
  rest_api_id = aws_api_gateway_rest_api._.id
  policy = jsonencode(
    {
      Version = "2012-10-17"
      Statement = [
        {
          Effect    = "Allow"
          Principal = "*"
          Action    = "execute-api:Invoke"
          Resource  = "${aws_api_gateway_rest_api._.execution_arn}"
        }
      ]
    }
  )
}

resource "aws_api_gateway_resource" "_" {
  rest_api_id = aws_api_gateway_rest_api._.id
  parent_id   = aws_api_gateway_rest_api._.root_resource_id
  path_part   = "default"
}

resource "aws_api_gateway_method" "_" {
  rest_api_id   = aws_api_gateway_rest_api._.id
  resource_id   = aws_api_gateway_resource._.id
  http_method   = "OPTIONS"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "_" {
  rest_api_id = aws_api_gateway_rest_api._.id
  resource_id = aws_api_gateway_resource._.id
  http_method = aws_api_gateway_method._.http_method
  type        = "MOCK"
}

resource "aws_api_gateway_deployment" "_" {
  rest_api_id = aws_api_gateway_rest_api._.id
  depends_on  = [aws_api_gateway_integration._]
  lifecycle {
    ignore_changes = [id]
  }
}

resource "aws_api_gateway_stage" "_" {
  rest_api_id   = aws_api_gateway_rest_api._.id
  deployment_id = aws_api_gateway_deployment._.id
  stage_name    = "v1"
  lifecycle {
    ignore_changes = [id]
  }
}

resource "aws_ssm_parameter" "api_gateway_id" {
  name  = "/rest-api-example/api-gateway/id"
  value = aws_api_gateway_rest_api._.id
  type  = "SecureString"
}
