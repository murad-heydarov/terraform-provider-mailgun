# WL Automation - Mailgun Enhanced Provider Deployment Guide

**Salam!** Bu guide sizin üçün step-by-step izahatdır ki, enhanced Mailgun provider-i necə deploy edəcəksiniz və istifadə edəcəksiniz.

## 📋 İçindəkilər

1. [Provider Setup](#1-provider-setup)
2. [GitHub Repository](#2-github-repository)
3. [Module Integration](#3-module-integration)
4. [Complete Example](#4-complete-example)
5. [Testing](#5-testing)
6. [Troubleshooting](#6-troubleshooting)

---

## 1. Provider Setup

### 1.1. Provider Kodu Hazırdır

Enhanced provider artıq hazırdır:
- **Yerləşmə**: `/workspace/terraform-provider-mailgun/`
- **Yeni Features**:
  - ✅ `use_automatic_sender_security` field
  - ✅ `trigger_verification` field
  - ✅ `verification_status` computed field

### 1.2. Provider Build

```bash
cd /workspace/terraform-provider-mailgun

# Build provider
make build

# Test locally (optional)
make test

# Install locally for testing
make install
```

---

## 2. GitHub Repository

### 2.1. GitHub Repo Yaradın

```bash
# 1. GitHub-da yeni repo yaradın: terraform-provider-mailgun
# URL: https://github.com/murad-heydarov/terraform-provider-mailgun

# 2. Provider kodunu push edin
cd /workspace/terraform-provider-mailgun

git init
git add .
git commit -m "feat: Add use_automatic_sender_security and trigger_verification features

- Add use_automatic_sender_security field for Mailgun API v4 support
- Add trigger_verification to automatically verify domains
- Add verification_status computed field
- Implement HTTP client for API v4 features
- Update documentation with WL automation examples"

git remote add origin https://github.com/murad-heydarov/terraform-provider-mailgun.git
git branch -M main
git push -u origin main
```

### 2.2. Tag və Release Yaradın

```bash
# Create initial release
git tag -a v0.9.0 -m "Initial enhanced release with WL automation features"
git push origin v0.9.0
```

GitHub Actions avtomatik olaraq provider-i build edəcək və release yaradacaq.

### 2.3. (Optional) Terraform Registry

Əgər provider-i Terraform Registry-də public etmək istəyirsinizsə:

1. [Terraform Registry](https://registry.terraform.io/publish) → Sign in with GitHub
2. GitHub repo-nu connect edin
3. Terraform avtomatik publish edəcək

---

## 3. Module Integration

### 3.1. Köhnə Module vs Yeni Module

#### ❌ Köhnə Module (`terraform/modules/mailgun/main.tf`)

**Problemlər:**
- `null_resource` + `curl` istifadə edir
- Manual API calls
- Çox kompleks kod
- Hard to maintain

#### ✅ Yeni Enhanced Module

**Location**: `/workspace/terraform-modules-mailgun-enhanced/`

**Üstünlüklər:**
- Native provider features
- Clean code
- Automatic verification
- Better state management

### 3.2. Module-u Proyektinizə Əlavə Edin

Sizin WL automation struktuuna əlavə edin:

```bash
# Köhnə mailgun module-u silin
rm -rf /workspace/terraform/modules/mailgun

# Yeni enhanced module-u copy edin
cp -r /workspace/terraform-modules-mailgun-enhanced /workspace/terraform/modules/mailgun
```

### 3.3. Provider Konfiqurasiyası

Update your `terraform/environments/prod/providers.tf`:

```hcl
provider "mailgun" {
  api_key = var.mailgun_api_key
}
```

Update `terraform/environments/prod/versions.tf`:

```hcl
terraform {
  required_version = ">= 1.5.0"

  required_providers {
    # ... existing providers ...
    
    mailgun = {
      source  = "murad-heydarov/mailgun"  # ✅ Changed
      version = "~> 0.9.0"                # ✅ Changed
    }
  }
}
```

---

## 4. Complete Example

### 4.1. Tam WL Deployment Example

**File**: `terraform/environments/prod/wl-configs/afftech.auto.tfvars`

```hcl
# ============================================================================
# AGENT WHITE LABEL CONFIGURATION
# ============================================================================

domain              = "afftech.xyz"
wl_type             = "agent"
platform_code       = "AFFTECH"
cloudflare_zone_id  = "84787ea66aa226406e7c736892c6d493"

# ============================================================================
# Subdomain Configuration
# ============================================================================

admin_subdomain = "admin"
agent_subdomain = "agent"

# ============================================================================
# Mailgun Configuration
# ============================================================================

mail_domain              = "support.afftech.xyz"
smtp_login_localpart     = "admin"
mailgun_dns_wait_seconds = 120

# ============================================================================
# ALB Configuration
# ============================================================================

alb_dns_name = "mt-apps-ingress-978b1006d8a9d559.elb.eu-central-1.amazonaws.com"

# ============================================================================
# GitLab CI/CD Configuration
# ============================================================================

gitlab_project_id = "marketingtech/pmaffiliate/pmaffiliate-react-front"
```

### 4.2. Deploy Edin

```bash
cd /workspace/terraform/environments/prod

# Initialize (download new provider)
terraform init -upgrade

# Plan
terraform plan -var-file="wl-configs/afftech.auto.tfvars"

# Apply
terraform apply -var-file="wl-configs/afftech.auto.tfvars" -auto-approve
```

### 4.3. Outputs

Deploy-dan sonra outputs:

```bash
terraform output -json

# Mailgun credentials
terraform output -json mailgun_smtp_credentials
```

Output nümunəsi:

```json
{
  "smtp_login": "postmaster@support.afftech.xyz",
  "smtp_password": "generated-secure-password-32-chars",
  "mail_domain": "support.afftech.xyz",
  "region": "eu",
  "verification_status": "verification_triggered"
}
```

---

## 5. Testing

### 5.1. DNS Records Yoxlayın

```bash
# SPF record
dig TXT support.afftech.xyz +short

# MX records
dig MX support.afftech.xyz +short

# DKIM records
dig CNAME pdk1._domainkey.support.afftech.xyz +short
```

### 5.2. Mailgun Console

1. [Mailgun EU Console](https://app.eu.mailgun.com/mg/sending/domains)
2. Select `support.afftech.xyz`
3. "Domain verification & DNS" tab-da bütün records green olmalıdır

### 5.3. Test Email Göndərin

```bash
curl -s --user 'api:YOUR_API_KEY' \
  https://api.eu.mailgun.net/v3/support.afftech.xyz/messages \
  -F from='Test <test@support.afftech.xyz>' \
  -F to='your-email@example.com' \
  -F subject='Test Email' \
  -F text='Testing Mailgun configuration'
```

---

## 6. Troubleshooting

### 6.1. Provider Not Found

**Error**: `Could not find provider murad-heydarov/mailgun`

**Həll**:
- GitHub-da release yaratdığınızdan əmin olun
- Tag push etmisiniz: `git push origin v0.9.0`
- Terraform Registry-də publish olunmasını gözləyin (5-10 dəq)

**Temporary Həll** (Local testing):

```bash
# Build və local install
cd /workspace/terraform-provider-mailgun
make install

# terraform-də local provider istifadə edin
# terraform/environments/prod/versions.tf
terraform {
  required_providers {
    mailgun = {
      source  = "local/murad-heydarov/mailgun"  # Local path
      version = "~> 0.9.0"
    }
  }
}
```

### 6.2. Verification Fails

**Səbəb**: DNS records hələ propagate olmayıb

**Həll**:

```bash
# 1. Wait artırın
mailgun_dns_wait_seconds = 180  # 3 minutes

# 2. Manual verify edin
curl -X PUT \
  "https://api.eu.mailgun.net/v3/domains/support.afftech.xyz/verify" \
  -u "api:YOUR_API_KEY"
```

### 6.3. API v4 Features Don't Work

**Səbəb**: Provider updated deyil

**Həll**:

```bash
# Provider-i yenidən build edin
cd /workspace/terraform-provider-mailgun
make clean
make build
make install

# Terraform-də yeniləyin
cd /workspace/terraform/environments/prod
terraform init -upgrade
```

---

## 7. Yekun Struktur

Sizin final struktur:

```
/workspace/
├── terraform-provider-mailgun/          # ✅ Enhanced provider
│   ├── mailgun/
│   │   ├── resource_mailgun_domain.go   # ✅ Enhanced with new fields
│   │   └── ...
│   ├── go.mod
│   ├── Makefile
│   ├── README.md                        # ✅ Complete documentation
│   └── .github/workflows/               # ✅ CI/CD
│
└── terraform/
    ├── modules/
    │   ├── mailgun/                     # ✅ Enhanced module (simplified)
    │   │   ├── main.tf
    │   │   ├── variables.tf
    │   │   ├── outputs.tf
    │   │   ├── versions.tf
    │   │   └── README.md
    │   ├── acm-certificates/
    │   ├── cloudfront-s3-website/
    │   └── wl-domain/
    │
    └── environments/
        └── prod/
            ├── main.tf
            ├── providers.tf             # ✅ Updated
            ├── versions.tf              # ✅ Updated
            ├── variables.tf
            ├── outputs.tf
            └── wl-configs/
                ├── afftech.auto.tfvars
                ├── brandx.auto.tfvars
                └── owinbet.auto.tfvars
```

---

## 8. Növbəti Addımlar

### 8.1. Provider GitHub-a Push

```bash
cd /workspace/terraform-provider-mailgun
git init
git add .
git commit -m "feat: Initial enhanced provider with WL automation"
git remote add origin https://github.com/murad-heydarov/terraform-provider-mailgun.git
git push -u origin main
git tag v0.9.0
git push origin v0.9.0
```

### 8.2. Module Test

```bash
cd /workspace/terraform/environments/prod

# Test single WL
terraform plan -var-file="wl-configs/afftech.auto.tfvars"

# Apply
terraform apply -var-file="wl-configs/afftech.auto.tfvars"
```

### 8.3. Documentation Update

Öz confluence page-nizi update edin:
- Provider link əlavə edin
- New automated process izah edin
- Old manual steps remove edin

---

## 9. Fərqlər (Old vs New)

| Feature | Old (Manual) | New (Automated) |
|---------|-------------|-----------------|
| Domain creation | Manual Mailgun UI | ✅ Terraform |
| SMTP user | Manual creation | ✅ Auto-generated |
| DNS records | Manual Cloudflare | ✅ Automatic |
| Verification | Manual button click | ✅ Automatic |
| Sender security | Not available | ✅ API v4 enabled |
| Time required | ~15 minutes | ✅ ~3 minutes |

---

## 10. FAQ

**Q: Provider Terraform Registry-də görünmür?**  
A: GitHub release yaradın və 10 dəqiqə gözləyin. Və ya local provider istifadə edin.

**Q: Verification niyə fail olur?**  
A: DNS propagation gözləyin. `mailgun_dns_wait_seconds` artırın.

**Q: use_automatic_sender_security nə işə yarayır?**  
A: Mailgun API v4 feature-dir. Domain security-ni avtomatik aktivləşdirir.

**Q: Köhnə wgebis provider ilə işləyir?**  
A: Xeyr, enhanced provider lazımdır. Amma migration asandır.

**Q: Production-da təhlükəsiz?**  
A: Bəli! Base wgebis provider-in üzərində build edilib, yalnız 3 field əlavə edilib.

---

## 🎉 Uğurlar!

Artıq tam avtomatik WL Mailgun deployment-niz hazırdır!

**Suallar?** Contact DevOps team.
