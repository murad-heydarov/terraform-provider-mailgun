#!/bin/bash
# WL Automation - Mailgun Provider Quick Start Script
# Bu script provider-i GitHub-a push edir və release yaradır

set -e

echo "🚀 WL Automation - Mailgun Provider Setup"
echo "=========================================="
echo ""

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Check if we're in the right directory
if [ ! -d "terraform-provider-mailgun" ]; then
    echo -e "${RED}❌ Error: terraform-provider-mailgun directory not found${NC}"
    echo "Please run this script from the workspace root"
    exit 1
fi

# Check if git is installed
if ! command -v git &> /dev/null; then
    echo -e "${RED}❌ Error: git is not installed${NC}"
    exit 1
fi

# Get GitHub repository URL
echo -e "${YELLOW}📝 GitHub Repository Setup${NC}"
read -p "Enter GitHub repository URL (default: https://github.com/murad-heydarov/terraform-provider-mailgun.git): " REPO_URL
REPO_URL=${REPO_URL:-https://github.com/murad-heydarov/terraform-provider-mailgun.git}

# Get version
read -p "Enter version (default: v0.9.0): " VERSION
VERSION=${VERSION:-v0.9.0}

echo ""
echo -e "${GREEN}✅ Configuration:${NC}"
echo "   Repository: $REPO_URL"
echo "   Version: $VERSION"
echo ""

read -p "Continue? (y/n): " -n 1 -r
echo ""
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Aborted."
    exit 1
fi

# Step 1: Initialize git in provider directory
echo ""
echo -e "${YELLOW}📦 Step 1: Initializing git repository...${NC}"
cd terraform-provider-mailgun

if [ -d ".git" ]; then
    echo "   Git already initialized"
else
    git init
    echo "   ✓ Git initialized"
fi

# Step 2: Add all files
echo ""
echo -e "${YELLOW}📝 Step 2: Adding files...${NC}"
git add .
echo "   ✓ Files added"

# Step 3: Create commit
echo ""
echo -e "${YELLOW}💾 Step 3: Creating commit...${NC}"
git commit -m "feat: Enhanced provider with WL automation features

- Add use_automatic_sender_security field for Mailgun API v4 support
- Add trigger_verification to automatically verify domains after DNS
- Add verification_status computed field
- Implement HTTP client for API v4 features not in mailgun-go SDK
- Update documentation with WL automation examples
- Add GitHub Actions workflows for CI/CD
- Add GoReleaser configuration for multi-platform builds

BREAKING CHANGE: Module name changed to github.com/murad-heydarov/terraform-provider-mailgun" || echo "   ✓ Commit already exists or no changes"

# Step 4: Add remote
echo ""
echo -e "${YELLOW}🔗 Step 4: Adding remote repository...${NC}"
git remote remove origin 2>/dev/null || true
git remote add origin "$REPO_URL"
echo "   ✓ Remote added"

# Step 5: Push to main
echo ""
echo -e "${YELLOW}🚀 Step 5: Pushing to GitHub...${NC}"
git branch -M main
if git push -u origin main; then
    echo "   ✓ Pushed to main branch"
else
    echo -e "${RED}   ❌ Push failed. Check your GitHub credentials and repository access${NC}"
    exit 1
fi

# Step 6: Create tag
echo ""
echo -e "${YELLOW}🏷️  Step 6: Creating release tag...${NC}"
git tag -a "$VERSION" -m "Initial enhanced release with WL automation features

Features:
- Automatic sender security (Mailgun API v4)
- Domain verification trigger
- Verification status tracking
- Simplified module implementation
- Complete documentation

This release adds critical features for automated WL deployments."

if git push origin "$VERSION"; then
    echo "   ✓ Tag $VERSION pushed"
else
    echo -e "${RED}   ❌ Tag push failed${NC}"
    exit 1
fi

# Step 7: Wait for GitHub Actions
echo ""
echo -e "${GREEN}✅ Provider successfully published!${NC}"
echo ""
echo "📋 Next steps:"
echo ""
echo "1. Check GitHub Actions: ${REPO_URL}/actions"
echo "   - Wait for workflows to complete (~5 minutes)"
echo "   - Release will be created automatically"
echo ""
echo "2. Verify release: ${REPO_URL}/releases/tag/${VERSION}"
echo ""
echo "3. (Optional) Publish to Terraform Registry:"
echo "   - Go to: https://registry.terraform.io/publish"
echo "   - Sign in with GitHub"
echo "   - Connect repository: murad-heydarov/terraform-provider-mailgun"
echo ""
echo "4. Update your Terraform configuration:"
echo '   terraform {'
echo '     required_providers {'
echo '       mailgun = {'
echo '         source  = "murad-heydarov/mailgun"'
echo "         version = \"~> 0.9.0\""
echo '       }'
echo '     }'
echo '   }'
echo ""
echo "5. Test the provider:"
echo "   cd /path/to/terraform/environments/prod"
echo "   terraform init -upgrade"
echo '   terraform plan -var-file="wl-configs/afftech.auto.tfvars"'
echo ""
echo -e "${GREEN}🎉 All done! Check DEPLOYMENT_GUIDE.md for detailed next steps.${NC}"
