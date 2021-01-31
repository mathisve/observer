terraform {
  backend "s3" {
    bucket = "gusbackend"
    key = "terraform"
    region = "eu-central-1"
  }
}

variable "region" {
  default = "eu-central-1"
}

provider "aws" {
  region = var.region
}

data "aws_caller_identity" "current" {}

resource "aws_dynamodb_table" "MessagesTable" {
  name = "gusBotMessages"
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
    "DDBTableGroupKey-observerBot" = "gusBot"
  }
}

resource "aws_dynamodb_table" "MemberAdd" {
  name = "gusMemberAdd"
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
    "DDBTableGroupKey-observerBot" = "gusBot"
  }
}
resource "aws_dynamodb_table" "BannedMembers" {
  name = "gusBannedMembers"
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
    "DDBTableGroupKey-observerBot" = "gusBot"
  }
}

