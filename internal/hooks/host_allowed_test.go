package hooks

import (
	"testing"

	"github.com/router-for-me/CLIProxyAPI/v7/sdk/pluginapi"

	"github.com/wkeking/clinepass-channel-monitor/internal/config"
	"github.com/wkeking/clinepass-channel-monitor/internal/store"
)

func TestHostAllowedAcceptsObservedPluginExecutorWithoutBaseURL(t *testing.T) {
	cfg := config.Default()
	st := store.New(cfg)
	record := &pluginapi.UsageRecord{Provider: "clinepassbridge", Model: "cline-pass/deepseek-v4.1-flash"}
	if !hostAllowed(record, cfg, st, &store.PendingChannel{}) {
		t.Fatal("observed plugin-executor usage without base_url was rejected")
	}
	if got := st.Totals().SkippedUnmached; got != 0 {
		t.Fatalf("skipped_unmatched_host = %d, want 0", got)
	}
}

func TestHostAllowedStillRejectsUnobservedUsageWithoutBaseURL(t *testing.T) {
	cfg := config.Default()
	st := store.New(cfg)
	record := &pluginapi.UsageRecord{Provider: "other-plugin", Model: "other-model"}
	if hostAllowed(record, cfg, st, nil) {
		t.Fatal("unobserved usage without base_url was accepted")
	}
	if got := st.Totals().SkippedUnmached; got != 1 {
		t.Fatalf("skipped_unmatched_host = %d, want 1", got)
	}
}
