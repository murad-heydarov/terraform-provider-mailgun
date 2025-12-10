# Mailgun Terraform Provider

This repository hosts `murad-heydarov/terraform-mailgun-provider`, a fork of the official Mailgun provider tailored for WL automation projects. It keeps feature-parity with the upstream provider and adds:

- `mailgun_domain_verification` – fully automated DNS verification with optional polling.
- First-class support for `use_automatic_sender_security` on domains.
- End-to-end documentation aimed at WL automation workflows.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) **1.6+**
- [Go](https://go.dev/doc/install) **1.21+** (only required when building from source)
- A Mailgun API key (`MAILGUN_API_KEY`)

## Installation

Until the provider is published on the Terraform Registry, build it locally and point Terraform to the resulting binary.

```sh
git clone https://github.com/murad-heydarov/terraform-mailgun-provider.git
cd terraform-mailgun-provider
make build
```

Terraform will place the compiled plugin inside `~/.terraform.d/plugins` (see `GNUmakefile` for the exact target path). Within your configuration:

```hcl
terraform {
  required_providers {
    mailgun = {
      source  = "murad-heydarov/mailgun"
      version = "0.1.0"
    }
  }
}

provider "mailgun" {
  api_key = var.mailgun_api_key
}
```

## Development

```sh
# Run unit tests
make test

# Run acceptance tests (requires a real Mailgun account and DNS control)
TF_ACC=1 MAILGUN_API_KEY=... make testacc
```

Acceptance tests create real domains, DNS records and credentials—clean up any leftovers in your Mailgun organization after running them.
