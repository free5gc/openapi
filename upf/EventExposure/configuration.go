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

import (
	"net/http"
	"strings"

	"github.com/free5gc/openapi"
)

type Configuration struct {
	url            string
	basePath       string
	host           string
	defaultHeader  map[string]string
	userAgent      string
	httpClient     *http.Client
	MetricsHook    openapi.RequestMetricsHook
	redirectPolicy openapi.RedirectPolicy
}

func NewConfiguration() *Configuration {
	return &Configuration{
		basePath:      "https://example.com/nupf-ee/v1",
		url:           "{apiRoot}/nupf-ee/v1",
		defaultHeader: make(map[string]string),
		userAgent:     "OpenAPI-Generator/1.0.0/go",
		MetricsHook:   nil,
	}
}

func (c *Configuration) SetBasePath(apiRoot string) {
	c.basePath = strings.Replace(c.url, "{"+"apiRoot"+"}", apiRoot, -1)
}

func (c *Configuration) BasePath() string {
	return c.basePath
}

func (c *Configuration) Host() string {
	return c.host
}

func (c *Configuration) SetHost(host string) {
	c.host = host
}

func (c *Configuration) UserAgent() string {
	return c.userAgent
}

func (c *Configuration) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Configuration) DefaultHeader() map[string]string {
	return c.defaultHeader
}

func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.defaultHeader[key] = value
}

func (c *Configuration) HTTPClient() *http.Client {
	return c.httpClient
}

func (c *Configuration) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *Configuration) Metrics() openapi.RequestMetricsHook {
	return c.MetricsHook
}

func (c *Configuration) SetMetrics(h openapi.RequestMetricsHook) {
	c.MetricsHook = h
}

func (c *Configuration) RedirectPolicy() openapi.RedirectPolicy {
	return c.redirectPolicy
}

func (c *Configuration) SetRedirectPolicy(policy openapi.RedirectPolicy) {
	c.redirectPolicy = policy
}
