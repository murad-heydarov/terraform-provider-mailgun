package mailgun

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailgun/mailgun-go/v5"
	"github.com/mailgun/mailgun-go/v5/mtypes"
)

func resourceMailgunDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceMailgunDomainCreate,
		ReadContext:   resourceMailgunDomainRead,
		UpdateContext: resourceMailgunDomainUpdate,
		DeleteContext: resourceMailgunDomainDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceMailgunDomainImport,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"region": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "us",
			},

			"spam_action": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "disabled",
			},

			"smtp_login": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"smtp_password": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  false,
				Sensitive: true,
			},

			"wildcard": {
				Type:     schema.TypeBool,
				ForceNew: true,
				Optional: true,
			},

			"dkim_selector": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"force_dkim_authority": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},

			"open_tracking": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},

			"click_tracking": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
			},

			"web_scheme": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "http",
			},

			"use_automatic_sender_security": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  true,
				Description: "Enable automatic sender security (Mailgun API v4)",
			},

			"trigger_verification": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  false,
				Description: "Trigger domain verification after DNS records are added",
			},

			"verification_status": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Domain verification status",
			},

			"receiving_records": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "Use `receiving_records_set` instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"receiving_records_set": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      domainRecordsSchemaSetFunc,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"sending_records": {
				Type:       schema.TypeList,
				Computed:   true,
				Deprecated: "Use `sending_records_set` instead.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"sending_records_set": {
				Type:     schema.TypeSet,
				Computed: true,
				Set:      domainRecordsSchemaSetFunc,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"valid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dkim_key_size": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
		},
		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				if diff.HasChange("name") {
					var sendingRecords []interface{}

					sendingRecords = append(sendingRecords, map[string]interface{}{"id": diff.Get("name").(string)})
					sendingRecords = append(sendingRecords, map[string]interface{}{"id": "_domainkey." + diff.Get("name").(string)})
					sendingRecords = append(sendingRecords, map[string]interface{}{"id": "email." + diff.Get("name").(string)})

					if err := diff.SetNew("sending_records_set", schema.NewSet(domainRecordsSchemaSetFunc, sendingRecords)); err != nil {
						return fmt.Errorf("error setting new sending_records_set diff: %w", err)
					}

					var receivingRecords []interface{}

					receivingRecords = append(receivingRecords, map[string]interface{}{"id": "mxa.mailgun.org"})
					receivingRecords = append(receivingRecords, map[string]interface{}{"id": "mxb.mailgun.org"})

					if err := diff.SetNew("receiving_records_set", schema.NewSet(domainRecordsSchemaSetFunc, receivingRecords)); err != nil {
						return fmt.Errorf("error setting new receiving_records_set diff: %w", err)
					}
				}

				return nil
			},
		),
	}
}

func domainRecordsSchemaSetFunc(v interface{}) int {
	m, ok := v.(map[string]interface{})

	if !ok {
		return 0
	}

	if v, ok := m["id"].(string); ok {
		return stringHashcode(v)
	}

	return 0
}

func resourceMailgunDomainImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	setDefaultRegionForImport(d)

	return []*schema.ResourceData{d}, nil
}

func resourceMailgunDomainUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var name = d.Get("name").(string)
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	var currentData schema.ResourceData
	var newPassword = d.Get("smtp_password").(string)
	var smtpLogin = d.Get("smtp_login").(string)
	var openTracking = d.Get("open_tracking").(bool)
	var clickTracking = d.Get("click_tracking").(bool)
	var webScheme = d.Get("web_scheme").(string)

	// Retrieve and update state of domain
	_, errc = resourceMailgunDomainRetrieve(d.Id(), client, &currentData)

	if errc != nil {
		return diag.FromErr(errc)
	}

	// Update default credential if changed
	if currentData.Get("smtp_password") != newPassword {
		errc = client.ChangeCredentialPassword(ctx, name, smtpLogin, newPassword)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	if currentData.Get("open_tracking") != openTracking {
		var openTrackingValue = "no"
		if openTracking {
			openTrackingValue = "yes"
		}
		errc = client.UpdateOpenTracking(ctx, name, openTrackingValue)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	if currentData.Get("click_tracking") != clickTracking {
		var clickTrackingValue = "no"
		if clickTracking {
			clickTrackingValue = "yes"
		}
		errc = client.UpdateClickTracking(ctx, d.Get("name").(string), clickTrackingValue)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	if currentData.Get("web_scheme") != webScheme {
		opts := mailgun.UpdateDomainOptions{}
		opts.WebScheme = webScheme
		errc = client.UpdateDomain(ctx, name, &opts)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	// Handle automatic sender security update
	if d.HasChange("use_automatic_sender_security") && d.Get("use_automatic_sender_security").(bool) {
		config := meta.(*Config)
		err := setAutomaticSenderSecurity(ctx, name, config.APIKey, d.Get("region").(string))
		if err != nil {
			log.Printf("[WARN] Failed to set automatic sender security: %v", err)
		}
	}

	// Handle verification trigger
	if d.HasChange("trigger_verification") && d.Get("trigger_verification").(bool) {
		config := meta.(*Config)
		status, err := verifyDomain(ctx, name, config.APIKey, d.Get("region").(string))
		if err != nil {
			log.Printf("[WARN] Failed to trigger verification: %v", err)
		} else {
			_ = d.Set("verification_status", status)
		}
	}

	return nil
}

func resourceMailgunDomainCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	opts := mailgun.CreateDomainOptions{}

	name := d.Get("name").(string)

	opts.SpamAction = mtypes.SpamAction(d.Get("spam_action").(string))
	opts.Password = d.Get("smtp_password").(string)
	opts.Wildcard = d.Get("wildcard").(bool)
	opts.DKIMKeySize = d.Get("dkim_key_size").(int)
	opts.ForceDKIMAuthority = d.Get("force_dkim_authority").(bool)
	opts.WebScheme = d.Get("web_scheme").(string)
	var dkimSelector = d.Get("dkim_selector").(string)
	var openTracking = d.Get("open_tracking").(bool)
	var clickTracking = d.Get("click_tracking").(bool)

	log.Printf("[DEBUG] Domain create configuration: %#v", opts)

	_, err := client.CreateDomain(context.Background(), name, &opts)

	if err != nil {
		return diag.FromErr(err)
	}

	if dkimSelector != "" {
		errc = client.UpdateDomainDkimSelector(ctx, name, dkimSelector)

		if errc != nil {
			return diag.FromErr(errc)
		}
	}
	if openTracking {
		errc = client.UpdateOpenTracking(ctx, name, "yes")

		if errc != nil {
			return diag.FromErr(errc)
		}
	}
	if clickTracking {
		errc = client.UpdateClickTracking(ctx, d.Get("name").(string), "yes")

		if errc != nil {
			return diag.FromErr(errc)
		}
	}

	d.SetId(name)

	log.Printf("[INFO] Domain ID: %s", d.Id())

	// Set automatic sender security if enabled
	if d.Get("use_automatic_sender_security").(bool) {
		config := meta.(*Config)
		err := setAutomaticSenderSecurity(ctx, name, config.APIKey, d.Get("region").(string))
		if err != nil {
			log.Printf("[WARN] Failed to set automatic sender security: %v", err)
			// Don't fail the entire create operation
		}
	}

	// Retrieve and update state of domain
	_, err = resourceMailgunDomainRetrieve(d.Id(), client, d)

	if err != nil {
		return diag.FromErr(errc)
	}

	// Trigger verification if requested
	if d.Get("trigger_verification").(bool) {
		config := meta.(*Config)
		status, err := verifyDomain(ctx, name, config.APIKey, d.Get("region").(string))
		if err != nil {
			log.Printf("[WARN] Failed to trigger verification: %v", err)
		} else {
			_ = d.Set("verification_status", status)
		}
	}

	return nil
}

func resourceMailgunDomainDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	log.Printf("[INFO] Deleting Domain: %s", d.Id())

	// Destroy the domain
	err := client.DeleteDomain(context.Background(), d.Id())
	if err != nil {
		return diag.Errorf("Error deleting domain: %s", err)
	}

	// Give the destroy a chance to take effect
	err = resource.RetryContext(ctx, 5*time.Minute, func() *resource.RetryError {
		_, err = client.GetDomain(ctx, d.Id(), nil)
		if err == nil {
			log.Printf("[INFO] Retrying until domain disappears...")
			return resource.RetryableError(
				fmt.Errorf("domain seems to still exist; will check again"))
		}
		log.Printf("[INFO] Got error looking for domain, seems gone: %s", err)
		return nil
	})

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMailgunDomainRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	_, err := resourceMailgunDomainRetrieve(d.Id(), client, d)

	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMailgunDomainRetrieve(id string, client *mailgun.Client, d *schema.ResourceData) (*mtypes.GetDomainResponse, error) {

	resp, err := client.GetDomain(context.Background(), id, nil)

	if err != nil {
		return nil, fmt.Errorf("Error retrieving domain: %s", err)
	}

	_ = d.Set("name", resp.Domain.Name)
	_ = d.Set("smtp_login", resp.Domain.SMTPLogin)
	_ = d.Set("wildcard", resp.Domain.Wildcard)
	_ = d.Set("spam_action", resp.Domain.SpamAction)
	_ = d.Set("web_scheme", resp.Domain.WebScheme)

	receivingRecords := make([]map[string]interface{}, len(resp.ReceivingDNSRecords))
	for i, r := range resp.ReceivingDNSRecords {
		receivingRecords[i] = make(map[string]interface{})
		receivingRecords[i]["id"] = r.Value
		receivingRecords[i]["priority"] = r.Priority
		receivingRecords[i]["valid"] = r.Valid
		receivingRecords[i]["value"] = r.Value
		receivingRecords[i]["record_type"] = r.RecordType
	}
	_ = d.Set("receiving_records", receivingRecords)
	_ = d.Set("receiving_records_set", receivingRecords)

	sendingRecords := make([]map[string]interface{}, len(resp.SendingDNSRecords))
	for i, r := range resp.SendingDNSRecords {
		sendingRecords[i] = make(map[string]interface{})
		sendingRecords[i]["id"] = r.Name
		sendingRecords[i]["name"] = r.Name
		sendingRecords[i]["valid"] = r.Valid
		sendingRecords[i]["value"] = r.Value
		sendingRecords[i]["record_type"] = r.RecordType

		if strings.Contains(r.Name, "._domainkey.") {
			sendingRecords[i]["id"] = "_domainkey." + resp.Domain.Name
		}
	}
	_ = d.Set("sending_records", sendingRecords)
	_ = d.Set("sending_records_set", sendingRecords)

	info, err := client.GetDomainTracking(context.Background(), id)
	var openTracking = false
	if info.Open.Active {
		openTracking = true
	}
	_ = d.Set("open_tracking", openTracking)

	var clickTracking = false
	if info.Click.Active {
		clickTracking = true
	}
	_ = d.Set("click_tracking", clickTracking)

	return &resp, nil
}

// setAutomaticSenderSecurity enables automatic sender security via API v4
// Uses multipart/form-data as required by Mailgun API v4
func setAutomaticSenderSecurity(ctx context.Context, domain string, apiKey string, region string) error {
	baseURL := "https://api.mailgun.net"
	if strings.ToLower(region) == "eu" {
		baseURL = "https://api.eu.mailgun.net"
	}

	url := fmt.Sprintf("%s/v4/domains/%s", baseURL, domain)

	// Create multipart form data
	body := &bytes.Buffer{}
	body.WriteString("use_automatic_sender_security=true")

	req, err := http.NewRequestWithContext(ctx, "PUT", url, body)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.SetBasicAuth("api", apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API v4 request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	log.Printf("[INFO] Automatic sender security enabled for domain: %s (API v4)", domain)
	return nil
}

// verifyDomain triggers domain verification via API v3
func verifyDomain(ctx context.Context, domain string, apiKey string, region string) (string, error) {
	baseURL := "https://api.mailgun.net"
	if strings.ToLower(region) == "eu" {
		baseURL = "https://api.eu.mailgun.net"
	}

	url := fmt.Sprintf("%s/v3/domains/%s/verify", baseURL, domain)

	req, err := http.NewRequestWithContext(ctx, "PUT", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.SetBasicAuth("api", apiKey)

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("[WARN] Domain verification returned status %d: %s", resp.StatusCode, string(body))
		return "verification_failed", nil
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("[WARN] Failed to parse verification response: %v", err)
		return "verification_triggered", nil
	}

	log.Printf("[INFO] Domain verification triggered for: %s", domain)
	return "verification_triggered", nil
}
