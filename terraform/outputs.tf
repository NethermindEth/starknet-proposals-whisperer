output "ecr_repository_url" {
  value = aws_ecr_repository.starknet-proposals-whisperer-repo.repository_url
}

output "lambda_function_arn" {
  value = aws_lambda_function.starknet-proposals-whisperer.arn
}