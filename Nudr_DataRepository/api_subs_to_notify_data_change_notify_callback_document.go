package Nudr_DataRepository

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
)

// Linger please
var (
	_ context.Context
)

type DataChangeNotifyCallbackDocumentApiService service

func (a *DataChangeNotifyCallbackDocumentApiService) OnDataChangeNotify(ctx context.Context, onDataChangeNotifyUrl string, request models.DataChangeNotify) (*http.Response, error) {
	var (
		localVarHttpMethod	= strings.ToUpper("Post")
		localVarPostBody	interface{}
		localVarFormFileName	string
		localVarFileName	string
		localVarFileBytes	[]byte
	)

	// create path and map variables
	localVarPath := onDataChangeNotifyUrl
	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHttpContentTypes := []string{"application/json"}
	localVarHeaderParams["Content-Type"] = localVarHttpContentTypes[0]	// use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHttpHeaderAccept := openapi.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}

	// body params
	localVarPostBody = &request

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:	localVarBody,
		ErrorStatus:	localVarHttpResponse.Status,
	}
	switch localVarHttpResponse.StatusCode {

	case 204:
		return localVarHttpResponse, err
	case 400:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 404:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 411:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 413:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 415:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 429:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 500:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	case 503:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHttpResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHttpResponse, apiError
	default:
		return localVarHttpResponse, openapi.ReportError("%d is not a valid status code in OnDataChangeNotify", localVarHttpResponse.StatusCode)
	}
}
