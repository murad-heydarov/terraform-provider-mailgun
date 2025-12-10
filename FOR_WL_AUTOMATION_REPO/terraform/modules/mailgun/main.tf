# terraform/modules/mailgun/main.tf
# Enhanced Mailgun Module for WL Automation
# Uses: murad-heydarov/mailgun provider v0.9.0

locals {
  region           = "eu"
  mailgun_api_base = "https://api.eu.mailgun.net"
}

# ============================================================================
# Random SMTP Password
# ============================================================================

resource "random_password" "smtp_password" {
  length  = 32
  special = true
  upper   = true
  lower   = true
  numeric = true

  keepers = {
    domain = var.mail_domain
  }
}

# ============================================================================
# Mailgun Domain (Enhanced Provider)
# ============================================================================

resource "mailgun_domain" "this" {
  name   = var.mail_domain
  region = local.region

  # ✅ ENHANCED FEATURES (Provider v0.9.0)
  use_automatic_sender_security = true
  trigger_verification          = false  # Trigger after DNS propagation

  # SMTP Settings
  spam_action   = "disabled"
  smtp_password = random_password.smtp_password.result

  # DKIM
  force_dkim_authority = true

  # Tracking
  web_scheme = "https"

  # Domain Settings
  wildcard = false

  lifecycle {
    create_before_destroy = true
  }
}

# ============================================================================
# SMTP Credentials
# ============================================================================

resource "mailgun_domain_credential" "smtp_user" {
  domain   = mailgun_domain.this.name
  login    = var.smtp_login_localpart
  password = random_password.smtp_password.result
  region   = local.region

  lifecycle {
    ignore_changes = [password]
  }

  depends_on = [mailgun_domain.this]
}

# ============================================================================
# DNS Records - Sending (SPF, DKIM, Tracking)
# ============================================================================

locals {
  sending_records = {
    for record in mailgun_domain.this.sending_records_set :
    "${record.record_type}-${record.name}" => record
  }

  receiving_records = {
    for record in mailgun_domain.this.receiving_records_set :
    "${record.record_type}-${record.value}" => record
  }
}

resource "cloudflare_dns_record" "mailgun_sending" {
  for_each = local.sending_records

  zone_id = var.cloudflare_zone_id
  name    = each.value.name
  type    = each.value.record_type
  ttl     = 1
  proxied = false

  content = each.value.record_type == "CNAME" ? trimsuffix(each.value.value, ".") : each.value.value

  comment = "Mailgun ${each.value.record_type} - ${var.mail_domain}"

  depends_on = [mailgun_domain.this]
}

# ============================================================================
# DNS Records - Receiving (MX)
# ============================================================================

resource "cloudflare_dns_record" "mailgun_receiving" {
  for_each = local.receiving_records

  zone_id = var.cloudflare_zone_id
  name    = var.mail_domain
  type    = each.value.record_type
  ttl     = 1
  proxied = false

  content  = each.value.value
  priority = each.value.record_type == "MX" ? each.value.priority : null

  comment = "Mailgun ${each.value.record_type} - ${var.mail_domain}"

  depends_on = [mailgun_domain.this]
}

# ============================================================================
# Wait for DNS Propagation
# ============================================================================

resource "time_sleep" "wait_for_dns" {
  depends_on = [
    cloudflare_dns_record.mailgun_sending,
    cloudflare_dns_record.mailgun_receiving
  ]

  create_duration = "${var.dns_wait_seconds}s"
}

# ============================================================================
# Trigger Domain Verification
# ============================================================================

resource "null_resource" "trigger_verification" {
  count = var.enable_verification_trigger ? 1 : 0

  triggers = {
    domain           = mailgun_domain.this.name
    dns_records_hash = sha256(jsonencode([
      values(cloudflare_dns_record.mailgun_sending)[*].id,
      values(cloudflare_dns_record.mailgun_receiving)[*].id
    ]))
    always_run = timestamp()
  }

  provisioner "local-exec" {
    command = "curl -X PUT ${local.mailgun_api_base}/v3/domains/${var.mail_domain}/verify -u api:${var.mailgun_api_key} --fail-with-body"
  }

  depends_on = [time_sleep.wait_for_dns]
}
