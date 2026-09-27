package pokeapi

import "testing"

func TestGetLocationAreas(t *testing.T) {
	resp, err := GetLocationAreas(BaseURL)
	if err != nil {
		t.Fatalf("GetLocationAreas returned error: %v", err)
	}

	if len(resp.Results) != 20 {
		t.Errorf("expected 20 results, got %d", len(resp.Results))
	}

	if resp.Next == nil {
		t.Errorf("expected Next to be set, got nil")
	}

	if resp.Previous != nil {
		t.Errorf("expected Previous to be nil on first page, got %q", *resp.Previous)
	}

	if resp.Count <= 0 {
		t.Errorf("expected Count > 0, got %d", resp.Count)
	}
}
