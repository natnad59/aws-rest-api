terraform {
  backend "s3" {
    bucket = "backend-dev-tfstate"
    region = "us-east-1"
    key    = "rest-api-example.tfstate"
  }
}

provider "aws" {
  region = "us-east-1"
}
