package mailgun

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mailgun/mailgun-go/v5"
	"github.com/mailgun/mailgun-go/v5/mtypes"
)

const (
	defaultVerificationTimeout     = 10 * time.Minute
	defaultVerificationPollSeconds = 15 * time.Second
)

func resourceMailgunDomainVerification() *schema.Resource {
	recordSchema := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"record_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"valid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cached": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}

	return &schema.Resource{
		CreateContext: resourceMailgunDomainVerificationCreate,
		ReadContext:   resourceMailgunDomainVerificationRead,
		DeleteContext: resourceMailgunDomainVerificationDelete,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "us",
				ForceNew: true,
			},
			"wait_for_active": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Wait until all DNS records are reported as valid by Mailgun before completing.",
				ForceNew:    true,
			},
			"poll_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultVerificationPollSeconds.String(),
				Description: "Polling interval used while waiting for verification (for example, \"15s\").",
				ForceNew:    true,
			},
			"timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     defaultVerificationTimeout.String(),
				Description: "Maximum time to wait for Mailgun to report all DNS records as valid (for example, \"10m\").",
				ForceNew:    true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"receiving_records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     recordSchema,
			},
			"sending_records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     recordSchema,
			},
		},
	}
}

func resourceMailgunDomainVerificationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	domain := d.Get("domain").(string)

	resp, err := client.VerifyDomain(ctx, domain)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setVerificationResourceState(d, &resp); err != nil {
		return diag.FromErr(err)
	}

	if d.Get("wait_for_active").(bool) {
		timeout, err := parseDurationWithDefault(d.Get("timeout").(string), defaultVerificationTimeout)
		if err != nil {
			return diag.FromErr(err)
		}

		pollInterval, err := parseDurationWithDefault(d.Get("poll_interval").(string), defaultVerificationPollSeconds)
		if err != nil {
			return diag.FromErr(err)
		}

		finalResp, err := waitForDomainVerification(ctx, client, domain, timeout, pollInterval)
		if err != nil {
			return diag.FromErr(err)
		}

		if err := setVerificationResourceState(d, finalResp); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(domain)

	return resourceMailgunDomainVerificationRead(ctx, d, meta)
}

func resourceMailgunDomainVerificationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client, errc := meta.(*Config).GetClient(d.Get("region").(string))
	if errc != nil {
		return diag.FromErr(errc)
	}

	resp, err := client.GetDomain(ctx, d.Id(), nil)
	if err != nil {
		if mailgun.GetStatusFromErr(err) == http.StatusNotFound {
			d.SetId("")
			return nil
		}

		return diag.FromErr(err)
	}

	if err := setVerificationResourceState(d, &resp); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceMailgunDomainVerificationDelete(_ context.Context, d *schema.ResourceData, _ interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func setVerificationResourceState(d *schema.ResourceData, resp *mtypes.GetDomainResponse) error {
	if err := d.Set("status", resp.Domain.State); err != nil {
		return err
	}

	if err := d.Set("receiving_records", flattenDomainVerificationRecords(resp.ReceivingDNSRecords)); err != nil {
		return err
	}

	if err := d.Set("sending_records", flattenDomainVerificationRecords(resp.SendingDNSRecords)); err != nil {
		return err
	}

	return nil
}

func flattenDomainVerificationRecords(records []mtypes.DNSRecord) []map[string]interface{} {
	result := make([]map[string]interface{}, len(records))

	for i, record := range records {
		cached := make([]string, len(record.Cached))
		copy(cached, record.Cached)

		result[i] = map[string]interface{}{
			"name":        record.Name,
			"record_type": record.RecordType,
			"value":       record.Value,
			"priority":    record.Priority,
			"valid":       record.Valid,
			"is_active":   record.Active,
			"cached":      cached,
		}
	}

	return result
}

func waitForDomainVerification(ctx context.Context, client *mailgun.Client, domain string, timeout, pollInterval time.Duration) (*mtypes.GetDomainResponse, error) {
	stateConf := &resource.StateChangeConf{
		Pending:        []string{"pending"},
		Target:         []string{"verified"},
		Timeout:        timeout,
		MinTimeout:     pollInterval,
		PollInterval:   pollInterval,
		NotFoundChecks: 5,
		Refresh: func() (interface{}, string, error) {
			resp, err := client.GetDomain(ctx, domain, nil)
			if err != nil {
				return nil, "", err
			}

			if domainRecordsAreValid(&resp) {
				return &resp, "verified", nil
			}

			return &resp, "pending", nil
		},
	}

	result, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}

	resp, ok := result.(*mtypes.GetDomainResponse)
	if !ok {
		return nil, fmt.Errorf("unexpected verification response type %T", result)
	}

	return resp, nil
}

func domainRecordsAreValid(resp *mtypes.GetDomainResponse) bool {
	for _, record := range resp.ReceivingDNSRecords {
		if !strings.EqualFold(record.Valid, "valid") {
			return false
		}
	}

	for _, record := range resp.SendingDNSRecords {
		if !strings.EqualFold(record.Valid, "valid") {
			return false
		}
	}

	return true
}

func parseDurationWithDefault(raw string, fallback time.Duration) (time.Duration, error) {
	trimmed := strings.TrimSpace(raw)
	if trimmed == "" {
		return fallback, nil
	}

	duration, err := time.ParseDuration(trimmed)
	if err != nil {
		return 0, fmt.Errorf("invalid duration value %q: %w", raw, err)
	}

	return duration, nil
}
