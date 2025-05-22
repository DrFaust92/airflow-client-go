/*
Airflow API (Stable)

# Overview  To facilitate management, Apache Airflow supports a range of REST API endpoints across its objects. This section provides an overview of the API design, methods, and supported use cases.  Most of the endpoints accept `JSON` as input and return `JSON` responses. This means that you must usually add the following headers to your request: ``` Content-type: application/json Accept: application/json ```  ## Resources  The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its endpoint's corresponding resource. The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.  Resource names are used as part of endpoint URLs, as well as in API parameters and responses.  ## CRUD Operations  The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources. You can review the standards for these operations and their standard parameters below.  Some endpoints have special behavior as exceptions.  ### Create  To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata in the request body. The response returns a `201 Created` response code upon success with the resource's metadata, including its internal `id`, in the response body.  ### Read  The HTTP `GET` request can be used to read a resource or to list a number of resources.  A resource's `id` can be submitted in the request parameters to read a specific resource. The response usually returns a `200 OK` response code upon success, with the resource's metadata in the response body.  If a `GET` request does not include a specific resource `id`, it is treated as a list request. The response usually returns a `200 OK` response code upon success, with an object containing a list of resources' metadata in the response body.  When reading resources, some common query parameters are usually available. e.g.: ``` v1/connections?limit=25&offset=25 ```  |Query Parameter|Type|Description| |---------------|----|-----------| |limit|integer|Maximum number of objects to fetch. Usually 25 by default| |offset|integer|Offset after which to start returning objects. For use with limit query parameter.|  ### Update  Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request, with the fields to modify in the request body. The response usually returns a `200 OK` response code upon success, with information about the modified resource in the response body.  ### Delete  Deleting a resource requires the resource `id` and is typically executed via an HTTP `DELETE` request. The response usually returns a `204 No Content` response code upon success.  ## Conventions  - Resource names are plural and expressed in camelCase. - Names are consistent between URL parameter name and field name.  - Field names are in snake_case. ```json {     \"description\": \"string\",     \"name\": \"string\",     \"occupied_slots\": 0,     \"open_slots\": 0     \"queued_slots\": 0,     \"running_slots\": 0,     \"scheduled_slots\": 0,     \"slots\": 0, } ```  ### Update Mask  Update mask is available as a query parameter in patch endpoints. It is used to notify the API which fields you want to update. Using `update_mask` makes it easier to update objects by helping the server know which fields to update in an object instead of updating all fields. The update request ignores any fields that aren't specified in the field mask, leaving them with their current values.  Example: ```   resource = request.get('/resource/my-id').json()   resource['my_field'] = 'new-value'   request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource)) ```  ## Versioning and Endpoint Lifecycle  - API versioning is not synchronized to specific releases of the Apache Airflow. - APIs are designed to be backward compatible. - Any changes to the API will first go through a deprecation phase.  # Trying the API  You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/), [Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test the Apache Airflow API.  Note that you will need to pass credentials data.  For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used: ```bash curl -X PATCH 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\ -H 'Content-Type: application/json' \\ --user \"username:password\" \\ -d '{     \"is_paused\": true }' ```  Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/), it is possible to import the API specifications directly:  1. Download the API specification by clicking the **Download** button at the top of this document 2. Import the JSON specification in the graphical tool of your choice.   - In *Postman*, you can click the **import** button at the top   - With *Insomnia*, you can just drag-and-drop the file on the UI  Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on the **Code** button.  ## Enabling CORS  [Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a browser security feature that restricts HTTP requests that are initiated from scripts running in the browser.  For details on enabling/configuring CORS, see [Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Authentication  To be able to meet the requirements of many organizations, Airflow supports many authentication methods, and it is even possible to add your own method.  If you want to check which auth backend is currently set, you can use `airflow config get-value api auth_backends` command as in the example below. ```bash $ airflow config get-value api auth_backends airflow.api.auth.backend.basic_auth ``` The default is to deny all requests.  For details on configuring the authentication, see [API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Errors  We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807) also known as Problem Details for HTTP APIs. As with our normal API responses, your client must be prepared to gracefully handle additional members of the response.  ## Unauthenticated  This indicates that the request has not been applied because it lacks valid authentication credentials for the target resource. Please check that you have valid credentials.  ## PermissionDenied  This response means that the server understood the request but refuses to authorize it because it lacks sufficient rights to the resource. It happens when you do not have the necessary permission to execute the action you performed. You need to get the appropriate permissions in other to resolve this error.  ## BadRequest  This response means that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.  ## NotFound  This client error response indicates that the server cannot find the requested resource.  ## MethodNotAllowed  Indicates that the request method is known by the server but is not supported by the target resource.  ## NotAcceptable  The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request, and the server is unwilling to supply a default representation.  ## AlreadyExists  The request could not be completed due to a conflict with the current state of the target resource, e.g. the resource it tries to create already exists.  ## Unknown  This means that the server encountered an unexpected condition that prevented it from fulfilling the request. 

API version: 2.10.5
Contact: dev@airflow.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airflow

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
	"reflect"
)


// TaskInstanceAPIService TaskInstanceAPI service
type TaskInstanceAPIService service

type ApiGetExtraLinksRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	mapIndex *int32
}

// Filter on map index for mapped task.
func (r ApiGetExtraLinksRequest) MapIndex(mapIndex int32) ApiGetExtraLinksRequest {
	r.mapIndex = &mapIndex
	return r
}

func (r ApiGetExtraLinksRequest) Execute() (*ExtraLinkCollection, *http.Response, error) {
	return r.ApiService.GetExtraLinksExecute(r)
}

/*
GetExtraLinks List extra links

List extra links for task instance.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return ApiGetExtraLinksRequest
*/
func (a *TaskInstanceAPIService) GetExtraLinks(ctx context.Context, dagId string, dagRunId string, taskId string) ApiGetExtraLinksRequest {
	return ApiGetExtraLinksRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return ExtraLinkCollection
func (a *TaskInstanceAPIService) GetExtraLinksExecute(r ApiGetExtraLinksRequest) (*ExtraLinkCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ExtraLinkCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetExtraLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.mapIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "map_index", r.mapIndex, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetLogRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	taskTryNumber int32
	fullContent *bool
	mapIndex *int32
	token *string
}

// A full content will be returned. By default, only the first fragment will be returned. 
func (r ApiGetLogRequest) FullContent(fullContent bool) ApiGetLogRequest {
	r.fullContent = &fullContent
	return r
}

// Filter on map index for mapped task.
func (r ApiGetLogRequest) MapIndex(mapIndex int32) ApiGetLogRequest {
	r.mapIndex = &mapIndex
	return r
}

// A token that allows you to continue fetching logs. If passed, it will specify the location from which the download should be continued. 
func (r ApiGetLogRequest) Token(token string) ApiGetLogRequest {
	r.token = &token
	return r
}

func (r ApiGetLogRequest) Execute() (*GetLog200Response, *http.Response, error) {
	return r.ApiService.GetLogExecute(r)
}

/*
GetLog Get logs

Get logs for a specific task instance and its try number.
To get log from specific character position, following way of using
URLSafeSerializer can be used.

Example:
```
from itsdangerous.url_safe import URLSafeSerializer

request_url = f"api/v1/dags/{DAG_ID}/dagRuns/{RUN_ID}/taskInstances/{TASK_ID}/logs/1"
key = app.config["SECRET_KEY"]
serializer = URLSafeSerializer(key)
token = serializer.dumps({"log_pos": 10000})

response = self.client.get(
    request_url,
    query_string={"token": token},
    headers={"Accept": "text/plain"},
    environ_overrides={"REMOTE_USER": "test"},
)
continuation_token = response.json["continuation_token"]
    metadata = URLSafeSerializer(key).loads(continuation_token)
    log_pos = metadata["log_pos"]
    end_of_log = metadata["end_of_log"]
```
If log_pos is passed as 10000 like the above example, it renders the logs starting
from char position 10000 to last (not the end as the logs may be tailing behind in
running state). This way pagination can be done with metadata as part of the token.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param taskTryNumber The task try number.
 @return ApiGetLogRequest
*/
func (a *TaskInstanceAPIService) GetLog(ctx context.Context, dagId string, dagRunId string, taskId string, taskTryNumber int32) ApiGetLogRequest {
	return ApiGetLogRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		taskTryNumber: taskTryNumber,
	}
}

// Execute executes the request
//  @return GetLog200Response
func (a *TaskInstanceAPIService) GetLogExecute(r ApiGetLogRequest) (*GetLog200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetLog200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetLog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/logs/{task_try_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_try_number"+"}", url.PathEscape(parameterValueToString(r.taskTryNumber, "taskTryNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fullContent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "full_content", r.fullContent, "form", "")
	}
	if r.mapIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "map_index", r.mapIndex, "form", "")
	}
	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetMappedTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
}

func (r ApiGetMappedTaskInstanceRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.GetMappedTaskInstanceExecute(r)
}

/*
GetMappedTaskInstance Get a mapped task instance

Get details of a mapped task instance.

*New in version 2.3.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return ApiGetMappedTaskInstanceRequest
*/
func (a *TaskInstanceAPIService) GetMappedTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) ApiGetMappedTaskInstanceRequest {
	return ApiGetMappedTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceAPIService) GetMappedTaskInstanceExecute(r ApiGetMappedTaskInstanceRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetMappedTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterValueToString(r.mapIndex, "mapIndex")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetMappedTaskInstanceDependenciesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
}

func (r ApiGetMappedTaskInstanceDependenciesRequest) Execute() (*TaskInstanceDependencyCollection, *http.Response, error) {
	return r.ApiService.GetMappedTaskInstanceDependenciesExecute(r)
}

/*
GetMappedTaskInstanceDependencies Get task dependencies blocking task from getting scheduled.

Get task dependencies blocking task from getting scheduled.

*New in version 2.10.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return ApiGetMappedTaskInstanceDependenciesRequest
*/
func (a *TaskInstanceAPIService) GetMappedTaskInstanceDependencies(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) ApiGetMappedTaskInstanceDependenciesRequest {
	return ApiGetMappedTaskInstanceDependenciesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstanceDependencyCollection
func (a *TaskInstanceAPIService) GetMappedTaskInstanceDependenciesExecute(r ApiGetMappedTaskInstanceDependenciesRequest) (*TaskInstanceDependencyCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceDependencyCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetMappedTaskInstanceDependencies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/dependencies"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterValueToString(r.mapIndex, "mapIndex")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetMappedTaskInstanceTriesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
	limit *int32
	offset *int32
	orderBy *string
}

// The numbers of items to return.
func (r ApiGetMappedTaskInstanceTriesRequest) Limit(limit int32) ApiGetMappedTaskInstanceTriesRequest {
	r.limit = &limit
	return r
}

// The number of items to skip before starting to collect the result set.
func (r ApiGetMappedTaskInstanceTriesRequest) Offset(offset int32) ApiGetMappedTaskInstanceTriesRequest {
	r.offset = &offset
	return r
}

// The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0* 
func (r ApiGetMappedTaskInstanceTriesRequest) OrderBy(orderBy string) ApiGetMappedTaskInstanceTriesRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetMappedTaskInstanceTriesRequest) Execute() (*TaskInstanceHistoryCollection, *http.Response, error) {
	return r.ApiService.GetMappedTaskInstanceTriesExecute(r)
}

/*
GetMappedTaskInstanceTries List mapped task instance tries

Get details of all task instance tries.

*New in version 2.10.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return ApiGetMappedTaskInstanceTriesRequest
*/
func (a *TaskInstanceAPIService) GetMappedTaskInstanceTries(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) ApiGetMappedTaskInstanceTriesRequest {
	return ApiGetMappedTaskInstanceTriesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstanceHistoryCollection
func (a *TaskInstanceAPIService) GetMappedTaskInstanceTriesExecute(r ApiGetMappedTaskInstanceTriesRequest) (*TaskInstanceHistoryCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceHistoryCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetMappedTaskInstanceTries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/tries"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterValueToString(r.mapIndex, "mapIndex")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetMappedTaskInstanceTryDetailsRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
	taskTryNumber int32
}

func (r ApiGetMappedTaskInstanceTryDetailsRequest) Execute() (*TaskInstanceHistory, *http.Response, error) {
	return r.ApiService.GetMappedTaskInstanceTryDetailsExecute(r)
}

/*
GetMappedTaskInstanceTryDetails get mapped taskinstance try

Get details of a mapped task instance try.

*New in version 2.10.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @param taskTryNumber The task try number.
 @return ApiGetMappedTaskInstanceTryDetailsRequest
*/
func (a *TaskInstanceAPIService) GetMappedTaskInstanceTryDetails(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32, taskTryNumber int32) ApiGetMappedTaskInstanceTryDetailsRequest {
	return ApiGetMappedTaskInstanceTryDetailsRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
		taskTryNumber: taskTryNumber,
	}
}

// Execute executes the request
//  @return TaskInstanceHistory
func (a *TaskInstanceAPIService) GetMappedTaskInstanceTryDetailsExecute(r ApiGetMappedTaskInstanceTryDetailsRequest) (*TaskInstanceHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetMappedTaskInstanceTryDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/tries/{task_try_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterValueToString(r.mapIndex, "mapIndex")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_try_number"+"}", url.PathEscape(parameterValueToString(r.taskTryNumber, "taskTryNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetMappedTaskInstancesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	limit *int32
	offset *int32
	executionDateGte *time.Time
	executionDateLte *time.Time
	startDateGte *time.Time
	startDateLte *time.Time
	endDateGte *time.Time
	endDateLte *time.Time
	updatedAtGte *time.Time
	updatedAtLte *time.Time
	durationGte *float32
	durationLte *float32
	state *[]string
	pool *[]string
	queue *[]string
	executor *[]string
	orderBy *string
}

// The numbers of items to return.
func (r ApiGetMappedTaskInstancesRequest) Limit(limit int32) ApiGetMappedTaskInstancesRequest {
	r.limit = &limit
	return r
}

// The number of items to skip before starting to collect the result set.
func (r ApiGetMappedTaskInstancesRequest) Offset(offset int32) ApiGetMappedTaskInstancesRequest {
	r.offset = &offset
	return r
}

// Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period. 
func (r ApiGetMappedTaskInstancesRequest) ExecutionDateGte(executionDateGte time.Time) ApiGetMappedTaskInstancesRequest {
	r.executionDateGte = &executionDateGte
	return r
}

// Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period. 
func (r ApiGetMappedTaskInstancesRequest) ExecutionDateLte(executionDateLte time.Time) ApiGetMappedTaskInstancesRequest {
	r.executionDateLte = &executionDateLte
	return r
}

// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r ApiGetMappedTaskInstancesRequest) StartDateGte(startDateGte time.Time) ApiGetMappedTaskInstancesRequest {
	r.startDateGte = &startDateGte
	return r
}

// Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r ApiGetMappedTaskInstancesRequest) StartDateLte(startDateLte time.Time) ApiGetMappedTaskInstancesRequest {
	r.startDateLte = &startDateLte
	return r
}

// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r ApiGetMappedTaskInstancesRequest) EndDateGte(endDateGte time.Time) ApiGetMappedTaskInstancesRequest {
	r.endDateGte = &endDateGte
	return r
}

// Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r ApiGetMappedTaskInstancesRequest) EndDateLte(endDateLte time.Time) ApiGetMappedTaskInstancesRequest {
	r.endDateLte = &endDateLte
	return r
}

// Returns objects greater or equal the specified date.  This can be combined with updated_at_lte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r ApiGetMappedTaskInstancesRequest) UpdatedAtGte(updatedAtGte time.Time) ApiGetMappedTaskInstancesRequest {
	r.updatedAtGte = &updatedAtGte
	return r
}

// Returns objects less or equal the specified date.  This can be combined with updated_at_gte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r ApiGetMappedTaskInstancesRequest) UpdatedAtLte(updatedAtLte time.Time) ApiGetMappedTaskInstancesRequest {
	r.updatedAtLte = &updatedAtLte
	return r
}

// Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period. 
func (r ApiGetMappedTaskInstancesRequest) DurationGte(durationGte float32) ApiGetMappedTaskInstancesRequest {
	r.durationGte = &durationGte
	return r
}

// Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range. 
func (r ApiGetMappedTaskInstancesRequest) DurationLte(durationLte float32) ApiGetMappedTaskInstancesRequest {
	r.durationLte = &durationLte
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetMappedTaskInstancesRequest) State(state []string) ApiGetMappedTaskInstancesRequest {
	r.state = &state
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetMappedTaskInstancesRequest) Pool(pool []string) ApiGetMappedTaskInstancesRequest {
	r.pool = &pool
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetMappedTaskInstancesRequest) Queue(queue []string) ApiGetMappedTaskInstancesRequest {
	r.queue = &queue
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetMappedTaskInstancesRequest) Executor(executor []string) ApiGetMappedTaskInstancesRequest {
	r.executor = &executor
	return r
}

// The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0* 
func (r ApiGetMappedTaskInstancesRequest) OrderBy(orderBy string) ApiGetMappedTaskInstancesRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetMappedTaskInstancesRequest) Execute() (*TaskInstanceCollection, *http.Response, error) {
	return r.ApiService.GetMappedTaskInstancesExecute(r)
}

/*
GetMappedTaskInstances List mapped task instances

Get details of all mapped task instances.

*New in version 2.3.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return ApiGetMappedTaskInstancesRequest
*/
func (a *TaskInstanceAPIService) GetMappedTaskInstances(ctx context.Context, dagId string, dagRunId string, taskId string) ApiGetMappedTaskInstancesRequest {
	return ApiGetMappedTaskInstancesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstanceCollection
func (a *TaskInstanceAPIService) GetMappedTaskInstancesExecute(r ApiGetMappedTaskInstancesRequest) (*TaskInstanceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetMappedTaskInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/listMapped"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.executionDateGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "execution_date_gte", r.executionDateGte, "form", "")
	}
	if r.executionDateLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "execution_date_lte", r.executionDateLte, "form", "")
	}
	if r.startDateGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_gte", r.startDateGte, "form", "")
	}
	if r.startDateLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_lte", r.startDateLte, "form", "")
	}
	if r.endDateGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_gte", r.endDateGte, "form", "")
	}
	if r.endDateLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_lte", r.endDateLte, "form", "")
	}
	if r.updatedAtGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updated_at_gte", r.updatedAtGte, "form", "")
	}
	if r.updatedAtLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updated_at_lte", r.updatedAtLte, "form", "")
	}
	if r.durationGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration_gte", r.durationGte, "form", "")
	}
	if r.durationLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration_lte", r.durationLte, "form", "")
	}
	if r.state != nil {
		t := *r.state
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "state", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "state", t, "form", "multi")
		}
	}
	if r.pool != nil {
		t := *r.pool
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pool", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pool", t, "form", "multi")
		}
	}
	if r.queue != nil {
		t := *r.queue
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "queue", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "queue", t, "form", "multi")
		}
	}
	if r.executor != nil {
		t := *r.executor
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "executor", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "executor", t, "form", "multi")
		}
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
}

func (r ApiGetTaskInstanceRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.GetTaskInstanceExecute(r)
}

/*
GetTaskInstance Get a task instance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return ApiGetTaskInstanceRequest
*/
func (a *TaskInstanceAPIService) GetTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string) ApiGetTaskInstanceRequest {
	return ApiGetTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceAPIService) GetTaskInstanceExecute(r ApiGetTaskInstanceRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetTaskInstanceDependenciesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
}

func (r ApiGetTaskInstanceDependenciesRequest) Execute() (*TaskInstanceDependencyCollection, *http.Response, error) {
	return r.ApiService.GetTaskInstanceDependenciesExecute(r)
}

/*
GetTaskInstanceDependencies Get task dependencies blocking task from getting scheduled.

Get task dependencies blocking task from getting scheduled.

*New in version 2.10.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return ApiGetTaskInstanceDependenciesRequest
*/
func (a *TaskInstanceAPIService) GetTaskInstanceDependencies(ctx context.Context, dagId string, dagRunId string, taskId string) ApiGetTaskInstanceDependenciesRequest {
	return ApiGetTaskInstanceDependenciesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstanceDependencyCollection
func (a *TaskInstanceAPIService) GetTaskInstanceDependenciesExecute(r ApiGetTaskInstanceDependenciesRequest) (*TaskInstanceDependencyCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceDependencyCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetTaskInstanceDependencies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/dependencies"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetTaskInstanceTriesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	limit *int32
	offset *int32
	orderBy *string
}

// The numbers of items to return.
func (r ApiGetTaskInstanceTriesRequest) Limit(limit int32) ApiGetTaskInstanceTriesRequest {
	r.limit = &limit
	return r
}

// The number of items to skip before starting to collect the result set.
func (r ApiGetTaskInstanceTriesRequest) Offset(offset int32) ApiGetTaskInstanceTriesRequest {
	r.offset = &offset
	return r
}

// The name of the field to order the results by. Prefix a field name with &#x60;-&#x60; to reverse the sort order.  *New in version 2.1.0* 
func (r ApiGetTaskInstanceTriesRequest) OrderBy(orderBy string) ApiGetTaskInstanceTriesRequest {
	r.orderBy = &orderBy
	return r
}

func (r ApiGetTaskInstanceTriesRequest) Execute() (*TaskInstanceHistoryCollection, *http.Response, error) {
	return r.ApiService.GetTaskInstanceTriesExecute(r)
}

/*
GetTaskInstanceTries List task instance tries

Get details of all task instance tries.

*New in version 2.10.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return ApiGetTaskInstanceTriesRequest
*/
func (a *TaskInstanceAPIService) GetTaskInstanceTries(ctx context.Context, dagId string, dagRunId string, taskId string) ApiGetTaskInstanceTriesRequest {
	return ApiGetTaskInstanceTriesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstanceHistoryCollection
func (a *TaskInstanceAPIService) GetTaskInstanceTriesExecute(r ApiGetTaskInstanceTriesRequest) (*TaskInstanceHistoryCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceHistoryCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetTaskInstanceTries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/tries"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetTaskInstanceTryDetailsRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	taskTryNumber int32
}

func (r ApiGetTaskInstanceTryDetailsRequest) Execute() (*TaskInstanceHistory, *http.Response, error) {
	return r.ApiService.GetTaskInstanceTryDetailsExecute(r)
}

/*
GetTaskInstanceTryDetails get taskinstance try

Get details of a task instance try.

*New in version 2.10.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param taskTryNumber The task try number.
 @return ApiGetTaskInstanceTryDetailsRequest
*/
func (a *TaskInstanceAPIService) GetTaskInstanceTryDetails(ctx context.Context, dagId string, dagRunId string, taskId string, taskTryNumber int32) ApiGetTaskInstanceTryDetailsRequest {
	return ApiGetTaskInstanceTryDetailsRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		taskTryNumber: taskTryNumber,
	}
}

// Execute executes the request
//  @return TaskInstanceHistory
func (a *TaskInstanceAPIService) GetTaskInstanceTryDetailsExecute(r ApiGetTaskInstanceTryDetailsRequest) (*TaskInstanceHistory, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceHistory
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetTaskInstanceTryDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/tries/{task_try_number}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_try_number"+"}", url.PathEscape(parameterValueToString(r.taskTryNumber, "taskTryNumber")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetTaskInstancesRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	executionDateGte *time.Time
	executionDateLte *time.Time
	startDateGte *time.Time
	startDateLte *time.Time
	endDateGte *time.Time
	endDateLte *time.Time
	updatedAtGte *time.Time
	updatedAtLte *time.Time
	durationGte *float32
	durationLte *float32
	state *[]string
	pool *[]string
	queue *[]string
	executor *[]string
	limit *int32
	offset *int32
}

// Returns objects greater or equal to the specified date.  This can be combined with execution_date_lte parameter to receive only the selected period. 
func (r ApiGetTaskInstancesRequest) ExecutionDateGte(executionDateGte time.Time) ApiGetTaskInstancesRequest {
	r.executionDateGte = &executionDateGte
	return r
}

// Returns objects less than or equal to the specified date.  This can be combined with execution_date_gte parameter to receive only the selected period. 
func (r ApiGetTaskInstancesRequest) ExecutionDateLte(executionDateLte time.Time) ApiGetTaskInstancesRequest {
	r.executionDateLte = &executionDateLte
	return r
}

// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r ApiGetTaskInstancesRequest) StartDateGte(startDateGte time.Time) ApiGetTaskInstancesRequest {
	r.startDateGte = &startDateGte
	return r
}

// Returns objects less or equal the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r ApiGetTaskInstancesRequest) StartDateLte(startDateLte time.Time) ApiGetTaskInstancesRequest {
	r.startDateLte = &startDateLte
	return r
}

// Returns objects greater or equal the specified date.  This can be combined with start_date_lte parameter to receive only the selected period. 
func (r ApiGetTaskInstancesRequest) EndDateGte(endDateGte time.Time) ApiGetTaskInstancesRequest {
	r.endDateGte = &endDateGte
	return r
}

// Returns objects less than or equal to the specified date.  This can be combined with start_date_gte parameter to receive only the selected period. 
func (r ApiGetTaskInstancesRequest) EndDateLte(endDateLte time.Time) ApiGetTaskInstancesRequest {
	r.endDateLte = &endDateLte
	return r
}

// Returns objects greater or equal the specified date.  This can be combined with updated_at_lte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r ApiGetTaskInstancesRequest) UpdatedAtGte(updatedAtGte time.Time) ApiGetTaskInstancesRequest {
	r.updatedAtGte = &updatedAtGte
	return r
}

// Returns objects less or equal the specified date.  This can be combined with updated_at_gte parameter to receive only the selected period.  *New in version 2.6.0* 
func (r ApiGetTaskInstancesRequest) UpdatedAtLte(updatedAtLte time.Time) ApiGetTaskInstancesRequest {
	r.updatedAtLte = &updatedAtLte
	return r
}

// Returns objects greater than or equal to the specified values.  This can be combined with duration_lte parameter to receive only the selected period. 
func (r ApiGetTaskInstancesRequest) DurationGte(durationGte float32) ApiGetTaskInstancesRequest {
	r.durationGte = &durationGte
	return r
}

// Returns objects less than or equal to the specified values.  This can be combined with duration_gte parameter to receive only the selected range. 
func (r ApiGetTaskInstancesRequest) DurationLte(durationLte float32) ApiGetTaskInstancesRequest {
	r.durationLte = &durationLte
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetTaskInstancesRequest) State(state []string) ApiGetTaskInstancesRequest {
	r.state = &state
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetTaskInstancesRequest) Pool(pool []string) ApiGetTaskInstancesRequest {
	r.pool = &pool
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetTaskInstancesRequest) Queue(queue []string) ApiGetTaskInstancesRequest {
	r.queue = &queue
	return r
}

// The value can be repeated to retrieve multiple matching values (OR condition).
func (r ApiGetTaskInstancesRequest) Executor(executor []string) ApiGetTaskInstancesRequest {
	r.executor = &executor
	return r
}

// The numbers of items to return.
func (r ApiGetTaskInstancesRequest) Limit(limit int32) ApiGetTaskInstancesRequest {
	r.limit = &limit
	return r
}

// The number of items to skip before starting to collect the result set.
func (r ApiGetTaskInstancesRequest) Offset(offset int32) ApiGetTaskInstancesRequest {
	r.offset = &offset
	return r
}

func (r ApiGetTaskInstancesRequest) Execute() (*TaskInstanceCollection, *http.Response, error) {
	return r.ApiService.GetTaskInstancesExecute(r)
}

/*
GetTaskInstances List task instances

This endpoint allows specifying `~` as the dag_id, dag_run_id to retrieve DAG runs for all DAGs and DAG runs.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @return ApiGetTaskInstancesRequest
*/
func (a *TaskInstanceAPIService) GetTaskInstances(ctx context.Context, dagId string, dagRunId string) ApiGetTaskInstancesRequest {
	return ApiGetTaskInstancesRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
	}
}

// Execute executes the request
//  @return TaskInstanceCollection
func (a *TaskInstanceAPIService) GetTaskInstancesExecute(r ApiGetTaskInstancesRequest) (*TaskInstanceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetTaskInstances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.executionDateGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "execution_date_gte", r.executionDateGte, "form", "")
	}
	if r.executionDateLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "execution_date_lte", r.executionDateLte, "form", "")
	}
	if r.startDateGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_gte", r.startDateGte, "form", "")
	}
	if r.startDateLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_date_lte", r.startDateLte, "form", "")
	}
	if r.endDateGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_gte", r.endDateGte, "form", "")
	}
	if r.endDateLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date_lte", r.endDateLte, "form", "")
	}
	if r.updatedAtGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updated_at_gte", r.updatedAtGte, "form", "")
	}
	if r.updatedAtLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updated_at_lte", r.updatedAtLte, "form", "")
	}
	if r.durationGte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration_gte", r.durationGte, "form", "")
	}
	if r.durationLte != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "duration_lte", r.durationLte, "form", "")
	}
	if r.state != nil {
		t := *r.state
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "state", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "state", t, "form", "multi")
		}
	}
	if r.pool != nil {
		t := *r.pool
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "pool", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "pool", t, "form", "multi")
		}
	}
	if r.queue != nil {
		t := *r.queue
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "queue", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "queue", t, "form", "multi")
		}
	}
	if r.executor != nil {
		t := *r.executor
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "executor", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "executor", t, "form", "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiGetTaskInstancesBatchRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	listTaskInstanceForm *ListTaskInstanceForm
}

func (r ApiGetTaskInstancesBatchRequest) ListTaskInstanceForm(listTaskInstanceForm ListTaskInstanceForm) ApiGetTaskInstancesBatchRequest {
	r.listTaskInstanceForm = &listTaskInstanceForm
	return r
}

func (r ApiGetTaskInstancesBatchRequest) Execute() (*TaskInstanceCollection, *http.Response, error) {
	return r.ApiService.GetTaskInstancesBatchExecute(r)
}

/*
GetTaskInstancesBatch List task instances (batch)

List task instances from all DAGs and DAG runs.
This endpoint is a POST to allow filtering across a large number of DAG IDs, where as a GET it would run in to maximum HTTP request URL length limits.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTaskInstancesBatchRequest
*/
func (a *TaskInstanceAPIService) GetTaskInstancesBatch(ctx context.Context) ApiGetTaskInstancesBatchRequest {
	return ApiGetTaskInstancesBatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TaskInstanceCollection
func (a *TaskInstanceAPIService) GetTaskInstancesBatchExecute(r ApiGetTaskInstancesBatchRequest) (*TaskInstanceCollection, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceCollection
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.GetTaskInstancesBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/~/dagRuns/~/taskInstances/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.listTaskInstanceForm == nil {
		return localVarReturnValue, nil, reportError("listTaskInstanceForm is required and must be specified")
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
	// body params
	localVarPostBody = r.listTaskInstanceForm
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPatchMappedTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
	updateTaskInstance *UpdateTaskInstance
}

// Parameters of action
func (r ApiPatchMappedTaskInstanceRequest) UpdateTaskInstance(updateTaskInstance UpdateTaskInstance) ApiPatchMappedTaskInstanceRequest {
	r.updateTaskInstance = &updateTaskInstance
	return r
}

func (r ApiPatchMappedTaskInstanceRequest) Execute() (*TaskInstanceReference, *http.Response, error) {
	return r.ApiService.PatchMappedTaskInstanceExecute(r)
}

/*
PatchMappedTaskInstance Updates the state of a mapped task instance

Updates the state for single mapped task instance.
*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return ApiPatchMappedTaskInstanceRequest
*/
func (a *TaskInstanceAPIService) PatchMappedTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) ApiPatchMappedTaskInstanceRequest {
	return ApiPatchMappedTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstanceReference
func (a *TaskInstanceAPIService) PatchMappedTaskInstanceExecute(r ApiPatchMappedTaskInstanceRequest) (*TaskInstanceReference, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.PatchMappedTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterValueToString(r.mapIndex, "mapIndex")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	// body params
	localVarPostBody = r.updateTaskInstance
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiPatchTaskInstanceRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	updateTaskInstance *UpdateTaskInstance
}

// Parameters of action
func (r ApiPatchTaskInstanceRequest) UpdateTaskInstance(updateTaskInstance UpdateTaskInstance) ApiPatchTaskInstanceRequest {
	r.updateTaskInstance = &updateTaskInstance
	return r
}

func (r ApiPatchTaskInstanceRequest) Execute() (*TaskInstanceReference, *http.Response, error) {
	return r.ApiService.PatchTaskInstanceExecute(r)
}

/*
PatchTaskInstance Updates the state of a task instance

Updates the state for single task instance.
*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return ApiPatchTaskInstanceRequest
*/
func (a *TaskInstanceAPIService) PatchTaskInstance(ctx context.Context, dagId string, dagRunId string, taskId string) ApiPatchTaskInstanceRequest {
	return ApiPatchTaskInstanceRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstanceReference
func (a *TaskInstanceAPIService) PatchTaskInstanceExecute(r ApiPatchTaskInstanceRequest) (*TaskInstanceReference, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstanceReference
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.PatchTaskInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateTaskInstance == nil {
		return localVarReturnValue, nil, reportError("updateTaskInstance is required and must be specified")
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
	// body params
	localVarPostBody = r.updateTaskInstance
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiSetMappedTaskInstanceNoteRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	mapIndex int32
	setTaskInstanceNote *SetTaskInstanceNote
}

// Parameters of set Task Instance note.
func (r ApiSetMappedTaskInstanceNoteRequest) SetTaskInstanceNote(setTaskInstanceNote SetTaskInstanceNote) ApiSetMappedTaskInstanceNoteRequest {
	r.setTaskInstanceNote = &setTaskInstanceNote
	return r
}

func (r ApiSetMappedTaskInstanceNoteRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.SetMappedTaskInstanceNoteExecute(r)
}

/*
SetMappedTaskInstanceNote Update the TaskInstance note.

Update the manual user note of a mapped Task Instance.

*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @param mapIndex The map index.
 @return ApiSetMappedTaskInstanceNoteRequest
*/
func (a *TaskInstanceAPIService) SetMappedTaskInstanceNote(ctx context.Context, dagId string, dagRunId string, taskId string, mapIndex int32) ApiSetMappedTaskInstanceNoteRequest {
	return ApiSetMappedTaskInstanceNoteRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
		mapIndex: mapIndex,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceAPIService) SetMappedTaskInstanceNoteExecute(r ApiSetMappedTaskInstanceNoteRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.SetMappedTaskInstanceNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/setNote"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"map_index"+"}", url.PathEscape(parameterValueToString(r.mapIndex, "mapIndex")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setTaskInstanceNote == nil {
		return localVarReturnValue, nil, reportError("setTaskInstanceNote is required and must be specified")
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
	// body params
	localVarPostBody = r.setTaskInstanceNote
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiSetTaskInstanceNoteRequest struct {
	ctx context.Context
	ApiService *TaskInstanceAPIService
	dagId string
	dagRunId string
	taskId string
	setTaskInstanceNote *SetTaskInstanceNote
}

// Parameters of set Task Instance note.
func (r ApiSetTaskInstanceNoteRequest) SetTaskInstanceNote(setTaskInstanceNote SetTaskInstanceNote) ApiSetTaskInstanceNoteRequest {
	r.setTaskInstanceNote = &setTaskInstanceNote
	return r
}

func (r ApiSetTaskInstanceNoteRequest) Execute() (*TaskInstance, *http.Response, error) {
	return r.ApiService.SetTaskInstanceNoteExecute(r)
}

/*
SetTaskInstanceNote Update the TaskInstance note.

Update the manual user note of a non-mapped Task Instance.

*New in version 2.5.0*


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dagId The DAG ID.
 @param dagRunId The DAG run ID.
 @param taskId The task ID.
 @return ApiSetTaskInstanceNoteRequest
*/
func (a *TaskInstanceAPIService) SetTaskInstanceNote(ctx context.Context, dagId string, dagRunId string, taskId string) ApiSetTaskInstanceNoteRequest {
	return ApiSetTaskInstanceNoteRequest{
		ApiService: a,
		ctx: ctx,
		dagId: dagId,
		dagRunId: dagRunId,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return TaskInstance
func (a *TaskInstanceAPIService) SetTaskInstanceNoteExecute(r ApiSetTaskInstanceNoteRequest) (*TaskInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TaskInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskInstanceAPIService.SetTaskInstanceNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/setNote"
	localVarPath = strings.Replace(localVarPath, "{"+"dag_id"+"}", url.PathEscape(parameterValueToString(r.dagId, "dagId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dag_run_id"+"}", url.PathEscape(parameterValueToString(r.dagRunId, "dagRunId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setTaskInstanceNote == nil {
		return localVarReturnValue, nil, reportError("setTaskInstanceNote is required and must be specified")
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
	// body params
	localVarPostBody = r.setTaskInstanceNote
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
