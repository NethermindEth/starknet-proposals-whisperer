module "starknet-proposals-whisperer" {
  source        = "./modules"
  s3_name       = "starknet-proposals-whisperer-s3"
  dynamodb_name = "starknet-proposals-whisperer-dynamodb"
  region        = "us-east-2"
}