package mailgun

import (
	"testing"
	"time"

	"github.com/mailgun/mailgun-go/v5/mtypes"
)

func TestParseDurationWithDefault(t *testing.T) {
	t.Parallel()

	d, err := parseDurationWithDefault("30s", time.Minute)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if d != 30*time.Second {
		t.Fatalf("expected 30s duration, got %s", d)
	}

	d, err = parseDurationWithDefault("", time.Minute)
	if err != nil {
		t.Fatalf("unexpected error for fallback: %v", err)
	}

	if d != time.Minute {
		t.Fatalf("expected fallback duration, got %s", d)
	}
}

func TestDomainRecordsAreValid(t *testing.T) {
	t.Parallel()

	resp := &mtypes.GetDomainResponse{
		ReceivingDNSRecords: []mtypes.DNSRecord{
			{Valid: "valid"},
		},
		SendingDNSRecords: []mtypes.DNSRecord{
			{Valid: "VALID"},
		},
	}

	if !domainRecordsAreValid(resp) {
		t.Fatal("expected records to be valid")
	}

	resp.ReceivingDNSRecords[0].Valid = "pending"

	if domainRecordsAreValid(resp) {
		t.Fatal("expected records to be invalid")
	}
}
