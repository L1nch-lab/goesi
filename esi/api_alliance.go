/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 1.30
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"fmt"

	"github.com/antihax/goesi/optional"
)

// Linger please
var (
	_ context.Context
)

type AllianceApiService service

/*
AllianceApiService List all alliances
List all active player alliances  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetAlliancesOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "IfNoneMatch" (optional.String) -  ETag from a previous request. A 304 will be returned if this matches the current ETag

@return []int32
*/

type GetAlliancesOpts struct {
	Datasource	optional.String
	IfNoneMatch	optional.String
}

func (a *AllianceApiService) GetAlliances(ctx context.Context, localVarOptionals *GetAlliancesOpts) ([]int32, *http.Response, error) {
	var (
		localVarHttpMethod	= strings.ToUpper("Get")
		localVarPostBody	interface{}
		localVarFileName	string
		localVarFileBytes	[]byte
		localVarReturnValue	[]int32
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/alliances/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 400 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("content-type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 400 {
		newErr := GenericSwaggerError{
			body:	localVarBody,
			error:	localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []int32
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 420 {
			var v ErrorLimited
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 504 {
			var v GatewayTimeout
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AllianceApiService Get alliance information
Public information about an alliance  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param allianceId An EVE alliance ID
 * @param optional nil or *GetAlliancesAllianceIdOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "IfNoneMatch" (optional.String) -  ETag from a previous request. A 304 will be returned if this matches the current ETag

@return GetAlliancesAllianceIdOk
*/

type GetAlliancesAllianceIdOpts struct {
	Datasource	optional.String
	IfNoneMatch	optional.String
}

func (a *AllianceApiService) GetAlliancesAllianceId(ctx context.Context, allianceId int32, localVarOptionals *GetAlliancesAllianceIdOpts) (GetAlliancesAllianceIdOk, *http.Response, error) {
	var (
		localVarHttpMethod	= strings.ToUpper("Get")
		localVarPostBody	interface{}
		localVarFileName	string
		localVarFileBytes	[]byte
		localVarReturnValue	GetAlliancesAllianceIdOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v3/alliances/{alliance_id}/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if allianceId < 1 {
		return localVarReturnValue, nil, reportError("allianceId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 400 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("content-type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 400 {
		newErr := GenericSwaggerError{
			body:	localVarBody,
			error:	localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v GetAlliancesAllianceIdOk
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 404 {
			var v GetAlliancesAllianceIdNotFound
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 420 {
			var v ErrorLimited
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 504 {
			var v GatewayTimeout
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AllianceApiService List alliance&#39;s corporations
List all current member corporations of an alliance  ---  This route is cached for up to 3600 seconds
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param allianceId An EVE alliance ID
 * @param optional nil or *GetAlliancesAllianceIdCorporationsOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "IfNoneMatch" (optional.String) -  ETag from a previous request. A 304 will be returned if this matches the current ETag

@return []int32
*/

type GetAlliancesAllianceIdCorporationsOpts struct {
	Datasource	optional.String
	IfNoneMatch	optional.String
}

func (a *AllianceApiService) GetAlliancesAllianceIdCorporations(ctx context.Context, allianceId int32, localVarOptionals *GetAlliancesAllianceIdCorporationsOpts) ([]int32, *http.Response, error) {
	var (
		localVarHttpMethod	= strings.ToUpper("Get")
		localVarPostBody	interface{}
		localVarFileName	string
		localVarFileBytes	[]byte
		localVarReturnValue	[]int32
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/alliances/{alliance_id}/corporations/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if allianceId < 1 {
		return localVarReturnValue, nil, reportError("allianceId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 400 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("content-type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 400 {
		newErr := GenericSwaggerError{
			body:	localVarBody,
			error:	localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []int32
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 420 {
			var v ErrorLimited
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 504 {
			var v GatewayTimeout
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AllianceApiService Get alliance icon
Get the icon urls for a alliance  ---  This route expires daily at 11:05
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param allianceId An EVE alliance ID
 * @param optional nil or *GetAlliancesAllianceIdIconsOpts - Optional Parameters:
     * @param "Datasource" (optional.String) -  The server name you would like data from
     * @param "IfNoneMatch" (optional.String) -  ETag from a previous request. A 304 will be returned if this matches the current ETag

@return GetAlliancesAllianceIdIconsOk
*/

type GetAlliancesAllianceIdIconsOpts struct {
	Datasource	optional.String
	IfNoneMatch	optional.String
}

func (a *AllianceApiService) GetAlliancesAllianceIdIcons(ctx context.Context, allianceId int32, localVarOptionals *GetAlliancesAllianceIdIconsOpts) (GetAlliancesAllianceIdIconsOk, *http.Response, error) {
	var (
		localVarHttpMethod	= strings.ToUpper("Get")
		localVarPostBody	interface{}
		localVarFileName	string
		localVarFileBytes	[]byte
		localVarReturnValue	GetAlliancesAllianceIdIconsOk
	)

	// create path and map variables
	localVarPath := a.client.basePath + "/v1/alliances/{alliance_id}/icons/"
	localVarPath = strings.Replace(localVarPath, "{"+"alliance_id"+"}", fmt.Sprintf("%v", allianceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if allianceId < 1 {
		return localVarReturnValue, nil, reportError("allianceId must be greater than 1")
	}

	if localVarOptionals != nil && localVarOptionals.Datasource.IsSet() {
		localVarQueryParams.Add("datasource", parameterToString(localVarOptionals.Datasource.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.IfNoneMatch.IsSet() {
		localVarHeaderParams["If-None-Match"] = parameterToString(localVarOptionals.IfNoneMatch.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 400 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("content-type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 400 {
		newErr := GenericSwaggerError{
			body:	localVarBody,
			error:	localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v GetAlliancesAllianceIdIconsOk
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 400 {
			var v BadRequest
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 404 {
			var v GetAlliancesAllianceIdIconsNotFound
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 420 {
			var v ErrorLimited
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 500 {
			var v InternalServerError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 503 {
			var v ServiceUnavailable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 504 {
			var v GatewayTimeout
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("content-type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
