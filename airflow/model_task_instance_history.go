/*
Airflow API (Stable)

# Overview  To facilitate management, Apache Airflow supports a range of REST API endpoints across its objects. This section provides an overview of the API design, methods, and supported use cases.  Most of the endpoints accept `JSON` as input and return `JSON` responses. This means that you must usually add the following headers to your request: ``` Content-type: application/json Accept: application/json ```  ## Resources  The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its endpoint's corresponding resource. The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.  Resource names are used as part of endpoint URLs, as well as in API parameters and responses.  ## CRUD Operations  The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources. You can review the standards for these operations and their standard parameters below.  Some endpoints have special behavior as exceptions.  ### Create  To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata in the request body. The response returns a `201 Created` response code upon success with the resource's metadata, including its internal `id`, in the response body.  ### Read  The HTTP `GET` request can be used to read a resource or to list a number of resources.  A resource's `id` can be submitted in the request parameters to read a specific resource. The response usually returns a `200 OK` response code upon success, with the resource's metadata in the response body.  If a `GET` request does not include a specific resource `id`, it is treated as a list request. The response usually returns a `200 OK` response code upon success, with an object containing a list of resources' metadata in the response body.  When reading resources, some common query parameters are usually available. e.g.: ``` v1/connections?limit=25&offset=25 ```  |Query Parameter|Type|Description| |---------------|----|-----------| |limit|integer|Maximum number of objects to fetch. Usually 25 by default| |offset|integer|Offset after which to start returning objects. For use with limit query parameter.|  ### Update  Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request, with the fields to modify in the request body. The response usually returns a `200 OK` response code upon success, with information about the modified resource in the response body.  ### Delete  Deleting a resource requires the resource `id` and is typically executed via an HTTP `DELETE` request. The response usually returns a `204 No Content` response code upon success.  ## Conventions  - Resource names are plural and expressed in camelCase. - Names are consistent between URL parameter name and field name.  - Field names are in snake_case. ```json {     \"description\": \"string\",     \"name\": \"string\",     \"occupied_slots\": 0,     \"open_slots\": 0     \"queued_slots\": 0,     \"running_slots\": 0,     \"scheduled_slots\": 0,     \"slots\": 0, } ```  ### Update Mask  Update mask is available as a query parameter in patch endpoints. It is used to notify the API which fields you want to update. Using `update_mask` makes it easier to update objects by helping the server know which fields to update in an object instead of updating all fields. The update request ignores any fields that aren't specified in the field mask, leaving them with their current values.  Example: ```   resource = request.get('/resource/my-id').json()   resource['my_field'] = 'new-value'   request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource)) ```  ## Versioning and Endpoint Lifecycle  - API versioning is not synchronized to specific releases of the Apache Airflow. - APIs are designed to be backward compatible. - Any changes to the API will first go through a deprecation phase.  # Trying the API  You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/), [Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test the Apache Airflow API.  Note that you will need to pass credentials data.  For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used: ```bash curl -X PATCH 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\ -H 'Content-Type: application/json' \\ --user \"username:password\" \\ -d '{     \"is_paused\": true }' ```  Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/), it is possible to import the API specifications directly:  1. Download the API specification by clicking the **Download** button at the top of this document 2. Import the JSON specification in the graphical tool of your choice.   - In *Postman*, you can click the **import** button at the top   - With *Insomnia*, you can just drag-and-drop the file on the UI  Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on the **Code** button.  ## Enabling CORS  [Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) is a browser security feature that restricts HTTP requests that are initiated from scripts running in the browser.  For details on enabling/configuring CORS, see [Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Authentication  To be able to meet the requirements of many organizations, Airflow supports many authentication methods, and it is even possible to add your own method.  If you want to check which auth backend is currently set, you can use `airflow config get-value api auth_backends` command as in the example below. ```bash $ airflow config get-value api auth_backends airflow.api.auth.backend.basic_auth ``` The default is to deny all requests.  For details on configuring the authentication, see [API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).  # Errors  We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807) also known as Problem Details for HTTP APIs. As with our normal API responses, your client must be prepared to gracefully handle additional members of the response.  ## Unauthenticated  This indicates that the request has not been applied because it lacks valid authentication credentials for the target resource. Please check that you have valid credentials.  ## PermissionDenied  This response means that the server understood the request but refuses to authorize it because it lacks sufficient rights to the resource. It happens when you do not have the necessary permission to execute the action you performed. You need to get the appropriate permissions in other to resolve this error.  ## BadRequest  This response means that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.  ## NotFound  This client error response indicates that the server cannot find the requested resource.  ## MethodNotAllowed  Indicates that the request method is known by the server but is not supported by the target resource.  ## NotAcceptable  The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request, and the server is unwilling to supply a default representation.  ## AlreadyExists  The request could not be completed due to a conflict with the current state of the target resource, e.g. the resource it tries to create already exists.  ## Unknown  This means that the server encountered an unexpected condition that prevented it from fulfilling the request. 

API version: 2.10.5
Contact: dev@airflow.apache.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airflow

import (
	"encoding/json"
)

// checks if the TaskInstanceHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskInstanceHistory{}

// TaskInstanceHistory struct for TaskInstanceHistory
type TaskInstanceHistory struct {
	DagId *string `json:"dag_id,omitempty"`
	// The DagRun ID for this task instance  *New in version 2.3.0* 
	DagRunId *string `json:"dag_run_id,omitempty"`
	Duration NullableFloat32 `json:"duration,omitempty"`
	EndDate NullableString `json:"end_date,omitempty"`
	// Executor the task is configured to run on or None (which indicates the default executor)  *New in version 2.10.0* 
	Executor NullableString `json:"executor,omitempty"`
	ExecutorConfig *string `json:"executor_config,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	MapIndex *int32 `json:"map_index,omitempty"`
	MaxTries *int32 `json:"max_tries,omitempty"`
	// *Changed in version 2.1.1*&#58; Field becomes nullable. 
	Operator NullableString `json:"operator,omitempty"`
	Pid NullableInt32 `json:"pid,omitempty"`
	Pool *string `json:"pool,omitempty"`
	PoolSlots *int32 `json:"pool_slots,omitempty"`
	PriorityWeight NullableInt32 `json:"priority_weight,omitempty"`
	Queue NullableString `json:"queue,omitempty"`
	// The datetime that the task enter the state QUEUE, also known as queue_at 
	QueuedWhen NullableString `json:"queued_when,omitempty"`
	StartDate NullableString `json:"start_date,omitempty"`
	State NullableTaskState `json:"state,omitempty"`
	// Human centric display text for the task.  *New in version 2.9.0* 
	TaskDisplayName *string `json:"task_display_name,omitempty"`
	TaskId *string `json:"task_id,omitempty"`
	TryNumber *int32 `json:"try_number,omitempty"`
	Unixname *string `json:"unixname,omitempty"`
}

// NewTaskInstanceHistory instantiates a new TaskInstanceHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskInstanceHistory() *TaskInstanceHistory {
	this := TaskInstanceHistory{}
	return &this
}

// NewTaskInstanceHistoryWithDefaults instantiates a new TaskInstanceHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskInstanceHistoryWithDefaults() *TaskInstanceHistory {
	this := TaskInstanceHistory{}
	return &this
}

// GetDagId returns the DagId field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetDagId() string {
	if o == nil || IsNil(o.DagId) {
		var ret string
		return ret
	}
	return *o.DagId
}

// GetDagIdOk returns a tuple with the DagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetDagIdOk() (*string, bool) {
	if o == nil || IsNil(o.DagId) {
		return nil, false
	}
	return o.DagId, true
}

// HasDagId returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasDagId() bool {
	if o != nil && !IsNil(o.DagId) {
		return true
	}

	return false
}

// SetDagId gets a reference to the given string and assigns it to the DagId field.
func (o *TaskInstanceHistory) SetDagId(v string) {
	o.DagId = &v
}

// GetDagRunId returns the DagRunId field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetDagRunId() string {
	if o == nil || IsNil(o.DagRunId) {
		var ret string
		return ret
	}
	return *o.DagRunId
}

// GetDagRunIdOk returns a tuple with the DagRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetDagRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.DagRunId) {
		return nil, false
	}
	return o.DagRunId, true
}

// HasDagRunId returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasDagRunId() bool {
	if o != nil && !IsNil(o.DagRunId) {
		return true
	}

	return false
}

// SetDagRunId gets a reference to the given string and assigns it to the DagRunId field.
func (o *TaskInstanceHistory) SetDagRunId(v string) {
	o.DagRunId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetDuration() float32 {
	if o == nil || IsNil(o.Duration.Get()) {
		var ret float32
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableFloat32 and assigns it to the Duration field.
func (o *TaskInstanceHistory) SetDuration(v float32) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *TaskInstanceHistory) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *TaskInstanceHistory) UnsetDuration() {
	o.Duration.Unset()
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetEndDate() string {
	if o == nil || IsNil(o.EndDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *TaskInstanceHistory) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *TaskInstanceHistory) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *TaskInstanceHistory) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetExecutor returns the Executor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetExecutor() string {
	if o == nil || IsNil(o.Executor.Get()) {
		var ret string
		return ret
	}
	return *o.Executor.Get()
}

// GetExecutorOk returns a tuple with the Executor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetExecutorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Executor.Get(), o.Executor.IsSet()
}

// HasExecutor returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasExecutor() bool {
	if o != nil && o.Executor.IsSet() {
		return true
	}

	return false
}

// SetExecutor gets a reference to the given NullableString and assigns it to the Executor field.
func (o *TaskInstanceHistory) SetExecutor(v string) {
	o.Executor.Set(&v)
}
// SetExecutorNil sets the value for Executor to be an explicit nil
func (o *TaskInstanceHistory) SetExecutorNil() {
	o.Executor.Set(nil)
}

// UnsetExecutor ensures that no value is present for Executor, not even an explicit nil
func (o *TaskInstanceHistory) UnsetExecutor() {
	o.Executor.Unset()
}

// GetExecutorConfig returns the ExecutorConfig field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetExecutorConfig() string {
	if o == nil || IsNil(o.ExecutorConfig) {
		var ret string
		return ret
	}
	return *o.ExecutorConfig
}

// GetExecutorConfigOk returns a tuple with the ExecutorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetExecutorConfigOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutorConfig) {
		return nil, false
	}
	return o.ExecutorConfig, true
}

// HasExecutorConfig returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasExecutorConfig() bool {
	if o != nil && !IsNil(o.ExecutorConfig) {
		return true
	}

	return false
}

// SetExecutorConfig gets a reference to the given string and assigns it to the ExecutorConfig field.
func (o *TaskInstanceHistory) SetExecutorConfig(v string) {
	o.ExecutorConfig = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *TaskInstanceHistory) SetHostname(v string) {
	o.Hostname = &v
}

// GetMapIndex returns the MapIndex field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetMapIndex() int32 {
	if o == nil || IsNil(o.MapIndex) {
		var ret int32
		return ret
	}
	return *o.MapIndex
}

// GetMapIndexOk returns a tuple with the MapIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetMapIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.MapIndex) {
		return nil, false
	}
	return o.MapIndex, true
}

// HasMapIndex returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasMapIndex() bool {
	if o != nil && !IsNil(o.MapIndex) {
		return true
	}

	return false
}

// SetMapIndex gets a reference to the given int32 and assigns it to the MapIndex field.
func (o *TaskInstanceHistory) SetMapIndex(v int32) {
	o.MapIndex = &v
}

// GetMaxTries returns the MaxTries field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetMaxTries() int32 {
	if o == nil || IsNil(o.MaxTries) {
		var ret int32
		return ret
	}
	return *o.MaxTries
}

// GetMaxTriesOk returns a tuple with the MaxTries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetMaxTriesOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTries) {
		return nil, false
	}
	return o.MaxTries, true
}

// HasMaxTries returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasMaxTries() bool {
	if o != nil && !IsNil(o.MaxTries) {
		return true
	}

	return false
}

// SetMaxTries gets a reference to the given int32 and assigns it to the MaxTries field.
func (o *TaskInstanceHistory) SetMaxTries(v int32) {
	o.MaxTries = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetOperator() string {
	if o == nil || IsNil(o.Operator.Get()) {
		var ret string
		return ret
	}
	return *o.Operator.Get()
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator.Get(), o.Operator.IsSet()
}

// HasOperator returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasOperator() bool {
	if o != nil && o.Operator.IsSet() {
		return true
	}

	return false
}

// SetOperator gets a reference to the given NullableString and assigns it to the Operator field.
func (o *TaskInstanceHistory) SetOperator(v string) {
	o.Operator.Set(&v)
}
// SetOperatorNil sets the value for Operator to be an explicit nil
func (o *TaskInstanceHistory) SetOperatorNil() {
	o.Operator.Set(nil)
}

// UnsetOperator ensures that no value is present for Operator, not even an explicit nil
func (o *TaskInstanceHistory) UnsetOperator() {
	o.Operator.Unset()
}

// GetPid returns the Pid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetPid() int32 {
	if o == nil || IsNil(o.Pid.Get()) {
		var ret int32
		return ret
	}
	return *o.Pid.Get()
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetPidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pid.Get(), o.Pid.IsSet()
}

// HasPid returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasPid() bool {
	if o != nil && o.Pid.IsSet() {
		return true
	}

	return false
}

// SetPid gets a reference to the given NullableInt32 and assigns it to the Pid field.
func (o *TaskInstanceHistory) SetPid(v int32) {
	o.Pid.Set(&v)
}
// SetPidNil sets the value for Pid to be an explicit nil
func (o *TaskInstanceHistory) SetPidNil() {
	o.Pid.Set(nil)
}

// UnsetPid ensures that no value is present for Pid, not even an explicit nil
func (o *TaskInstanceHistory) UnsetPid() {
	o.Pid.Unset()
}

// GetPool returns the Pool field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetPool() string {
	if o == nil || IsNil(o.Pool) {
		var ret string
		return ret
	}
	return *o.Pool
}

// GetPoolOk returns a tuple with the Pool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetPoolOk() (*string, bool) {
	if o == nil || IsNil(o.Pool) {
		return nil, false
	}
	return o.Pool, true
}

// HasPool returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasPool() bool {
	if o != nil && !IsNil(o.Pool) {
		return true
	}

	return false
}

// SetPool gets a reference to the given string and assigns it to the Pool field.
func (o *TaskInstanceHistory) SetPool(v string) {
	o.Pool = &v
}

// GetPoolSlots returns the PoolSlots field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetPoolSlots() int32 {
	if o == nil || IsNil(o.PoolSlots) {
		var ret int32
		return ret
	}
	return *o.PoolSlots
}

// GetPoolSlotsOk returns a tuple with the PoolSlots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetPoolSlotsOk() (*int32, bool) {
	if o == nil || IsNil(o.PoolSlots) {
		return nil, false
	}
	return o.PoolSlots, true
}

// HasPoolSlots returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasPoolSlots() bool {
	if o != nil && !IsNil(o.PoolSlots) {
		return true
	}

	return false
}

// SetPoolSlots gets a reference to the given int32 and assigns it to the PoolSlots field.
func (o *TaskInstanceHistory) SetPoolSlots(v int32) {
	o.PoolSlots = &v
}

// GetPriorityWeight returns the PriorityWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetPriorityWeight() int32 {
	if o == nil || IsNil(o.PriorityWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.PriorityWeight.Get()
}

// GetPriorityWeightOk returns a tuple with the PriorityWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetPriorityWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PriorityWeight.Get(), o.PriorityWeight.IsSet()
}

// HasPriorityWeight returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasPriorityWeight() bool {
	if o != nil && o.PriorityWeight.IsSet() {
		return true
	}

	return false
}

// SetPriorityWeight gets a reference to the given NullableInt32 and assigns it to the PriorityWeight field.
func (o *TaskInstanceHistory) SetPriorityWeight(v int32) {
	o.PriorityWeight.Set(&v)
}
// SetPriorityWeightNil sets the value for PriorityWeight to be an explicit nil
func (o *TaskInstanceHistory) SetPriorityWeightNil() {
	o.PriorityWeight.Set(nil)
}

// UnsetPriorityWeight ensures that no value is present for PriorityWeight, not even an explicit nil
func (o *TaskInstanceHistory) UnsetPriorityWeight() {
	o.PriorityWeight.Unset()
}

// GetQueue returns the Queue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetQueue() string {
	if o == nil || IsNil(o.Queue.Get()) {
		var ret string
		return ret
	}
	return *o.Queue.Get()
}

// GetQueueOk returns a tuple with the Queue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetQueueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Queue.Get(), o.Queue.IsSet()
}

// HasQueue returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasQueue() bool {
	if o != nil && o.Queue.IsSet() {
		return true
	}

	return false
}

// SetQueue gets a reference to the given NullableString and assigns it to the Queue field.
func (o *TaskInstanceHistory) SetQueue(v string) {
	o.Queue.Set(&v)
}
// SetQueueNil sets the value for Queue to be an explicit nil
func (o *TaskInstanceHistory) SetQueueNil() {
	o.Queue.Set(nil)
}

// UnsetQueue ensures that no value is present for Queue, not even an explicit nil
func (o *TaskInstanceHistory) UnsetQueue() {
	o.Queue.Unset()
}

// GetQueuedWhen returns the QueuedWhen field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetQueuedWhen() string {
	if o == nil || IsNil(o.QueuedWhen.Get()) {
		var ret string
		return ret
	}
	return *o.QueuedWhen.Get()
}

// GetQueuedWhenOk returns a tuple with the QueuedWhen field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetQueuedWhenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.QueuedWhen.Get(), o.QueuedWhen.IsSet()
}

// HasQueuedWhen returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasQueuedWhen() bool {
	if o != nil && o.QueuedWhen.IsSet() {
		return true
	}

	return false
}

// SetQueuedWhen gets a reference to the given NullableString and assigns it to the QueuedWhen field.
func (o *TaskInstanceHistory) SetQueuedWhen(v string) {
	o.QueuedWhen.Set(&v)
}
// SetQueuedWhenNil sets the value for QueuedWhen to be an explicit nil
func (o *TaskInstanceHistory) SetQueuedWhenNil() {
	o.QueuedWhen.Set(nil)
}

// UnsetQueuedWhen ensures that no value is present for QueuedWhen, not even an explicit nil
func (o *TaskInstanceHistory) UnsetQueuedWhen() {
	o.QueuedWhen.Unset()
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetStartDate() string {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret string
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableString and assigns it to the StartDate field.
func (o *TaskInstanceHistory) SetStartDate(v string) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *TaskInstanceHistory) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *TaskInstanceHistory) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskInstanceHistory) GetState() TaskState {
	if o == nil || IsNil(o.State.Get()) {
		var ret TaskState
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskInstanceHistory) GetStateOk() (*TaskState, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableTaskState and assigns it to the State field.
func (o *TaskInstanceHistory) SetState(v TaskState) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *TaskInstanceHistory) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *TaskInstanceHistory) UnsetState() {
	o.State.Unset()
}

// GetTaskDisplayName returns the TaskDisplayName field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetTaskDisplayName() string {
	if o == nil || IsNil(o.TaskDisplayName) {
		var ret string
		return ret
	}
	return *o.TaskDisplayName
}

// GetTaskDisplayNameOk returns a tuple with the TaskDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetTaskDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.TaskDisplayName) {
		return nil, false
	}
	return o.TaskDisplayName, true
}

// HasTaskDisplayName returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasTaskDisplayName() bool {
	if o != nil && !IsNil(o.TaskDisplayName) {
		return true
	}

	return false
}

// SetTaskDisplayName gets a reference to the given string and assigns it to the TaskDisplayName field.
func (o *TaskInstanceHistory) SetTaskDisplayName(v string) {
	o.TaskDisplayName = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *TaskInstanceHistory) SetTaskId(v string) {
	o.TaskId = &v
}

// GetTryNumber returns the TryNumber field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetTryNumber() int32 {
	if o == nil || IsNil(o.TryNumber) {
		var ret int32
		return ret
	}
	return *o.TryNumber
}

// GetTryNumberOk returns a tuple with the TryNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetTryNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.TryNumber) {
		return nil, false
	}
	return o.TryNumber, true
}

// HasTryNumber returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasTryNumber() bool {
	if o != nil && !IsNil(o.TryNumber) {
		return true
	}

	return false
}

// SetTryNumber gets a reference to the given int32 and assigns it to the TryNumber field.
func (o *TaskInstanceHistory) SetTryNumber(v int32) {
	o.TryNumber = &v
}

// GetUnixname returns the Unixname field value if set, zero value otherwise.
func (o *TaskInstanceHistory) GetUnixname() string {
	if o == nil || IsNil(o.Unixname) {
		var ret string
		return ret
	}
	return *o.Unixname
}

// GetUnixnameOk returns a tuple with the Unixname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskInstanceHistory) GetUnixnameOk() (*string, bool) {
	if o == nil || IsNil(o.Unixname) {
		return nil, false
	}
	return o.Unixname, true
}

// HasUnixname returns a boolean if a field has been set.
func (o *TaskInstanceHistory) HasUnixname() bool {
	if o != nil && !IsNil(o.Unixname) {
		return true
	}

	return false
}

// SetUnixname gets a reference to the given string and assigns it to the Unixname field.
func (o *TaskInstanceHistory) SetUnixname(v string) {
	o.Unixname = &v
}

func (o TaskInstanceHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskInstanceHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DagId) {
		toSerialize["dag_id"] = o.DagId
	}
	if !IsNil(o.DagRunId) {
		toSerialize["dag_run_id"] = o.DagRunId
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.Executor.IsSet() {
		toSerialize["executor"] = o.Executor.Get()
	}
	if !IsNil(o.ExecutorConfig) {
		toSerialize["executor_config"] = o.ExecutorConfig
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.MapIndex) {
		toSerialize["map_index"] = o.MapIndex
	}
	if !IsNil(o.MaxTries) {
		toSerialize["max_tries"] = o.MaxTries
	}
	if o.Operator.IsSet() {
		toSerialize["operator"] = o.Operator.Get()
	}
	if o.Pid.IsSet() {
		toSerialize["pid"] = o.Pid.Get()
	}
	if !IsNil(o.Pool) {
		toSerialize["pool"] = o.Pool
	}
	if !IsNil(o.PoolSlots) {
		toSerialize["pool_slots"] = o.PoolSlots
	}
	if o.PriorityWeight.IsSet() {
		toSerialize["priority_weight"] = o.PriorityWeight.Get()
	}
	if o.Queue.IsSet() {
		toSerialize["queue"] = o.Queue.Get()
	}
	if o.QueuedWhen.IsSet() {
		toSerialize["queued_when"] = o.QueuedWhen.Get()
	}
	if o.StartDate.IsSet() {
		toSerialize["start_date"] = o.StartDate.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if !IsNil(o.TaskDisplayName) {
		toSerialize["task_display_name"] = o.TaskDisplayName
	}
	if !IsNil(o.TaskId) {
		toSerialize["task_id"] = o.TaskId
	}
	if !IsNil(o.TryNumber) {
		toSerialize["try_number"] = o.TryNumber
	}
	if !IsNil(o.Unixname) {
		toSerialize["unixname"] = o.Unixname
	}
	return toSerialize, nil
}

type NullableTaskInstanceHistory struct {
	value *TaskInstanceHistory
	isSet bool
}

func (v NullableTaskInstanceHistory) Get() *TaskInstanceHistory {
	return v.value
}

func (v *NullableTaskInstanceHistory) Set(val *TaskInstanceHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskInstanceHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskInstanceHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskInstanceHistory(val *TaskInstanceHistory) *NullableTaskInstanceHistory {
	return &NullableTaskInstanceHistory{value: val, isSet: true}
}

func (v NullableTaskInstanceHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskInstanceHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


