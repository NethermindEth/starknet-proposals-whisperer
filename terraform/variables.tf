variable "project_name" {
  type = string
  description = "The name of the project"
  default = "starknet-proposals-whisperer"
}

variable "iam_role_name" {
  type = string
  description = "The name of the IAM role"
  default = "StarknetProposalWhispererLambdaRole"
}