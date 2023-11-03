locals {
  tags = {
    "Project" : "nubia-slack-bot"
  }
}

data "aws_iam_policy_document" "lambda-doc-role" {
  statement {
    actions = ["sts:AssumeRole"]
    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda-role" {
  name               = var.iam-role-name
  description        = "IAM role for starknet-proposals-whisperer lambda function"
  assume_role_policy = data.aws_iam_policy_document.lambda-doc-role.json
  tags               = local.tags
}

data "aws_iam_policy_document" "lambda-doc-policy" {
  statement {
    actions = [
      "logs:CreateLogGroup",
      "logs:CreateLogStream",
      "logs:PutLogEvents",
    ]
    resources = ["arn:aws:logs:*:*:*"]
  }
}

resource "aws_iam_policy" "lambda-policy" {
  name        = var.iam-policy-name
  description = "IAM policy for function lambda"
  policy      = data.aws_iam_policy_document.lambda-doc-policy.json

  tags = local.tags
}

resource "aws_iam_role_policy_attachment" "lambda-attachment" {
  role       = aws_iam_role.lambda-role.name
  policy_arn = aws_iam_policy.lambda-policy.arn
}

resource "aws_lambda_function" "starknet-proposals-whisperer" {
  function_name = var.project-name
  role          = aws_iam_role.lambda-role.arn
  image_uri     = "${var.ecr-repository-name}:latest"
  tags          = local.tags

  package_type = "Image"
}
