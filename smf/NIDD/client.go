/*
 * Nsmf_NIDD
 *
 * SMF NIDD Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.542 V17.4.0; 5G System; Session Management Services for Non-IP Data Delivery (NIDD); Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.542/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package NIDD

// APIClient manages communication with the Nsmf_NIDD API v1.1.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	IndividualPDUSessionApi *IndividualPDUSessionApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.IndividualPDUSessionApi = (*IndividualPDUSessionApiService)(&c.common)

	return c
}
