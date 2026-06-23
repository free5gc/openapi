/*
 * Nupf_EventExposure
 *
 * UPF Event Exposure Service.   Copyright 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.564 V19.5.0; 5G System; User Plane Function Services; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.564/
 *
 * API version: 1.2.0
 */

package EventExposure

// APIClient manages communication with the Nupf_EventExposure API v1.2.0.
type APIClient struct {
	cfg    *Configuration
	common service

	IndividualSubscriptionDocumentApi *IndividualSubscriptionDocumentApiService
	SubscriptionsCollectionApi        *SubscriptionsCollectionApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	c.IndividualSubscriptionDocumentApi = (*IndividualSubscriptionDocumentApiService)(&c.common)
	c.SubscriptionsCollectionApi = (*SubscriptionsCollectionApiService)(&c.common)

	return c
}
