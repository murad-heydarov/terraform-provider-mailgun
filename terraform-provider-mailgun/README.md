# Terraform Mailgun Provider (Enhanced)

> **Fork of [wgebis/terraform-provider-mailgun](https://github.com/wgebis/terraform-provider-mailgun)** with additional features for enterprise WL automation.

[![Go Version](https://img.shields.io/github/go-mod/go-version/murad-heydarov/terraform-provider-mailgun)](https://go.dev/)
[![License](https://img.shields.io/github/license/murad-heydarov/terraform-provider-mailgun)](LICENSE)

## ✨ Enhanced Features

This fork adds critical features missing from the official provider:

### 🆕 New Features

1. **`use_automatic_sender_security`** - Enable automatic sender security (Mailgun API v4)
2. **`trigger_verification`** - Automatically verify domain after DNS records are added
3. **`verification_status`** - Track domain verification status in Terraform state

### 🎯 Use Case

Perfect for **automated WL (White Label) deployments** where domains need to be:
- Created programmatically
- Verified automatically after DNS propagation
- Secured with latest Mailgun security features

## 📦 Installation

### Terraform 0.13+

Add to your `terraform` block:

```hcl
terraform {
  required_providers {
    mailgun = {
      source  = "murad-heydarov/mailgun"
      version = "~> 0.9.0"
    }
  }
}

provider "mailgun" {
  api_key = var.mailgun_api_key
}
```

## 🚀 Usage Examples

### Basic Domain with Auto-Verification

```hcl
resource "mailgun_domain" "example" {
  name   = "support.example.com"
  region = "eu"

  # 🆕 Enhanced features
  use_automatic_sender_security = true
  trigger_verification          = true

  # Standard features
  spam_action          = "disabled"
  force_dkim_authority = true
  web_scheme           = "https"
}

output "verification_status" {
  value = mailgun_domain.example.verification_status
}
```

### Complete WL Automation Example

```hcl
# 1. Create domain with enhanced security
resource "mailgun_domain" "wl_domain" {
  name   = var.mail_domain
  region = "eu"

  use_automatic_sender_security = true
  trigger_verification          = false  # Wait for DNS first

  spam_action          = "disabled"
  smtp_password        = random_password.smtp.result
  force_dkim_authority = true
  web_scheme           = "https"
}

# 2. Add DNS records to Cloudflare
resource "cloudflare_dns_record" "mailgun_spf" {
  zone_id = var.cloudflare_zone_id
  name    = var.mail_domain
  type    = "TXT"
  content = "v=spf1 include:mailgun.org ~all"
  proxied = false
}

# ... (add other DNS records)

# 3. Wait for DNS propagation
resource "time_sleep" "dns_propagation" {
  depends_on      = [cloudflare_dns_record.mailgun_spf]
  create_duration = "120s"
}

# 4. Trigger verification
resource "null_resource" "trigger_verification" {
  depends_on = [time_sleep.dns_propagation]

  triggers = {
    domain = mailgun_domain.wl_domain.name
  }

  provisioner "local-exec" {
    command = "curl -X PUT https://api.eu.mailgun.net/v3/domains/${var.mail_domain}/verify -u 'api:${var.mailgun_api_key}'"
  }
}

# 5. Create SMTP credentials
resource "mailgun_domain_credential" "smtp_user" {
  domain   = mailgun_domain.wl_domain.name
  login    = "admin"
  password = random_password.smtp.result
  region   = "eu"
}
```

## 🔧 Resource: `mailgun_domain`

### Schema

#### New Arguments

- **`use_automatic_sender_security`** - (Optional, Boolean, Default: `true`) Enable automatic sender security via Mailgun API v4.
- **`trigger_verification`** - (Optional, Boolean, Default: `false`) Trigger domain verification immediately after creation.

#### New Attributes

- **`verification_status`** - (Computed, String) Current verification status.

#### Existing Arguments

All existing arguments from `wgebis/mailgun` are supported:

- `name` - (Required) Domain name
- `region` - (Optional) `us` or `eu` (default: `us`)
- `spam_action` - (Optional) Spam action (default: `disabled`)
- `smtp_password` - (Optional, Sensitive) SMTP password
- `wildcard` - (Optional) Enable wildcard
- `force_dkim_authority` - (Optional) Force DKIM authority
- `web_scheme` - (Optional) `http` or `https`
- And more... (see [official docs](https://registry.terraform.io/providers/wgebis/mailgun/latest/docs/resources/domain))

## 🏗️ Building from Source

### Prerequisites

- [Go](https://golang.org/doc/install) 1.21+
- [Terraform](https://www.terraform.io/downloads.html) 1.5+

### Build

```bash
git clone https://github.com/murad-heydarov/terraform-provider-mailgun.git
cd terraform-provider-mailgun
go build -o terraform-provider-mailgun
```

### Local Development

```bash
# Install locally
make install

# Run tests
make test

# Run acceptance tests (requires Mailgun API key)
export MAILGUN_API_KEY="your-api-key"
make testacc
```

## 🤝 Contributing

Contributions are welcome! Please open an issue or PR.

## 📄 License

Mozilla Public License 2.0 - same as the original provider.

## 🙏 Credits

- Original provider by [wgebis](https://github.com/wgebis)
- Enhanced by [Murad Heydarov](https://github.com/murad-heydarov)

## 📚 Resources

- [Mailgun API Documentation](https://documentation.mailgun.com/en/latest/api_reference.html)
- [Original Provider](https://github.com/wgebis/terraform-provider-mailgun)
- [Terraform Registry](https://registry.terraform.io/)

---

**Note:** This is an enhanced fork specifically designed for automated WL deployments. For general Mailgun usage, consider the [official provider](https://github.com/wgebis/terraform-provider-mailgun).
