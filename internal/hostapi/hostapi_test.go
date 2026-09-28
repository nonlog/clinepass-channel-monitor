package hostapi

import (
	"bytes"
	"testing"
)

func TestUnwrapHostResponseEnvelope(t *testing.T) {
	raw := []byte(`{"ok":true,"result":{"files":[{"name":"cline.json"}]}}`)
	got, ok := unwrapHostResponse(raw)
	if !ok {
		t.Fatal("unwrapHostResponse() ok = false, want true")
	}
	want := []byte(`{"files":[{"name":"cline.json"}]}`)
	if !bytes.Equal(got, want) {
		t.Fatalf("unwrapHostResponse() = %s, want %s", got, want)
	}
}

func TestUnwrapHostResponseErrorEnvelope(t *testing.T) {
	if got, ok := unwrapHostResponse([]byte(`{"ok":false,"error":{"code":"host_error","message":"boom"}}`)); ok || got != nil {
		t.Fatalf("unwrapHostResponse(error) = (%q, %v), want (nil, false)", got, ok)
	}
}

func TestUnwrapHostResponseLegacyRawPayload(t *testing.T) {
	raw := []byte(`{"files":[{"name":"cline.json"}]}`)
	got, ok := unwrapHostResponse(raw)
	if !ok || !bytes.Equal(got, raw) {
		t.Fatalf("unwrapHostResponse(legacy) = (%s, %v), want unchanged payload", got, ok)
	}
}
