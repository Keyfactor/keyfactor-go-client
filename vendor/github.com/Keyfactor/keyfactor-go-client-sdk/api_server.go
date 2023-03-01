/*
Keyfactor-v1

This reference serves to document REST-based methods to manage and integrate with Keyfactor. In addition, an embedded interface allows for the execution of calls against the current Keyfactor API instance.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keyfactor_command_client_api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// ServerApiService ServerApi service
type ServerApiService service

type ApiServerAddAccessRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	xKeyfactorRequestedWith *string
	serverAccess *ModelsSSHAccessServerAccessRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerAddAccessRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerAddAccessRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Logons and users to be applied to the existing server
func (r ApiServerAddAccessRequest) ServerAccess(serverAccess ModelsSSHAccessServerAccessRequest) ApiServerAddAccessRequest {
	r.serverAccess = &serverAccess
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerAddAccessRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerAddAccessRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiServerAddAccessRequest) Execute() (*ModelsSSHAccessServerAccessResponse, *http.Response, error) {
	return r.ApiService.ServerAddAccessExecute(r)
}

/*
ServerAddAccess Updates logons and users with access to those logons for an existing server

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerAddAccessRequest
*/
func (a *ServerApiService) ServerAddAccess(ctx context.Context) ApiServerAddAccessRequest {
	return ApiServerAddAccessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsSSHAccessServerAccessResponse
func (a *ServerApiService) ServerAddAccessExecute(r ApiServerAddAccessRequest) (*ModelsSSHAccessServerAccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsSSHAccessServerAccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerAddAccess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers/Access"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.serverAccess == nil {
		return localVarReturnValue, nil, reportError("serverAccess is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	// body params
	localVarPostBody = r.serverAccess
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServerCreateServerRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	xKeyfactorRequestedWith *string
	creationRequest *ModelsSSHServersServerCreationRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerCreateServerRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerCreateServerRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Server properties to be applied to the newserver
func (r ApiServerCreateServerRequest) CreationRequest(creationRequest ModelsSSHServersServerCreationRequest) ApiServerCreateServerRequest {
	r.creationRequest = &creationRequest
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerCreateServerRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerCreateServerRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiServerCreateServerRequest) Execute() (*ModelsSSHServersServerResponse, *http.Response, error) {
	return r.ApiService.ServerCreateServerExecute(r)
}

/*
ServerCreateServer Creates a server with the provided properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerCreateServerRequest
*/
func (a *ServerApiService) ServerCreateServer(ctx context.Context) ApiServerCreateServerRequest {
	return ApiServerCreateServerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsSSHServersServerResponse
func (a *ServerApiService) ServerCreateServerExecute(r ApiServerCreateServerRequest) (*ModelsSSHServersServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsSSHServersServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerCreateServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.creationRequest == nil {
		return localVarReturnValue, nil, reportError("creationRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	// body params
	localVarPostBody = r.creationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServerDeleteRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	id int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerDeleteRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerDeleteRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerDeleteRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerDeleteRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiServerDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ServerDeleteExecute(r)
}

/*
ServerDelete Deletes a Server associated with the provided identifier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Keyfactor identifer of the Server to be deleted
 @return ApiServerDeleteRequest
*/
func (a *ServerApiService) ServerDelete(ctx context.Context, id int32) ApiServerDeleteRequest {
	return ApiServerDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ServerApiService) ServerDeleteExecute(r ApiServerDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiServerGetRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	id int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerGetRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerGetRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerGetRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerGetRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiServerGetRequest) Execute() (*ModelsSSHServersServerResponse, *http.Response, error) {
	return r.ApiService.ServerGetExecute(r)
}

/*
ServerGet Returns a Server associated with the provided identifier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Keyfactor identifier of the Server
 @return ApiServerGetRequest
*/
func (a *ServerApiService) ServerGet(ctx context.Context, id int32) ApiServerGetRequest {
	return ApiServerGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsSSHServersServerResponse
func (a *ServerApiService) ServerGetExecute(r ApiServerGetRequest) (*ModelsSSHServersServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsSSHServersServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServerGetAccessRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	id int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerGetAccessRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerGetAccessRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerGetAccessRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerGetAccessRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiServerGetAccessRequest) Execute() (*ModelsSSHAccessServerAccessResponse, *http.Response, error) {
	return r.ApiService.ServerGetAccessExecute(r)
}

/*
ServerGetAccess Retrieves logons and users with access to those logons for an existing server

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Id of the existing server
 @return ApiServerGetAccessRequest
*/
func (a *ServerApiService) ServerGetAccess(ctx context.Context, id int32) ApiServerGetAccessRequest {
	return ApiServerGetAccessRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsSSHAccessServerAccessResponse
func (a *ServerApiService) ServerGetAccessExecute(r ApiServerGetAccessRequest) (*ModelsSSHAccessServerAccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsSSHAccessServerAccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerGetAccess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers/Access/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServerQueryServersRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
	pqQueryString *string
	pqPageReturned *int32
	pqReturnLimit *int32
	pqSortField *string
	pqSortAscending *int32
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerQueryServersRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerQueryServersRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerQueryServersRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerQueryServersRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

// Contents of the query (ex: field1 -eq value1 AND field2 -gt value2)
func (r ApiServerQueryServersRequest) PqQueryString(pqQueryString string) ApiServerQueryServersRequest {
	r.pqQueryString = &pqQueryString
	return r
}

// The current page within the result set to be returned
func (r ApiServerQueryServersRequest) PqPageReturned(pqPageReturned int32) ApiServerQueryServersRequest {
	r.pqPageReturned = &pqPageReturned
	return r
}

// Maximum number of records to be returned in a single call
func (r ApiServerQueryServersRequest) PqReturnLimit(pqReturnLimit int32) ApiServerQueryServersRequest {
	r.pqReturnLimit = &pqReturnLimit
	return r
}

// Field by which the results should be sorted (view results via Management Portal for sortable columns)
func (r ApiServerQueryServersRequest) PqSortField(pqSortField string) ApiServerQueryServersRequest {
	r.pqSortField = &pqSortField
	return r
}

// Field sort direction [0&#x3D;ascending, 1&#x3D;descending]
func (r ApiServerQueryServersRequest) PqSortAscending(pqSortAscending int32) ApiServerQueryServersRequest {
	r.pqSortAscending = &pqSortAscending
	return r
}

func (r ApiServerQueryServersRequest) Execute() ([]ModelsSSHServersServerResponse, *http.Response, error) {
	return r.ApiService.ServerQueryServersExecute(r)
}

/*
ServerQueryServers Returns all servers according to the provided filter parameters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerQueryServersRequest
*/
func (a *ServerApiService) ServerQueryServers(ctx context.Context) ApiServerQueryServersRequest {
	return ApiServerQueryServersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ModelsSSHServersServerResponse
func (a *ServerApiService) ServerQueryServersExecute(r ApiServerQueryServersRequest) ([]ModelsSSHServersServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelsSSHServersServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerQueryServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	if r.pqQueryString != nil {
		parameterAddToQuery(localVarQueryParams, "pq.queryString", r.pqQueryString, "")
	}
	if r.pqPageReturned != nil {
		parameterAddToQuery(localVarQueryParams, "pq.pageReturned", r.pqPageReturned, "")
	}
	if r.pqReturnLimit != nil {
		parameterAddToQuery(localVarQueryParams, "pq.returnLimit", r.pqReturnLimit, "")
	}
	if r.pqSortField != nil {
		parameterAddToQuery(localVarQueryParams, "pq.sortField", r.pqSortField, "")
	}
	if r.pqSortAscending != nil {
		parameterAddToQuery(localVarQueryParams, "pq.sortAscending", r.pqSortAscending, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServerRemoveAccessRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	xKeyfactorRequestedWith *string
	serverAccess *ModelsSSHAccessServerAccessRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerRemoveAccessRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerRemoveAccessRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Logons and users to be removed from the existing server
func (r ApiServerRemoveAccessRequest) ServerAccess(serverAccess ModelsSSHAccessServerAccessRequest) ApiServerRemoveAccessRequest {
	r.serverAccess = &serverAccess
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerRemoveAccessRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerRemoveAccessRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiServerRemoveAccessRequest) Execute() (*ModelsSSHAccessServerAccessResponse, *http.Response, error) {
	return r.ApiService.ServerRemoveAccessExecute(r)
}

/*
ServerRemoveAccess Updates logons and users with access to those logons for an existing server

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerRemoveAccessRequest
*/
func (a *ServerApiService) ServerRemoveAccess(ctx context.Context) ApiServerRemoveAccessRequest {
	return ApiServerRemoveAccessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsSSHAccessServerAccessResponse
func (a *ServerApiService) ServerRemoveAccessExecute(r ApiServerRemoveAccessRequest) (*ModelsSSHAccessServerAccessResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsSSHAccessServerAccessResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerRemoveAccess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers/Access"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.serverAccess == nil {
		return localVarReturnValue, nil, reportError("serverAccess is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	// body params
	localVarPostBody = r.serverAccess
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiServerUpdateServerRequest struct {
	ctx context.Context
	ApiService *ServerApiService
	xKeyfactorRequestedWith *string
	updateRequest *ModelsSSHServersServerUpdateRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiServerUpdateServerRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiServerUpdateServerRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Server properties to be applied to the existing server
func (r ApiServerUpdateServerRequest) UpdateRequest(updateRequest ModelsSSHServersServerUpdateRequest) ApiServerUpdateServerRequest {
	r.updateRequest = &updateRequest
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiServerUpdateServerRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiServerUpdateServerRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiServerUpdateServerRequest) Execute() (*ModelsSSHServersServerResponse, *http.Response, error) {
	return r.ApiService.ServerUpdateServerExecute(r)
}

/*
ServerUpdateServer Updates an existing server with the provided properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerUpdateServerRequest
*/
func (a *ServerApiService) ServerUpdateServer(ctx context.Context) ApiServerUpdateServerRequest {
	return ApiServerUpdateServerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsSSHServersServerResponse
func (a *ServerApiService) ServerUpdateServerExecute(r ApiServerUpdateServerRequest) (*ModelsSSHServersServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsSSHServersServerResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServerApiService.ServerUpdateServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/SSH/Servers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.updateRequest == nil {
		return localVarReturnValue, nil, reportError("updateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/xml", "text/xml", "application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/json", "application/xml", "text/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xKeyfactorApiVersion != nil {
		parameterAddToQuery(localVarHeaderParams, "x-keyfactor-api-version", r.xKeyfactorApiVersion, "")
	}
	parameterAddToQuery(localVarHeaderParams, "x-keyfactor-requested-with", r.xKeyfactorRequestedWith, "")
	// body params
	localVarPostBody = r.updateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
