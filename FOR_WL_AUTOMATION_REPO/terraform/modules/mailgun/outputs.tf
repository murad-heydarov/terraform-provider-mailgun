# terraform/modules/mailgun/outputs.tf

output "mail_domain" {
  description = "Mailgun registered domain"
  value       = mailgun_domain.this.name
}

output "smtp_login" {
  description = "SMTP login email (from mailgun_domain)"
  value       = mailgun_domain.this.smtp_login
}

output "smtp_user_email" {
  description = "SMTP user credential email (full)"
  value       = "${var.smtp_login_localpart}@${var.mail_domain}"
}

output "smtp_password" {
  description = "SMTP password (generated)"
  value       = random_password.smtp_password.result
  sensitive   = true
}

output "region" {
  description = "Mailgun region"
  value       = local.region
}

output "domain_verification_status" {
  description = "Domain verification status from provider"
  value       = mailgun_domain.this.verification_status
}

output "domain_verification_records" {
  description = "Domain verification records"
  value = {
    receiving_records = mailgun_domain.this.receiving_records_set
    sending_records   = mailgun_domain.this.sending_records_set
  }
}

output "dns_records_created" {
  description = "DNS records created in Cloudflare"
  value = {
    sending   = [for r in cloudflare_dns_record.mailgun_sending : "${r.type} ${r.name}"]
    receiving = [for r in cloudflare_dns_record.mailgun_receiving : "${r.type} ${r.name}"]
  }
}

output "verification_triggered" {
  description = "Whether verification was triggered"
  value       = var.enable_verification_trigger
}
