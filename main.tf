variable "region" {
  default = "eu-central-1"
}

provider "aws" {
  region = var.region
}

data "aws_caller_identity" "current" {}

resource "aws_dynamodb_table" "MessagesTable" {
  name = "observerBotMessages"
  billing_mode = "PAY_PER_REQUEST"
  hash_key = "authorId"
  range_key = "messageId"

  attribute {
    name = "authorId"
    type = "S"
  }

  attribute {
    name = "messageId"
    type = "S"
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    "DDBTableGroupKey-observerBot" = "observerBot"
  }
}

resource "aws_dynamodb_table" "MemberAdd" {
  name = "observerMemberAdd"
  billing_mode = "PAY_PER_REQUEST"
  hash_key = "userId"
  range_key = "timestamp"

  attribute {
    name = "userId"
    type = "S"
  }

  attribute {
    name = "timestamp"
    type = "N"
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    "DDBTableGroupKey-observerBot" = "observerBot"
  }
}

resource "aws_dynamodb_table" "AttachmentTable" {
  name = "observerBotAttachments"
  billing_mode = "PAY_PER_REQUEST"
  hash_key = "authorId"
  range_key = "messageId"

  attribute {
    name = "authorId"
    type = "S"
  }

  attribute {
    name = "messageId"
    type = "S"
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    "DDBTableGroupKey-observerBot" = "observerBot"
  }
}

resource "aws_dynamodb_table" "VoiceEventsTable" {
  name = "observerBotVoiceEvents"
  billing_mode = "PAY_PER_REQUEST"
  hash_key = "userId"
  range_key = "timestamp"

  attribute {
    name = "userId"
    type = "S"
  }

  attribute {
    name = "timestamp"
    type = "N"
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    "DDBTableGroupKey-observerBot" = "observerBot"
  }
}

resource "aws_dynamodb_table" "BannedMembers" {
  name = "observerBannedMembers"
  billing_mode = "PAY_PER_REQUEST"
  hash_key = "userId"
  range_key = "timestamp"

  attribute {
    name = "userId"
    type = "S"
  }

  attribute {
    name = "timestamp"
    type = "N"
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = {
    "DDBTableGroupKey-observerBot" = "observerBot"
  }
}

resource "aws_s3_bucket" "AttachmentBucket" {
  bucket = "observerbotattachments"
  acl    = "private"
}

resource "aws_lambda_function" "AttachmentDownloader" {
  function_name = "observerBotAttachments"

  filename = "attachmentLambda/attachmentLambda.zip"
  source_code_hash = filebase64sha256("attachmentLambda/attachmentLambda.zip")

  handler = "contentLambda"
  role = "arn:aws:iam::043039367084:role/service-role/discordSaveAttachments-role-paknnvaq"

  runtime = "go1.x"
  timeout = 1
  memory_size = 512


  environment {
    variables = {
      REGION = "eu-central-1",
      BUCKET = aws_s3_bucket.AttachmentBucket.bucket
      TABLE = aws_dynamodb_table.AttachmentTable.name
    }
  }
}

