package main

import (
	"net/http"
	"testing"
)

// TestProbeEmptyListReturnsArrayNotNull asserts that GET /blobs on an empty
// store returns an empty JSON array "[]" for the blobs field, not null.
func TestProbeEmptyListReturnsArrayNotNull(t *testing.T) {
	_, h := newTestServer()
	code, body := doJSON(t, h, http.MethodGet, "/blobs", nil)
	if code != http.StatusOK {
		t.Fatalf("code = %d body=%v", code, body)
	}
	blobs, ok := body["blobs"]
	if !ok {
		t.Fatalf("missing blobs field: %v", body)
	}
	if blobs == nil {
		t.Fatalf("blobs is null; want empty array []")
	}
	arr, ok := blobs.([]any)
	if !ok {
		t.Fatalf("blobs is %T; want JSON array", blobs)
	}
	if len(arr) != 0 {
		t.Fatalf("blobs len = %d want 0", len(arr))
	}
}
