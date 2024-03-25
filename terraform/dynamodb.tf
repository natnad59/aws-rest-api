resource "aws_dynamodb_table" "_" {
  name         = "rest-api-example-${terraform.workspace}"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  global_secondary_index {
    name            = "name-index"
    hash_key        = "name"
    projection_type = "ALL"
  }

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "name"
    type = "S"
  }
}
