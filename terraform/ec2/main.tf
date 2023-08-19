terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.5.0"
}

provider "aws" {
  region  = "ap-southeast-2"
}

# Create EC2 instance
resource "aws_instance" "app_server" {
  ami           = "ami-0a709bebf4fa9246f"
  instance_type = "t2.micro"

  tags = {
    Name = "ExampleAppServerInstance"
  }
}
