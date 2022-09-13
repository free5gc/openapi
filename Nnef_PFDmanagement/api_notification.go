package Nnef_PFDmanagement

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/free5gc/openapi"
	. "github.com/free5gc/openapi/models"
)

// Linger please
var (
	_ context.Context
)

type NotificationApiService service

/*
NotificationApiService Notifies about update to PFD change
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param notifyUri string - Notification URI provided in subscription
 * @param pfdChangeNotifications []PfdChangeNotification - Notifications about updates to PFDs
*/

func (a *NotificationApiService) NotificationPost(ctx context.Context, notifyUri string, pfdChangeNotifications []PfdChangeNotification) ([]PfdChangeReport, *http.Response, error) {
	var (
		localVarHTTPMethod	= strings.ToUpper("Post")
		localVarPostBody	interface{}
		localVarFormFileName	string
		localVarFileName	string
		localVarFileBytes	[]byte
		localVarReturnValue	[]PfdChangeReport
	)

	// create path and map variables
	localVarPath := notifyUri + "/notify"
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}
	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0]	// use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = &pfdChangeNotifications

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:	localVarBody,
		ErrorStatus:	localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 200:
		err = openapi.Deserialize(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
		}
		return localVarReturnValue, localVarHTTPResponse, nil
	case 204:
		return localVarReturnValue, localVarHTTPResponse, nil
	case 400:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 401:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 403:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 404:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 411:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 413:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 415:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 429:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 500:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	case 503:
		var v ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarReturnValue, localVarHTTPResponse, nil
		}
		apiError.ErrorModel = v
		return localVarReturnValue, localVarHTTPResponse, nil
	default:
		return localVarReturnValue, localVarHTTPResponse, nil
	}
}
