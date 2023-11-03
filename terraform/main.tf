resource "aws_ecr_repository" "starknet-proposals-whisperer-repo" {
  name = var.project_name 
}

data "aws_iam_policy_document" "lambda_policy" {
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents"
    ]
    resources = ["*"]
  }
}

resource "aws_iam_role" "lambda_role" {
  name               = var.iam_role_name
  assume_role_policy = data.aws_iam_policy_document.lambda_policy.json
}

resource "aws_lambda_function" "starknet-proposals-whisperer" {
  function_name = var.project_name
  role          = aws_iam_role.lambda_role.arn
  image_uri     = "${aws_ecr_repository.starknet-proposals-whisperer-repo.repository_url}:latest"

  package_type = "Image"
}