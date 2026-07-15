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
	"context"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

type SubscriptionsCollectionApiService service

// CreateSubscriptionRequest contains the Nupf Event Exposure create request body.
type CreateSubscriptionRequest struct {
	UpfCreateEventSubscription *models.UpfCreateEventSubscription
}

func (r *CreateSubscriptionRequest) SetUpfCreateEventSubscription(subscription models.UpfCreateEventSubscription) {
	r.UpfCreateEventSubscription = &subscription
}

type CreateSubscriptionResponse struct {
	Location                    string
	UpfCreatedEventSubscription models.UpfCreatedEventSubscription
}

type CreateSubscriptionError struct {
	Location             string
	Var3gppSbiTargetNfId string
	ProblemDetails       models.ProblemDetails
	RedirectResponse     models.RedirectResponse
}

// CreateSubscription creates an individual UPF Event Exposure subscription.
func (a *SubscriptionsCollectionApiService) CreateSubscription(ctx context.Context, request *CreateSubscriptionRequest) (*CreateSubscriptionResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreateSubscriptionResponse
	)

	localVarPath := a.client.cfg.BasePath() + "/ee-subscriptions"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json, application/problem+json"
	localVarPostBody = request.UpfCreateEventSubscription

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	if err = localVarHTTPResponse.Body.Close(); err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.UpfCreatedEventSubscription, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		return &localVarReturnValue, nil
	case 307, 308:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.RedirectResponse, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		v.Location = localVarHTTPResponse.Header.Get("Location")
		v.Var3gppSbiTargetNfId = localVarHTTPResponse.Header.Get("3gpp-Sbi-Target-Nf-Id")
		apiError.ErrorModel = v
		return nil, apiError
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 501, 502, 503:
		var v CreateSubscriptionError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	default:
		return nil, apiError
	}
}
