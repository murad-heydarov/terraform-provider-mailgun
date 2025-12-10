---
page_title: "Mailgun: mailgun_domain_verification"
---

# mailgun\_domain\_verification

Triggers Mailgun’s domain verification workflow via the public API. Use this resource after DNS records are created (for example, via the Cloudflare provider) to replace the manual “Verify DNS settings” button in the Mailgun UI.

## Example Usage

```hcl
resource "mailgun_domain" "example" {
  name        = "support.example.com"
  region      = "eu"
  spam_action = "disabled"
}

# Cloudflare records omitted for brevity

resource "mailgun_domain_verification" "example" {
  domain          = mailgun_domain.example.name
  region          = "eu"
  wait_for_active = true
  poll_interval   = "20s"
  timeout         = "15m"
}
```

## Argument Reference

The following arguments are supported:

* `domain` – (Required) Domain name that should be verified. This usually references the `mailgun_domain` resource.
* `region` – (Optional) Mailgun region (`us` or `eu`). Defaults to `us`.
* `wait_for_active` – (Optional) When `true` (default), Terraform will poll Mailgun until all DNS records are reported as valid.
* `poll_interval` – (Optional) Interval between verification status checks while waiting. Accepts Go duration strings such as `"15s"`. Default: `15s`.
* `timeout` – (Optional) Maximum amount of time to wait for Mailgun to report success. Accepts Go duration strings such as `"10m"`. Default: `10m`.

All arguments are `ForceNew`—changing any of them will re-trigger verification.

## Attributes Reference

The following attributes are exported:

* `status` – Current Mailgun state for the domain (for example, `active` or `unverified`).
* `receiving_records` – Snapshot of Mailgun’s receiving DNS records and their validation state.
* `sending_records` – Snapshot of Mailgun’s sending DNS records and their validation state.
  * Each record includes `name`, `value`, `record_type`, `priority`, `valid`, `is_active`, and the `cached` values Mailgun sees.

## Import

This resource represents an action rather than a persisted object, so it cannot be imported.
