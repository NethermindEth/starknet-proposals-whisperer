variable "project-name" {
  type = string
  description = "The name of the project"
  default = "starknet-proposals-whisperer"
}

locals {
  tags = {
    "Project" : "nubia-slack-bot"
  }
}

resource "aws_ecr_repository" "starknet-proposals-whisperer-repo" {
  name = var.project-name
  tags = local.tags
}

output "ecr_repository_url" {
  value = aws_ecr_repository.starknet-proposals-whisperer-repo.repository_url
}