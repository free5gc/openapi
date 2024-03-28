/*
 * LMF Broadcast
 *
 * LMF Broadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.572 V17.5.0; 5G System; Location Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.572/
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Broadcast

import (
	"net/http"
	"strings"
)

type Configuration struct {
	url           string
	basePath      string
	host          string
	defaultHeader map[string]string
	userAgent     string
	httpClient    *http.Client
}

func NewConfiguration() *Configuration {
	cfg := &Configuration{
		basePath:      "https://example.com/nlmf-broadcast/v1",
		url:           "{apiRoot}/nlmf-broadcast/v1",
		defaultHeader: make(map[string]string),
		userAgent:     "OpenAPI-Generator/1.0.0/go",
	}
	return cfg
}

func (c *Configuration) SetBasePath(apiRoot string) {
	url := c.url

	// Replace apiRoot
	url = strings.Replace(url, "{"+"apiRoot"+"}", apiRoot, -1)

	c.basePath = url
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