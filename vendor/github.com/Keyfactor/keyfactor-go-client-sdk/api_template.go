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


// TemplateApiService TemplateApi service
type TemplateApiService service

type ApiTemplateGetGlobalSettingsRequest struct {
	ctx context.Context
	ApiService *TemplateApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiTemplateGetGlobalSettingsRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiTemplateGetGlobalSettingsRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiTemplateGetGlobalSettingsRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiTemplateGetGlobalSettingsRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiTemplateGetGlobalSettingsRequest) Execute() (*KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse, *http.Response, error) {
	return r.ApiService.TemplateGetGlobalSettingsExecute(r)
}

/*
TemplateGetGlobalSettings Gets the global template settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTemplateGetGlobalSettingsRequest
*/
func (a *TemplateApiService) TemplateGetGlobalSettings(ctx context.Context) ApiTemplateGetGlobalSettingsRequest {
	return ApiTemplateGetGlobalSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse
func (a *TemplateApiService) TemplateGetGlobalSettingsExecute(r ApiTemplateGetGlobalSettingsRequest) (*KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateApiService.TemplateGetGlobalSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Templates/Settings"

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

type ApiTemplateGetTemplateRequest struct {
	ctx context.Context
	ApiService *TemplateApiService
	id int32
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiTemplateGetTemplateRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiTemplateGetTemplateRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiTemplateGetTemplateRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiTemplateGetTemplateRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiTemplateGetTemplateRequest) Execute() (*ModelsTemplateRetrievalResponse, *http.Response, error) {
	return r.ApiService.TemplateGetTemplateExecute(r)
}

/*
TemplateGetTemplate Returns the certificate template associated with the provided id

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Keyfactor identifier of the certificate template
 @return ApiTemplateGetTemplateRequest
*/
func (a *TemplateApiService) TemplateGetTemplate(ctx context.Context, id int32) ApiTemplateGetTemplateRequest {
	return ApiTemplateGetTemplateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelsTemplateRetrievalResponse
func (a *TemplateApiService) TemplateGetTemplateExecute(r ApiTemplateGetTemplateRequest) (*ModelsTemplateRetrievalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsTemplateRetrievalResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateApiService.TemplateGetTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Templates/{id}"
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

type ApiTemplateGetTemplatesRequest struct {
	ctx context.Context
	ApiService *TemplateApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
	sqQueryString *string
	sqPageReturned *int32
	sqReturnLimit *int32
	sqSortField *string
	sqSortAscending *int32
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiTemplateGetTemplatesRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiTemplateGetTemplatesRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiTemplateGetTemplatesRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiTemplateGetTemplatesRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

// Contents of the query (ex: field1 -eq value1 AND field2 -gt value2)
func (r ApiTemplateGetTemplatesRequest) SqQueryString(sqQueryString string) ApiTemplateGetTemplatesRequest {
	r.sqQueryString = &sqQueryString
	return r
}

// The current page within the result set to be returned
func (r ApiTemplateGetTemplatesRequest) SqPageReturned(sqPageReturned int32) ApiTemplateGetTemplatesRequest {
	r.sqPageReturned = &sqPageReturned
	return r
}

// Maximum number of records to be returned in a single call
func (r ApiTemplateGetTemplatesRequest) SqReturnLimit(sqReturnLimit int32) ApiTemplateGetTemplatesRequest {
	r.sqReturnLimit = &sqReturnLimit
	return r
}

// Field by which the results should be sorted (view results via Management Portal for sortable columns)
func (r ApiTemplateGetTemplatesRequest) SqSortField(sqSortField string) ApiTemplateGetTemplatesRequest {
	r.sqSortField = &sqSortField
	return r
}

// Field sort direction [0&#x3D;ascending, 1&#x3D;descending]
func (r ApiTemplateGetTemplatesRequest) SqSortAscending(sqSortAscending int32) ApiTemplateGetTemplatesRequest {
	r.sqSortAscending = &sqSortAscending
	return r
}

func (r ApiTemplateGetTemplatesRequest) Execute() ([]ModelsTemplateCollectionRetrievalResponse, *http.Response, error) {
	return r.ApiService.TemplateGetTemplatesExecute(r)
}

/*
TemplateGetTemplates Returns all certificate templates according to the provided filter and output parameters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTemplateGetTemplatesRequest
*/
func (a *TemplateApiService) TemplateGetTemplates(ctx context.Context) ApiTemplateGetTemplatesRequest {
	return ApiTemplateGetTemplatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ModelsTemplateCollectionRetrievalResponse
func (a *TemplateApiService) TemplateGetTemplatesExecute(r ApiTemplateGetTemplatesRequest) ([]ModelsTemplateCollectionRetrievalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelsTemplateCollectionRetrievalResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateApiService.TemplateGetTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}

	if r.sqQueryString != nil {
		parameterAddToQuery(localVarQueryParams, "sq.queryString", r.sqQueryString, "")
	}
	if r.sqPageReturned != nil {
		parameterAddToQuery(localVarQueryParams, "sq.pageReturned", r.sqPageReturned, "")
	}
	if r.sqReturnLimit != nil {
		parameterAddToQuery(localVarQueryParams, "sq.returnLimit", r.sqReturnLimit, "")
	}
	if r.sqSortField != nil {
		parameterAddToQuery(localVarQueryParams, "sq.sortField", r.sqSortField, "")
	}
	if r.sqSortAscending != nil {
		parameterAddToQuery(localVarQueryParams, "sq.sortAscending", r.sqSortAscending, "")
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

type ApiTemplateGetValidSubjectPartsRequest struct {
	ctx context.Context
	ApiService *TemplateApiService
	xKeyfactorRequestedWith *string
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiTemplateGetValidSubjectPartsRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiTemplateGetValidSubjectPartsRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiTemplateGetValidSubjectPartsRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiTemplateGetValidSubjectPartsRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiTemplateGetValidSubjectPartsRequest) Execute() ([]KeyfactorApiModelsTemplatesValidSubjectPartResponse, *http.Response, error) {
	return r.ApiService.TemplateGetValidSubjectPartsExecute(r)
}

/*
TemplateGetValidSubjectParts Method for TemplateGetValidSubjectParts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTemplateGetValidSubjectPartsRequest
*/
func (a *TemplateApiService) TemplateGetValidSubjectParts(ctx context.Context) ApiTemplateGetValidSubjectPartsRequest {
	return ApiTemplateGetValidSubjectPartsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []KeyfactorApiModelsTemplatesValidSubjectPartResponse
func (a *TemplateApiService) TemplateGetValidSubjectPartsExecute(r ApiTemplateGetValidSubjectPartsRequest) ([]KeyfactorApiModelsTemplatesValidSubjectPartResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []KeyfactorApiModelsTemplatesValidSubjectPartResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateApiService.TemplateGetValidSubjectParts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Templates/SubjectParts"

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

type ApiTemplateImportRequest struct {
	ctx context.Context
	ApiService *TemplateApiService
	xKeyfactorRequestedWith *string
	configurationTenantRequest *KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiTemplateImportRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiTemplateImportRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Configuration tenant to import from
func (r ApiTemplateImportRequest) ConfigurationTenantRequest(configurationTenantRequest KeyfactorApiModelsConfigurationTenantConfigurationTenantRequest) ApiTemplateImportRequest {
	r.configurationTenantRequest = &configurationTenantRequest
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiTemplateImportRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiTemplateImportRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiTemplateImportRequest) Execute() (*http.Response, error) {
	return r.ApiService.TemplateImportExecute(r)
}

/*
TemplateImport Imports templates from the provided configuration tenant

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTemplateImportRequest
*/
func (a *TemplateApiService) TemplateImport(ctx context.Context) ApiTemplateImportRequest {
	return ApiTemplateImportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TemplateApiService) TemplateImportExecute(r ApiTemplateImportRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateApiService.TemplateImport")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Templates/Import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.configurationTenantRequest == nil {
		return nil, reportError("configurationTenantRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.configurationTenantRequest
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

type ApiTemplateUpdateGlobalSettingsRequest struct {
	ctx context.Context
	ApiService *TemplateApiService
	xKeyfactorRequestedWith *string
	settings *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiTemplateUpdateGlobalSettingsRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiTemplateUpdateGlobalSettingsRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// The new global template settings.
func (r ApiTemplateUpdateGlobalSettingsRequest) Settings(settings KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) ApiTemplateUpdateGlobalSettingsRequest {
	r.settings = &settings
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiTemplateUpdateGlobalSettingsRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiTemplateUpdateGlobalSettingsRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiTemplateUpdateGlobalSettingsRequest) Execute() (*KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse, *http.Response, error) {
	return r.ApiService.TemplateUpdateGlobalSettingsExecute(r)
}

/*
TemplateUpdateGlobalSettings Replaces the existing global template settings.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTemplateUpdateGlobalSettingsRequest
*/
func (a *TemplateApiService) TemplateUpdateGlobalSettings(ctx context.Context) ApiTemplateUpdateGlobalSettingsRequest {
	return ApiTemplateUpdateGlobalSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse
func (a *TemplateApiService) TemplateUpdateGlobalSettingsExecute(r ApiTemplateUpdateGlobalSettingsRequest) (*KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyfactorApiModelsTemplatesGlobalTemplateSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateApiService.TemplateUpdateGlobalSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Templates/Settings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.settings == nil {
		return localVarReturnValue, nil, reportError("settings is required and must be specified")
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
	localVarPostBody = r.settings
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

type ApiTemplateUpdateTemplateRequest struct {
	ctx context.Context
	ApiService *TemplateApiService
	xKeyfactorRequestedWith *string
	template *ModelsTemplateUpdateRequest
	xKeyfactorApiVersion *string
}

// Type of the request [XMLHttpRequest, APIClient]
func (r ApiTemplateUpdateTemplateRequest) XKeyfactorRequestedWith(xKeyfactorRequestedWith string) ApiTemplateUpdateTemplateRequest {
	r.xKeyfactorRequestedWith = &xKeyfactorRequestedWith
	return r
}

// Properties of the certificate template to be updated
func (r ApiTemplateUpdateTemplateRequest) Template(template ModelsTemplateUpdateRequest) ApiTemplateUpdateTemplateRequest {
	r.template = &template
	return r
}

// Desired version of the api, if not provided defaults to v1
func (r ApiTemplateUpdateTemplateRequest) XKeyfactorApiVersion(xKeyfactorApiVersion string) ApiTemplateUpdateTemplateRequest {
	r.xKeyfactorApiVersion = &xKeyfactorApiVersion
	return r
}

func (r ApiTemplateUpdateTemplateRequest) Execute() (*ModelsTemplateRetrievalResponse, *http.Response, error) {
	return r.ApiService.TemplateUpdateTemplateExecute(r)
}

/*
TemplateUpdateTemplate Updates a certificate template according to the provided properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTemplateUpdateTemplateRequest
*/
func (a *TemplateApiService) TemplateUpdateTemplate(ctx context.Context) ApiTemplateUpdateTemplateRequest {
	return ApiTemplateUpdateTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelsTemplateRetrievalResponse
func (a *TemplateApiService) TemplateUpdateTemplateExecute(r ApiTemplateUpdateTemplateRequest) (*ModelsTemplateRetrievalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelsTemplateRetrievalResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateApiService.TemplateUpdateTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xKeyfactorRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xKeyfactorRequestedWith is required and must be specified")
	}
	if r.template == nil {
		return localVarReturnValue, nil, reportError("template is required and must be specified")
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
	localVarPostBody = r.template
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
