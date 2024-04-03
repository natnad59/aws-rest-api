data "aws_partition" "current" {}

data aws_ssm_parameter "artifacts_bucket" {
    name = "/rest-api-example/artifacts-bucket"
}
