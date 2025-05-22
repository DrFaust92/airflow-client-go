<!--
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
 -->
# Go API client for airflow

# Overview

To facilitate management, Apache Airflow supports a range of REST API endpoints across its
objects.
This section provides an overview of the API design, methods, and supported use cases.

Most of the endpoints accept `JSON` as input and return `JSON` responses.
This means that you must usually add the following headers to your request:
```
Content-type: application/json
Accept: application/json
```

## Resources

The term `resource` refers to a single type of object in the Airflow metadata. An API is broken up by its
endpoint's corresponding resource.
The name of a resource is typically plural and expressed in camelCase. Example: `dagRuns`.

Resource names are used as part of endpoint URLs, as well as in API parameters and responses.

## CRUD Operations

The platform supports **C**reate, **R**ead, **U**pdate, and **D**elete operations on most resources.
You can review the standards for these operations and their standard parameters below.

Some endpoints have special behavior as exceptions.

### Create

To create a resource, you typically submit an HTTP `POST` request with the resource's required metadata
in the request body.
The response returns a `201 Created` response code upon success with the resource's metadata, including
its internal `id`, in the response body.

### Read

The HTTP `GET` request can be used to read a resource or to list a number of resources.

A resource's `id` can be submitted in the request parameters to read a specific resource.
The response usually returns a `200 OK` response code upon success, with the resource's metadata in
the response body.

If a `GET` request does not include a specific resource `id`, it is treated as a list request.
The response usually returns a `200 OK` response code upon success, with an object containing a list
of resources' metadata in the response body.

When reading resources, some common query parameters are usually available. e.g.:
```
v1/connections?limit=25&offset=25
```

|Query Parameter|Type|Description|
|---------------|----|-----------|
|limit|integer|Maximum number of objects to fetch. Usually 25 by default|
|offset|integer|Offset after which to start returning objects. For use with limit query parameter.|

### Update

Updating a resource requires the resource `id`, and is typically done using an HTTP `PATCH` request,
with the fields to modify in the request body.
The response usually returns a `200 OK` response code upon success, with information about the modified
resource in the response body.

### Delete

Deleting a resource requires the resource `id` and is typically executed via an HTTP `DELETE` request.
The response usually returns a `204 No Content` response code upon success.

## Conventions

- Resource names are plural and expressed in camelCase.
- Names are consistent between URL parameter name and field name.

- Field names are in snake_case.
```json
{
    \"description\": \"string\",
    \"name\": \"string\",
    \"occupied_slots\": 0,
    \"open_slots\": 0
    \"queued_slots\": 0,
    \"running_slots\": 0,
    \"scheduled_slots\": 0,
    \"slots\": 0,
}
```

### Update Mask

Update mask is available as a query parameter in patch endpoints. It is used to notify the
API which fields you want to update. Using `update_mask` makes it easier to update objects
by helping the server know which fields to update in an object instead of updating all fields.
The update request ignores any fields that aren't specified in the field mask, leaving them with
their current values.

Example:
```
  resource = request.get('/resource/my-id').json()
  resource['my_field'] = 'new-value'
  request.patch('/resource/my-id?update_mask=my_field', data=json.dumps(resource))
```

## Versioning and Endpoint Lifecycle

- API versioning is not synchronized to specific releases of the Apache Airflow.
- APIs are designed to be backward compatible.
- Any changes to the API will first go through a deprecation phase.

# Trying the API

You can use a third party client, such as [curl](https://curl.haxx.se/), [HTTPie](https://httpie.org/),
[Postman](https://www.postman.com/) or [the Insomnia rest client](https://insomnia.rest/) to test
the Apache Airflow API.

Note that you will need to pass credentials data.

For e.g., here is how to pause a DAG with [curl](https://curl.haxx.se/), when basic authorization is used:
```bash
curl -X PATCH 'https://example.com/api/v1/dags/{dag_id}?update_mask=is_paused' \\
-H 'Content-Type: application/json' \\
--user \"username:password\" \\
-d '{
    \"is_paused\": true
}'
```

Using a graphical tool such as [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/),
it is possible to import the API specifications directly:

1. Download the API specification by clicking the **Download** button at the top of this document
2. Import the JSON specification in the graphical tool of your choice.
  - In *Postman*, you can click the **import** button at the top
  - With *Insomnia*, you can just drag-and-drop the file on the UI

Note that with *Postman*, you can also generate code snippets by selecting a request and clicking on
the **Code** button.

## Enabling CORS

[Cross-origin resource sharing (CORS)](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
is a browser security feature that restricts HTTP requests that are
initiated from scripts running in the browser.

For details on enabling/configuring CORS, see
[Enabling CORS](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).

# Authentication

To be able to meet the requirements of many organizations, Airflow supports many authentication methods,
and it is even possible to add your own method.

If you want to check which auth backend is currently set, you can use
`airflow config get-value api auth_backends` command as in the example below.
```bash
$ airflow config get-value api auth_backends
airflow.api.auth.backend.basic_auth
```
The default is to deny all requests.

For details on configuring the authentication, see
[API Authorization](https://airflow.apache.org/docs/apache-airflow/stable/security/api.html).

# Errors

We follow the error response format proposed in [RFC 7807](https://tools.ietf.org/html/rfc7807)
also known as Problem Details for HTTP APIs. As with our normal API responses,
your client must be prepared to gracefully handle additional members of the response.

## Unauthenticated

This indicates that the request has not been applied because it lacks valid authentication
credentials for the target resource. Please check that you have valid credentials.

## PermissionDenied

This response means that the server understood the request but refuses to authorize
it because it lacks sufficient rights to the resource. It happens when you do not have the
necessary permission to execute the action you performed. You need to get the appropriate
permissions in other to resolve this error.

## BadRequest

This response means that the server cannot or will not process the request due to something
that is perceived to be a client error (e.g., malformed request syntax, invalid request message
framing, or deceptive request routing). To resolve this, please ensure that your syntax is correct.

## NotFound

This client error response indicates that the server cannot find the requested resource.

## MethodNotAllowed

Indicates that the request method is known by the server but is not supported by the target resource.

## NotAcceptable

The target resource does not have a current representation that would be acceptable to the user
agent, according to the proactive negotiation header fields received in the request, and the
server is unwilling to supply a default representation.

## AlreadyExists

The request could not be completed due to a conflict with the current state of the target
resource, e.g. the resource it tries to create already exists.

## Unknown

This means that the server encountered an unexpected condition that prevented it from
fulfilling the request.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.10.5
- Package version: 2.10.5
- Generator version: 7.13.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://airflow.apache.org](https://airflow.apache.org)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import fab "github.com/airflow/airflow-client-go/airflow/"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `airflow.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), airflow.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `airflow.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), airflow.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `airflow.ContextOperationServerIndices` and `airflow.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), airflow.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), airflow.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigAPI* | [**GetConfig**](docs/ConfigAPI.md#getconfig) | **Get** /config | Get current configuration
*ConfigAPI* | [**GetValue**](docs/ConfigAPI.md#getvalue) | **Get** /config/section/{section}/option/{option} | Get a option from configuration
*ConnectionAPI* | [**DeleteConnection**](docs/ConnectionAPI.md#deleteconnection) | **Delete** /connections/{connection_id} | Delete a connection
*ConnectionAPI* | [**GetConnection**](docs/ConnectionAPI.md#getconnection) | **Get** /connections/{connection_id} | Get a connection
*ConnectionAPI* | [**GetConnections**](docs/ConnectionAPI.md#getconnections) | **Get** /connections | List connections
*ConnectionAPI* | [**PatchConnection**](docs/ConnectionAPI.md#patchconnection) | **Patch** /connections/{connection_id} | Update a connection
*ConnectionAPI* | [**PostConnection**](docs/ConnectionAPI.md#postconnection) | **Post** /connections | Create a connection
*ConnectionAPI* | [**TestConnection**](docs/ConnectionAPI.md#testconnection) | **Post** /connections/test | Test a connection
*DAGAPI* | [**DeleteDag**](docs/DAGAPI.md#deletedag) | **Delete** /dags/{dag_id} | Delete a DAG
*DAGAPI* | [**GetDag**](docs/DAGAPI.md#getdag) | **Get** /dags/{dag_id} | Get basic information about a DAG
*DAGAPI* | [**GetDagDetails**](docs/DAGAPI.md#getdagdetails) | **Get** /dags/{dag_id}/details | Get a simplified representation of DAG
*DAGAPI* | [**GetDagSource**](docs/DAGAPI.md#getdagsource) | **Get** /dagSources/{file_token} | Get a source code
*DAGAPI* | [**GetDags**](docs/DAGAPI.md#getdags) | **Get** /dags | List DAGs
*DAGAPI* | [**GetTask**](docs/DAGAPI.md#gettask) | **Get** /dags/{dag_id}/tasks/{task_id} | Get simplified representation of a task
*DAGAPI* | [**GetTasks**](docs/DAGAPI.md#gettasks) | **Get** /dags/{dag_id}/tasks | Get tasks for DAG
*DAGAPI* | [**PatchDag**](docs/DAGAPI.md#patchdag) | **Patch** /dags/{dag_id} | Update a DAG
*DAGAPI* | [**PatchDags**](docs/DAGAPI.md#patchdags) | **Patch** /dags | Update DAGs
*DAGAPI* | [**PostClearTaskInstances**](docs/DAGAPI.md#postcleartaskinstances) | **Post** /dags/{dag_id}/clearTaskInstances | Clear a set of task instances
*DAGAPI* | [**PostSetTaskInstancesState**](docs/DAGAPI.md#postsettaskinstancesstate) | **Post** /dags/{dag_id}/updateTaskInstancesState | Set a state of task instances
*DAGAPI* | [**ReparseDagFile**](docs/DAGAPI.md#reparsedagfile) | **Put** /parseDagFile/{file_token} | Request re-parsing of a DAG file
*DAGRunAPI* | [**ClearDagRun**](docs/DAGRunAPI.md#cleardagrun) | **Post** /dags/{dag_id}/dagRuns/{dag_run_id}/clear | Clear a DAG run
*DAGRunAPI* | [**DeleteDagRun**](docs/DAGRunAPI.md#deletedagrun) | **Delete** /dags/{dag_id}/dagRuns/{dag_run_id} | Delete a DAG run
*DAGRunAPI* | [**GetDagRun**](docs/DAGRunAPI.md#getdagrun) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id} | Get a DAG run
*DAGRunAPI* | [**GetDagRuns**](docs/DAGRunAPI.md#getdagruns) | **Get** /dags/{dag_id}/dagRuns | List DAG runs
*DAGRunAPI* | [**GetDagRunsBatch**](docs/DAGRunAPI.md#getdagrunsbatch) | **Post** /dags/~/dagRuns/list | List DAG runs (batch)
*DAGRunAPI* | [**GetUpstreamDatasetEvents**](docs/DAGRunAPI.md#getupstreamdatasetevents) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents | Get dataset events for a DAG run
*DAGRunAPI* | [**PostDagRun**](docs/DAGRunAPI.md#postdagrun) | **Post** /dags/{dag_id}/dagRuns | Trigger a new DAG run.
*DAGRunAPI* | [**SetDagRunNote**](docs/DAGRunAPI.md#setdagrunnote) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/setNote | Update the DagRun note.
*DAGRunAPI* | [**UpdateDagRunState**](docs/DAGRunAPI.md#updatedagrunstate) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id} | Modify a DAG run
*DagStatsAPI* | [**GetDagStats**](docs/DagStatsAPI.md#getdagstats) | **Get** /dagStats | List Dag statistics
*DagWarningAPI* | [**GetDagWarnings**](docs/DagWarningAPI.md#getdagwarnings) | **Get** /dagWarnings | List dag warnings
*DatasetAPI* | [**CreateDatasetEvent**](docs/DatasetAPI.md#createdatasetevent) | **Post** /datasets/events | Create dataset event
*DatasetAPI* | [**DeleteDagDatasetQueuedEvent**](docs/DatasetAPI.md#deletedagdatasetqueuedevent) | **Delete** /dags/{dag_id}/datasets/queuedEvent/{uri} | Delete a queued Dataset event for a DAG.
*DatasetAPI* | [**DeleteDagDatasetQueuedEvents**](docs/DatasetAPI.md#deletedagdatasetqueuedevents) | **Delete** /dags/{dag_id}/datasets/queuedEvent | Delete queued Dataset events for a DAG.
*DatasetAPI* | [**DeleteDatasetQueuedEvents**](docs/DatasetAPI.md#deletedatasetqueuedevents) | **Delete** /datasets/queuedEvent/{uri} | Delete queued Dataset events for a Dataset.
*DatasetAPI* | [**GetDagDatasetQueuedEvent**](docs/DatasetAPI.md#getdagdatasetqueuedevent) | **Get** /dags/{dag_id}/datasets/queuedEvent/{uri} | Get a queued Dataset event for a DAG
*DatasetAPI* | [**GetDagDatasetQueuedEvents**](docs/DatasetAPI.md#getdagdatasetqueuedevents) | **Get** /dags/{dag_id}/datasets/queuedEvent | Get queued Dataset events for a DAG.
*DatasetAPI* | [**GetDataset**](docs/DatasetAPI.md#getdataset) | **Get** /datasets/{uri} | Get a dataset
*DatasetAPI* | [**GetDatasetEvents**](docs/DatasetAPI.md#getdatasetevents) | **Get** /datasets/events | Get dataset events
*DatasetAPI* | [**GetDatasetQueuedEvents**](docs/DatasetAPI.md#getdatasetqueuedevents) | **Get** /datasets/queuedEvent/{uri} | Get queued Dataset events for a Dataset.
*DatasetAPI* | [**GetDatasets**](docs/DatasetAPI.md#getdatasets) | **Get** /datasets | List datasets
*DatasetAPI* | [**GetUpstreamDatasetEvents**](docs/DatasetAPI.md#getupstreamdatasetevents) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/upstreamDatasetEvents | Get dataset events for a DAG run
*EventLogAPI* | [**GetEventLog**](docs/EventLogAPI.md#geteventlog) | **Get** /eventLogs/{event_log_id} | Get a log entry
*EventLogAPI* | [**GetEventLogs**](docs/EventLogAPI.md#geteventlogs) | **Get** /eventLogs | List log entries
*ImportErrorAPI* | [**GetImportError**](docs/ImportErrorAPI.md#getimporterror) | **Get** /importErrors/{import_error_id} | Get an import error
*ImportErrorAPI* | [**GetImportErrors**](docs/ImportErrorAPI.md#getimporterrors) | **Get** /importErrors | List import errors
*MonitoringAPI* | [**GetHealth**](docs/MonitoringAPI.md#gethealth) | **Get** /health | Get instance status
*MonitoringAPI* | [**GetVersion**](docs/MonitoringAPI.md#getversion) | **Get** /version | Get version information
*PermissionAPI* | [**GetPermissions**](docs/PermissionAPI.md#getpermissions) | **Get** /permissions | List permissions
*PluginAPI* | [**GetPlugins**](docs/PluginAPI.md#getplugins) | **Get** /plugins | Get a list of loaded plugins
*PoolAPI* | [**DeletePool**](docs/PoolAPI.md#deletepool) | **Delete** /pools/{pool_name} | Delete a pool
*PoolAPI* | [**GetPool**](docs/PoolAPI.md#getpool) | **Get** /pools/{pool_name} | Get a pool
*PoolAPI* | [**GetPools**](docs/PoolAPI.md#getpools) | **Get** /pools | List pools
*PoolAPI* | [**PatchPool**](docs/PoolAPI.md#patchpool) | **Patch** /pools/{pool_name} | Update a pool
*PoolAPI* | [**PostPool**](docs/PoolAPI.md#postpool) | **Post** /pools | Create a pool
*ProviderAPI* | [**GetProviders**](docs/ProviderAPI.md#getproviders) | **Get** /providers | List providers
*RoleAPI* | [**DeleteRole**](docs/RoleAPI.md#deleterole) | **Delete** /roles/{role_name} | Delete a role
*RoleAPI* | [**GetRole**](docs/RoleAPI.md#getrole) | **Get** /roles/{role_name} | Get a role
*RoleAPI* | [**GetRoles**](docs/RoleAPI.md#getroles) | **Get** /roles | List roles
*RoleAPI* | [**PatchRole**](docs/RoleAPI.md#patchrole) | **Patch** /roles/{role_name} | Update a role
*RoleAPI* | [**PostRole**](docs/RoleAPI.md#postrole) | **Post** /roles | Create a role
*TaskInstanceAPI* | [**GetExtraLinks**](docs/TaskInstanceAPI.md#getextralinks) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/links | List extra links
*TaskInstanceAPI* | [**GetLog**](docs/TaskInstanceAPI.md#getlog) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/logs/{task_try_number} | Get logs
*TaskInstanceAPI* | [**GetMappedTaskInstance**](docs/TaskInstanceAPI.md#getmappedtaskinstance) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index} | Get a mapped task instance
*TaskInstanceAPI* | [**GetMappedTaskInstanceDependencies**](docs/TaskInstanceAPI.md#getmappedtaskinstancedependencies) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/dependencies | Get task dependencies blocking task from getting scheduled.
*TaskInstanceAPI* | [**GetMappedTaskInstanceTries**](docs/TaskInstanceAPI.md#getmappedtaskinstancetries) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/tries | List mapped task instance tries
*TaskInstanceAPI* | [**GetMappedTaskInstanceTryDetails**](docs/TaskInstanceAPI.md#getmappedtaskinstancetrydetails) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/tries/{task_try_number} | get mapped taskinstance try
*TaskInstanceAPI* | [**GetMappedTaskInstances**](docs/TaskInstanceAPI.md#getmappedtaskinstances) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/listMapped | List mapped task instances
*TaskInstanceAPI* | [**GetTaskInstance**](docs/TaskInstanceAPI.md#gettaskinstance) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id} | Get a task instance
*TaskInstanceAPI* | [**GetTaskInstanceDependencies**](docs/TaskInstanceAPI.md#gettaskinstancedependencies) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/dependencies | Get task dependencies blocking task from getting scheduled.
*TaskInstanceAPI* | [**GetTaskInstanceTries**](docs/TaskInstanceAPI.md#gettaskinstancetries) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/tries | List task instance tries
*TaskInstanceAPI* | [**GetTaskInstanceTryDetails**](docs/TaskInstanceAPI.md#gettaskinstancetrydetails) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/tries/{task_try_number} | get taskinstance try
*TaskInstanceAPI* | [**GetTaskInstances**](docs/TaskInstanceAPI.md#gettaskinstances) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances | List task instances
*TaskInstanceAPI* | [**GetTaskInstancesBatch**](docs/TaskInstanceAPI.md#gettaskinstancesbatch) | **Post** /dags/~/dagRuns/~/taskInstances/list | List task instances (batch)
*TaskInstanceAPI* | [**PatchMappedTaskInstance**](docs/TaskInstanceAPI.md#patchmappedtaskinstance) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index} | Updates the state of a mapped task instance
*TaskInstanceAPI* | [**PatchTaskInstance**](docs/TaskInstanceAPI.md#patchtaskinstance) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id} | Updates the state of a task instance
*TaskInstanceAPI* | [**SetMappedTaskInstanceNote**](docs/TaskInstanceAPI.md#setmappedtaskinstancenote) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/{map_index}/setNote | Update the TaskInstance note.
*TaskInstanceAPI* | [**SetTaskInstanceNote**](docs/TaskInstanceAPI.md#settaskinstancenote) | **Patch** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/setNote | Update the TaskInstance note.
*UserAPI* | [**DeleteUser**](docs/UserAPI.md#deleteuser) | **Delete** /users/{username} | Delete a user
*UserAPI* | [**GetUser**](docs/UserAPI.md#getuser) | **Get** /users/{username} | Get a user
*UserAPI* | [**GetUsers**](docs/UserAPI.md#getusers) | **Get** /users | List users
*UserAPI* | [**PatchUser**](docs/UserAPI.md#patchuser) | **Patch** /users/{username} | Update a user
*UserAPI* | [**PostUser**](docs/UserAPI.md#postuser) | **Post** /users | Create a user
*VariableAPI* | [**DeleteVariable**](docs/VariableAPI.md#deletevariable) | **Delete** /variables/{variable_key} | Delete a variable
*VariableAPI* | [**GetVariable**](docs/VariableAPI.md#getvariable) | **Get** /variables/{variable_key} | Get a variable
*VariableAPI* | [**GetVariables**](docs/VariableAPI.md#getvariables) | **Get** /variables | List variables
*VariableAPI* | [**PatchVariable**](docs/VariableAPI.md#patchvariable) | **Patch** /variables/{variable_key} | Update a variable
*VariableAPI* | [**PostVariables**](docs/VariableAPI.md#postvariables) | **Post** /variables | Create a variable
*XComAPI* | [**GetXcomEntries**](docs/XComAPI.md#getxcomentries) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries | List XCom entries
*XComAPI* | [**GetXcomEntry**](docs/XComAPI.md#getxcomentry) | **Get** /dags/{dag_id}/dagRuns/{dag_run_id}/taskInstances/{task_id}/xcomEntries/{xcom_key} | Get an XCom entry


## Documentation For Models

 - [Action](docs/Action.md)
 - [ActionCollection](docs/ActionCollection.md)
 - [ActionResource](docs/ActionResource.md)
 - [BasicDAGRun](docs/BasicDAGRun.md)
 - [ClassReference](docs/ClassReference.md)
 - [ClearDagRun](docs/ClearDagRun.md)
 - [ClearDagRun200Response](docs/ClearDagRun200Response.md)
 - [ClearTaskInstances](docs/ClearTaskInstances.md)
 - [CollectionInfo](docs/CollectionInfo.md)
 - [Config](docs/Config.md)
 - [ConfigOption](docs/ConfigOption.md)
 - [ConfigSection](docs/ConfigSection.md)
 - [Connection](docs/Connection.md)
 - [ConnectionCollection](docs/ConnectionCollection.md)
 - [ConnectionCollectionItem](docs/ConnectionCollectionItem.md)
 - [ConnectionTest](docs/ConnectionTest.md)
 - [CreateDatasetEvent](docs/CreateDatasetEvent.md)
 - [CronExpression](docs/CronExpression.md)
 - [DAG](docs/DAG.md)
 - [DAGCollection](docs/DAGCollection.md)
 - [DAGDetail](docs/DAGDetail.md)
 - [DAGRun](docs/DAGRun.md)
 - [DAGRunCollection](docs/DAGRunCollection.md)
 - [DagProcessorStatus](docs/DagProcessorStatus.md)
 - [DagScheduleDatasetReference](docs/DagScheduleDatasetReference.md)
 - [DagState](docs/DagState.md)
 - [DagStatsCollectionItem](docs/DagStatsCollectionItem.md)
 - [DagStatsCollectionSchema](docs/DagStatsCollectionSchema.md)
 - [DagStatsStateCollectionItem](docs/DagStatsStateCollectionItem.md)
 - [DagWarning](docs/DagWarning.md)
 - [DagWarningCollection](docs/DagWarningCollection.md)
 - [Dataset](docs/Dataset.md)
 - [DatasetCollection](docs/DatasetCollection.md)
 - [DatasetEvent](docs/DatasetEvent.md)
 - [DatasetEventCollection](docs/DatasetEventCollection.md)
 - [Error](docs/Error.md)
 - [EventLog](docs/EventLog.md)
 - [EventLogCollection](docs/EventLogCollection.md)
 - [ExtraLink](docs/ExtraLink.md)
 - [ExtraLinkCollection](docs/ExtraLinkCollection.md)
 - [GetDagSource200Response](docs/GetDagSource200Response.md)
 - [GetLog200Response](docs/GetLog200Response.md)
 - [GetProviders200Response](docs/GetProviders200Response.md)
 - [HealthInfo](docs/HealthInfo.md)
 - [HealthStatus](docs/HealthStatus.md)
 - [ImportError](docs/ImportError.md)
 - [ImportErrorCollection](docs/ImportErrorCollection.md)
 - [Job](docs/Job.md)
 - [ListDagRunsForm](docs/ListDagRunsForm.md)
 - [ListTaskInstanceForm](docs/ListTaskInstanceForm.md)
 - [MetadatabaseStatus](docs/MetadatabaseStatus.md)
 - [PluginCollection](docs/PluginCollection.md)
 - [PluginCollectionItem](docs/PluginCollectionItem.md)
 - [Pool](docs/Pool.md)
 - [PoolCollection](docs/PoolCollection.md)
 - [Provider](docs/Provider.md)
 - [ProviderCollection](docs/ProviderCollection.md)
 - [QueuedEvent](docs/QueuedEvent.md)
 - [QueuedEventCollection](docs/QueuedEventCollection.md)
 - [RelativeDelta](docs/RelativeDelta.md)
 - [Resource](docs/Resource.md)
 - [Role](docs/Role.md)
 - [RoleCollection](docs/RoleCollection.md)
 - [SLAMiss](docs/SLAMiss.md)
 - [ScheduleInterval](docs/ScheduleInterval.md)
 - [SchedulerStatus](docs/SchedulerStatus.md)
 - [SetDagRunNote](docs/SetDagRunNote.md)
 - [SetTaskInstanceNote](docs/SetTaskInstanceNote.md)
 - [Tag](docs/Tag.md)
 - [Task](docs/Task.md)
 - [TaskCollection](docs/TaskCollection.md)
 - [TaskExtraLinksInner](docs/TaskExtraLinksInner.md)
 - [TaskFailedDependency](docs/TaskFailedDependency.md)
 - [TaskInstance](docs/TaskInstance.md)
 - [TaskInstanceCollection](docs/TaskInstanceCollection.md)
 - [TaskInstanceDependencyCollection](docs/TaskInstanceDependencyCollection.md)
 - [TaskInstanceHistory](docs/TaskInstanceHistory.md)
 - [TaskInstanceHistoryCollection](docs/TaskInstanceHistoryCollection.md)
 - [TaskInstanceReference](docs/TaskInstanceReference.md)
 - [TaskInstanceReferenceCollection](docs/TaskInstanceReferenceCollection.md)
 - [TaskOutletDatasetReference](docs/TaskOutletDatasetReference.md)
 - [TaskState](docs/TaskState.md)
 - [TimeDelta](docs/TimeDelta.md)
 - [Trigger](docs/Trigger.md)
 - [TriggerRule](docs/TriggerRule.md)
 - [TriggererStatus](docs/TriggererStatus.md)
 - [UpdateDagRunState](docs/UpdateDagRunState.md)
 - [UpdateTaskInstance](docs/UpdateTaskInstance.md)
 - [UpdateTaskInstancesState](docs/UpdateTaskInstancesState.md)
 - [UpdateTaskState](docs/UpdateTaskState.md)
 - [User](docs/User.md)
 - [UserCollection](docs/UserCollection.md)
 - [UserCollectionItem](docs/UserCollectionItem.md)
 - [UserCollectionItemRolesInner](docs/UserCollectionItemRolesInner.md)
 - [Variable](docs/Variable.md)
 - [VariableCollection](docs/VariableCollection.md)
 - [VariableCollectionItem](docs/VariableCollectionItem.md)
 - [VersionInfo](docs/VersionInfo.md)
 - [WeightRule](docs/WeightRule.md)
 - [XCom](docs/XCom.md)
 - [XComCollection](docs/XComCollection.md)
 - [XComCollectionItem](docs/XComCollectionItem.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Basic

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), airflow.ContextBasicAuth, airflow.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

### GoogleOpenId

### Kerberos


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

dev@airflow.apache.org

