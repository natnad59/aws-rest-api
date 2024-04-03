resource "aws_dynamodb_table" "_" {
  name         = "rest-api-example-${terraform.workspace}"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }
}
