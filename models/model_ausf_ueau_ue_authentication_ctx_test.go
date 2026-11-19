package models

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAusfUEAuthenticationCtxUnmarshalJSON(t *testing.T) {
	const uri1 = "http://ausf.example/confirmation/1"
	const uri2 = "http://ausf.example/confirmation/2"

	tests := []struct {
		name      string
		input     string
		wantLinks map[string][]Link
		wantErr   bool
	}{
		{
			name:      "single object",
			input:     `{"authType":"5G_AKA","_links":{"5g-aka":{"href":"` + uri1 + `"}},"servingNetworkName":"5G:mnc001.mcc001.3gppnetwork.org"}`,
			wantLinks: map[string][]Link{"5g-aka": {{Href: uri1}}},
		},
		{
			name:      "one-element array",
			input:     `{"_links":{"5g-aka":[{"href":"` + uri1 + `"}]}}`,
			wantLinks: map[string][]Link{"5g-aka": {{Href: uri1}}},
		},
		{
			name:      "multiple links in array",
			input:     `{"_links":{"5g-aka":[{"href":"` + uri1 + `"},{"href":"` + uri2 + `"}]}}`,
			wantLinks: map[string][]Link{"5g-aka": {{Href: uri1}, {Href: uri2}}},
		},
		{
			name:      "mixed object and array values",
			input:     `{"_links":{"5g-aka":{"href":"` + uri1 + `"},"eap-session":[{"href":"` + uri2 + `"}]}}`,
			wantLinks: map[string][]Link{"5g-aka": {{Href: uri1}}, "eap-session": {{Href: uri2}}},
		},
		{
			name:      "empty link object remains accepted",
			input:     `{"_links":{"5g-aka":{}}}`,
			wantLinks: map[string][]Link{"5g-aka": {{}}},
		},
		{
			name:      "empty links map",
			input:     `{"_links":{}}`,
			wantLinks: map[string][]Link{},
		},
		{
			name:    "empty array",
			input:   `{"_links":{"5g-aka":[]}}`,
			wantErr: true,
		},
		{
			name:    "null link value",
			input:   `{"_links":{"5g-aka":null}}`,
			wantErr: true,
		},
		{
			name:    "string link value",
			input:   `{"_links":{"5g-aka":"invalid"}}`,
			wantErr: true,
		},
		{
			name:    "numeric link value",
			input:   `{"_links":{"5g-aka":42}}`,
			wantErr: true,
		},
		{
			name:    "wrong href type",
			input:   `{"_links":{"5g-aka":{"href":42}}}`,
			wantErr: true,
		},
		{
			name:    "invalid element in array",
			input:   `{"_links":{"5g-aka":[{"href":"valid"},42]}}`,
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			input:   `{"_links":`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Ausf_UEAU_UEAuthenticationCtx
			err := json.Unmarshal([]byte(tt.input), &got)
			if (err != nil) != tt.wantErr {
				t.Fatalf("json.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got.Links, tt.wantLinks) {
				t.Errorf("Links = %#v, want %#v", got.Links, tt.wantLinks)
			}
			if tt.name == "single object" {
				if string(got.AuthType) != "5G_AKA" {
					t.Errorf("AuthType = %q, want 5G_AKA", got.AuthType)
				}
				if got.ServingNetworkName != "5G:mnc001.mcc001.3gppnetwork.org" {
					t.Errorf("ServingNetworkName = %q", got.ServingNetworkName)
				}
			}
		})
	}
}

func TestAusfUEAuthenticationCtxUnmarshalJSONDoesNotPartiallyUpdate(t *testing.T) {
	original := Ausf_UEAU_UEAuthenticationCtx{
		ServingNetworkName: "original-network",
		Links: map[string][]Link{
			"original": {{Href: "http://existing.example/"}},
		},
	}

	// The JSON changes a normal field, but its _links contains an invalid value.
	// The receiver must remain unchanged when parsing returns an error.
	input := []byte(`{"servingNetworkName":"new-network","_links":{"5g-aka":42}}`)
	if err := json.Unmarshal(input, &original); err == nil {
		t.Fatal("json.Unmarshal() succeeded, want an error")
	}
	if original.ServingNetworkName != "original-network" {
		t.Errorf("ServingNetworkName changed after error: %q", original.ServingNetworkName)
	}
	wantLinks := map[string][]Link{"original": {{Href: "http://existing.example/"}}}
	if !reflect.DeepEqual(original.Links, wantLinks) {
		t.Errorf("Links changed after error: %#v", original.Links)
	}
}

func TestAusfUEAuthenticationCtxMarshalJSONKeepsArrayFormat(t *testing.T) {
	ctx := Ausf_UEAU_UEAuthenticationCtx{
		Links: map[string][]Link{
			"5g-aka": {{Href: "http://ausf.example/confirmation"}},
		},
	}
	data, err := json.Marshal(&ctx)
	if err != nil {
		t.Fatalf("json.Marshal() error: %v", err)
	}
	var result struct {
		Links map[string]json.RawMessage `json:"_links"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("parse marshaled JSON: %v", err)
	}
	if string(result.Links["5g-aka"][0]) != "[" {
		t.Errorf("marshaled link = %s, want JSON array", result.Links["5g-aka"])
	}
}
