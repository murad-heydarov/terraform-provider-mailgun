# Terraform Mailgun Module (Enhanced)

> Uses [murad-heydarov/terraform-provider-mailgun](https://github.com/murad-heydarov/terraform-provider-mailgun) with automatic verification features.

## Features

✅ **Automatic Sender Security** - Enabled by default via Mailgun API v4  
✅ **Automatic Domain Verification** - Trigger verification after DNS propagation  
✅ **SMTP Credentials** - Auto-generated secure password  
✅ **DNS Automation** - Automatically creates all required Cloudflare records  
✅ **EU Region Support** - Defaults to EU Mailgun region  

## Usage

### Basic Example

```hcl
module "mailgun" {
  source = "./modules/mailgun"

  mail_domain        = "support.afftech.xyz"
  cloudflare_zone_id = "84787ea66aa226406e7c736892c6d493"
  mailgun_api_key    = var.mailgun_api_key
}

output "smtp_credentials" {
  value = {
    login    = module.mailgun.smtp_login
    password = module.mailgun.smtp_password
  }
  sensitive = true
}
```

### Advanced Example with Custom Settings

```hcl
module "mailgun_wl" {
  source = "./modules/mailgun"

  mail_domain              = "support.${var.domain}"
  cloudflare_zone_id       = var.cloudflare_zone_id
  mailgun_api_key          = var.mailgun_api_key
  smtp_login_localpart     = "noreply"
  dns_wait_seconds         = 180
  enable_verification_trigger = true

  tags = {
    Environment = "production"
    Project     = "WL-Automation"
    Domain      = var.domain
  }
}
```

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.5.0 |
| mailgun | ~> 0.9.0 |
| cloudflare | ~> 5.0 |
| random | ~> 3.6 |
| time | ~> 0.12 |
| null | ~> 3.2 |

## Providers

| Name | Version |
|------|---------|
| mailgun | ~> 0.9.0 |
| cloudflare | ~> 5.0 |
| random | ~> 3.6 |
| time | ~> 0.12 |
| null | ~> 3.2 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| mail_domain | Mail domain (e.g., 'support.afftech.xyz') | `string` | n/a | yes |
| cloudflare_zone_id | Cloudflare Zone ID for DNS records | `string` | n/a | yes |
| mailgun_api_key | Mailgun API key (sensitive) | `string` | n/a | yes |
| smtp_login_localpart | SMTP user local part | `string` | `"admin"` | no |
| dns_wait_seconds | Seconds to wait after DNS creation before verification | `number` | `120` | no |
| enable_verification_trigger | Enable automatic domain verification | `bool` | `true` | no |
| tags | Resource tags | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| mail_domain | Mailgun registered domain |
| smtp_login | SMTP login email |
| smtp_user_email | Full SMTP user credential email |
| smtp_password | Generated SMTP password (sensitive) |
| region | Mailgun region (eu) |
| domain_verification_status | Domain verification status from provider |
| domain_verification_records | Domain verification records |
| dns_records_created | DNS records created in Cloudflare |
| verification_triggered | Whether verification was triggered |

## How It Works

1. **Domain Creation** - Creates Mailgun domain with automatic sender security enabled
2. **SMTP User** - Generates secure random password and creates SMTP credential
3. **DNS Records** - Automatically creates all required DNS records in Cloudflare:
   - SPF (TXT record)
   - MX records (2x)
   - DKIM (CNAME records)
   - Email tracking (CNAME)
4. **DNS Propagation** - Waits for specified time (default 120s)
5. **Verification** - Triggers domain verification via Mailgun API

## DNS Records Created

The module automatically creates these records in Cloudflare:

### Sending Records (SPF, DKIM, Tracking)
- **TXT** - SPF record (`v=spf1 include:mailgun.org ~all`)
- **CNAME** - DKIM keys (pdk1, pdk2)
- **CNAME** - Email tracking

### Receiving Records (MX)
- **MX** - `mxa.eu.mailgun.org` (priority 10)
- **MX** - `mxb.eu.mailgun.org` (priority 10)

## Enhanced vs Original Module

### ❌ Original Module Issues

- Used `null_resource` with `curl` for verification
- Manual API calls via shell scripts
- No automatic sender security
- Complex implementation

### ✅ Enhanced Module Benefits

- Native provider support for verification
- Cleaner, more maintainable code
- Automatic sender security enabled
- Better error handling
- Proper Terraform state management

## Migration from Original Module

If you're using the original wgebis provider, migration is simple:

1. Update provider source:

```hcl
terraform {
  required_providers {
    mailgun = {
      source  = "murad-heydarov/mailgun"  # Changed
      version = "~> 0.9.0"                # Changed
    }
  }
}
```

2. No changes needed to module code - it's backward compatible!

3. Run terraform init to download new provider:

```bash
terraform init -upgrade
```

## Troubleshooting

### Verification Fails

If verification fails, you can manually trigger it:

```bash
curl -X PUT \
  "https://api.eu.mailgun.net/v3/domains/YOUR_DOMAIN/verify" \
  -u "api:YOUR_API_KEY"
```

### DNS Records Not Resolving

Check DNS propagation:

```bash
dig TXT support.yourdomain.com
dig MX support.yourdomain.com
```

Wait 5-10 minutes for global propagation.

### SMTP Authentication Fails

Verify credentials in Mailgun console:
1. Go to [Mailgun EU](https://app.eu.mailgun.com/)
2. Select your domain
3. Check "Sending" → "Domain verification & DNS"
4. Verify all records show green checkmarks

## Examples

See the [examples](../../examples/mailgun/) directory for complete working examples.

## License

Same as parent Terraform configuration.

## Support

For issues related to:
- **Provider** - [GitHub Issues](https://github.com/murad-heydarov/terraform-provider-mailgun/issues)
- **Module** - Contact your DevOps team
- **Mailgun API** - [Mailgun Support](https://help.mailgun.com/)
