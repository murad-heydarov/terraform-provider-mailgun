# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.9.0] - 2024-12-10

### Added

- 🚀 **`use_automatic_sender_security`** field for enabling automatic sender security via Mailgun API v4
- 🚀 **`trigger_verification`** field to automatically verify domain after creation
- 🚀 **`verification_status`** computed field to track domain verification status
- ✨ HTTP client-based API calls for v4 features not available in mailgun-go SDK
- 📚 Enhanced documentation with WL automation examples
- 🔧 GitHub Actions workflows for automated testing and releases
- 🔧 GoReleaser configuration for multi-platform builds

### Changed

- 📝 Updated README with detailed usage examples
- 🔄 Modified domain create/update logic to support new fields
- ⚡ Improved error handling for API v4 calls

### Fixed

- 🐛 API v4 endpoint compatibility for EU region
- 🐛 Proper context handling in verification calls

## [0.8.1] - 2023-XX-XX (wgebis/terraform-provider-mailgun)

### Base Version

This is the fork point from the original [wgebis/terraform-provider-mailgun](https://github.com/wgebis/terraform-provider-mailgun).

All previous changes and features are maintained from the original provider.

---

**Note:** Version 0.9.0 and above are maintained by this enhanced fork.
For versions 0.8.1 and below, see the [original provider changelog](https://github.com/wgebis/terraform-provider-mailgun).
