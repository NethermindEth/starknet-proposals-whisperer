terraform {
  backend "s3" {
    bucket         = "starknet-proposals-whisperer-s3"
    key            = "terraform.tfstate"
    dynamodb_table = "starknet-proposals-whisperer-dynamodb"
    region         = "us-east-2"
  }
}

provider "aws" {
  region = "us-east-2"
}
