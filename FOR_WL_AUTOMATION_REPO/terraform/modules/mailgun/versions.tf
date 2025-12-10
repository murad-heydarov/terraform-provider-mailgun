# terraform/modules/mailgun/versions.tf

terraform {
  required_version = ">= 1.5.0"

  required_providers {
    mailgun = {
      source  = "murad-heydarov/mailgun"
      version = "~> 0.9.0"
    }
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 5.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6"
    }
    time = {
      source  = "hashicorp/time"
      version = "~> 0.12"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }
}
