variable "project-name" {
  type = string
  description = "The name of the project"
  default = "starknet-proposals-whisperer"
}

variable "iam-role-name" {
  type = string
  description = "The name of the Lambda IAM role"
  default = "StarknetProposalWhispererLambdaRole"
}

variable "iam-policy-name" {
  type = string
  description = "The name of the Lambda IAM policy"
  default = "StarknetProposalWhispererLambdaPolicy"
}

variable "ecr-repository-name" {
  type = string
  description = "The name of the ECR repository"
}