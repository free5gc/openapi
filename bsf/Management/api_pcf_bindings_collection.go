/*
 * Nbsf_Management
 *
 * Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.521 V17.7.0; 5G System; Binding Support Management Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.521/
 *
 * API version: 1.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Management

import (
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"

	"context"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type PCFBindingsCollectionApiService service

/*
PCFBindingsCollectionApiService Create a new Individual PCF for a PDU Session binding information
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param PcfBinding -

@return CreatePCFBindingResponse
*/

// CreatePCFBindingRequest
type CreatePCFBindingRequest struct {
	PcfBinding *models.PcfBinding
}

func (r *CreatePCFBindingRequest) SetPcfBinding(PcfBinding models.PcfBinding) {
	r.PcfBinding = &PcfBinding
}

type CreatePCFBindingResponse struct {
	Location   string
	PcfBinding models.PcfBinding
}

type CreatePCFBindingError struct {
	BsfManagementExtProblemDetails models.BsfManagementExtProblemDetails
	ProblemDetails                 models.ProblemDetails
}

func (a *PCFBindingsCollectionApiService) CreatePCFBinding(ctx context.Context, request *CreatePCFBindingRequest) (*CreatePCFBindingResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CreatePCFBindingResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/pcfBindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = request.PcfBinding

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
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 201:
		err = openapi.Deserialize(&localVarReturnValue.PcfBinding, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		localVarReturnValue.Location = localVarHTTPResponse.Header.Get("Location")
		return &localVarReturnValue, nil
	case 400:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.BsfManagementExtProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 411:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 413:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 415:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v CreatePCFBindingError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v CreatePCFBindingError
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

/*
PCFBindingsCollectionApiService Read PCF for a PDU Session Bindings information
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param Ipv4Addr - The IPv4 Address of the served UE.
 * @param Ipv6Prefix - The IPv6 Address of the served UE. The NF service consumer shall append '/128' to the  IPv6 address in the attribute value. E.g. '2001:db8:85a3::8a2e:370:7334/128'.
 * @param MacAddr48 - The MAC Address of the served UE.
 * @param Dnn - DNN.
 * @param Supi - Subscription Permanent Identifier.
 * @param Gpsi - Generic Public Subscription Identifier
 * @param Snssai - The identification of slice.
 * @param IpDomain - The IPv4 address domain identifier.
 * @param SuppFeat - To filter irrelevant responses related to unsupported features.

@return GetPCFBindingsResponse
*/

// GetPCFBindingsRequest
type GetPCFBindingsRequest struct {
	Ipv4Addr   *string
	Ipv6Prefix *string
	MacAddr48  *string
	Dnn        *string
	Supi       *string
	Gpsi       *string
	Snssai     *models.Snssai
	IpDomain   *string
	SuppFeat   *string
}

func (r *GetPCFBindingsRequest) SetIpv4Addr(Ipv4Addr string) {
	r.Ipv4Addr = &Ipv4Addr
}
func (r *GetPCFBindingsRequest) SetIpv6Prefix(Ipv6Prefix string) {
	r.Ipv6Prefix = &Ipv6Prefix
}
func (r *GetPCFBindingsRequest) SetMacAddr48(MacAddr48 string) {
	r.MacAddr48 = &MacAddr48
}
func (r *GetPCFBindingsRequest) SetDnn(Dnn string) {
	r.Dnn = &Dnn
}
func (r *GetPCFBindingsRequest) SetSupi(Supi string) {
	r.Supi = &Supi
}
func (r *GetPCFBindingsRequest) SetGpsi(Gpsi string) {
	r.Gpsi = &Gpsi
}
func (r *GetPCFBindingsRequest) SetSnssai(Snssai models.Snssai) {
	r.Snssai = &Snssai
}
func (r *GetPCFBindingsRequest) SetIpDomain(IpDomain string) {
	r.IpDomain = &IpDomain
}
func (r *GetPCFBindingsRequest) SetSuppFeat(SuppFeat string) {
	r.SuppFeat = &SuppFeat
}

type GetPCFBindingsResponse struct {
	PcfBinding models.PcfBinding
}

type GetPCFBindingsError struct {
	ProblemDetails models.ProblemDetails
}

func (a *PCFBindingsCollectionApiService) GetPCFBindings(ctx context.Context, request *GetPCFBindingsRequest) (*GetPCFBindingsResponse, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetPCFBindingsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/pcfBindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if request.Ipv4Addr != nil {
		localVarQueryParams.Add("ipv4Addr", openapi.ParameterToString(request.Ipv4Addr, "multi"))
	}
	if request.Ipv6Prefix != nil {
		localVarQueryParams.Add("ipv6Prefix", openapi.ParameterToString(request.Ipv6Prefix, "multi"))
	}
	if request.MacAddr48 != nil {
		localVarQueryParams.Add("macAddr48", openapi.ParameterToString(request.MacAddr48, "multi"))
	}
	if request.Dnn != nil {
		localVarQueryParams.Add("dnn", openapi.ParameterToString(request.Dnn, "multi"))
	}
	if request.Supi != nil {
		localVarQueryParams.Add("supi", openapi.ParameterToString(request.Supi, "multi"))
	}
	if request.Gpsi != nil {
		localVarQueryParams.Add("gpsi", openapi.ParameterToString(request.Gpsi, "multi"))
	}
	if request.Snssai != nil {
		localVarQueryParams.Add("snssai", openapi.ParameterToString(request.Snssai, "application/json"))
	}
	if request.IpDomain != nil {
		localVarQueryParams.Add("ipDomain", openapi.ParameterToString(request.IpDomain, "multi"))
	}
	if request.SuppFeat != nil {
		localVarQueryParams.Add("supp-feat", openapi.ParameterToString(request.SuppFeat, "multi"))
	}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := strings.Join(localVarHTTPHeaderAccepts, ", ")
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

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
	err = localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.StatusCode,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue.PcfBinding, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		return &localVarReturnValue, nil
	case 204:
		return &localVarReturnValue, nil
	case 400:
		var v GetPCFBindingsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 401:
		var v GetPCFBindingsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 403:
		var v GetPCFBindingsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 404:
		var v GetPCFBindingsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 406:
		return &localVarReturnValue, nil
	case 414:
		var v GetPCFBindingsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 429:
		var v GetPCFBindingsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 500:
		var v GetPCFBindingsError
		err = openapi.Deserialize(&v.ProblemDetails, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			return nil, err
		}
		apiError.ErrorModel = v
		return nil, apiError
	case 503:
		var v GetPCFBindingsError
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