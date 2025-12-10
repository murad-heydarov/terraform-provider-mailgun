# terraform/modules/mailgun/variables.tf (ENHANCED VERSION)

variable "mail_domain" {
  description = "Mail domain (e.g., 'support.afftech.xyz')"
  type        = string

  validation {
    condition     = can(regex("^[a-z0-9][a-z0-9.-]*[a-z0-9]\\.[a-z]{2,}$", var.mail_domain))
    error_message = "Mail domain must be valid format (e.g., 'support.example.com')."
  }
}

variable "cloudflare_zone_id" {
  description = "Cloudflare Zone ID for DNS records"
  type        = string

  validation {
    condition     = can(regex("^[a-f0-9]{32}$", var.cloudflare_zone_id))
    error_message = "Cloudflare Zone ID must be 32 hexadecimal characters."
  }
}

variable "mailgun_api_key" {
  description = "Mailgun API key (sensitive) - required for verification"
  type        = string
  sensitive   = true
}

variable "smtp_login_localpart" {
  description = "SMTP user local part (e.g., 'admin' becomes admin@domain)"
  type        = string
  default     = "admin"

  validation {
    condition     = can(regex("^[a-z0-9._-]+$", var.smtp_login_localpart))
    error_message = "SMTP login must be valid email local part."
  }
}

variable "dns_wait_seconds" {
  description = "Seconds to wait after DNS record creation before verification"
  type        = number
  default     = 120

  validation {
    condition     = var.dns_wait_seconds >= 30 && var.dns_wait_seconds <= 300
    error_message = "DNS wait must be between 30 and 300 seconds."
  }
}

variable "enable_verification_trigger" {
  description = "Enable automatic domain verification after DNS propagation"
  type        = bool
  default     = true
}

variable "tags" {
  description = "Resource tags"
  type        = map(string)
  default     = {}
}
