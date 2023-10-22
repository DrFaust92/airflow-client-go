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

# DagProcessorStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NullableHealthStatus**](HealthStatus.md) |  | [optional] 
**LatestDagProcessorHeartbeat** | Pointer to **NullableString** | The time the dag processor last did a heartbeat. | [optional] [readonly] 

## Methods

### NewDagProcessorStatus

`func NewDagProcessorStatus() *DagProcessorStatus`

NewDagProcessorStatus instantiates a new DagProcessorStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDagProcessorStatusWithDefaults

`func NewDagProcessorStatusWithDefaults() *DagProcessorStatus`

NewDagProcessorStatusWithDefaults instantiates a new DagProcessorStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DagProcessorStatus) GetStatus() HealthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DagProcessorStatus) GetStatusOk() (*HealthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DagProcessorStatus) SetStatus(v HealthStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DagProcessorStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DagProcessorStatus) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DagProcessorStatus) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetLatestDagProcessorHeartbeat

`func (o *DagProcessorStatus) GetLatestDagProcessorHeartbeat() string`

GetLatestDagProcessorHeartbeat returns the LatestDagProcessorHeartbeat field if non-nil, zero value otherwise.

### GetLatestDagProcessorHeartbeatOk

`func (o *DagProcessorStatus) GetLatestDagProcessorHeartbeatOk() (*string, bool)`

GetLatestDagProcessorHeartbeatOk returns a tuple with the LatestDagProcessorHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDagProcessorHeartbeat

`func (o *DagProcessorStatus) SetLatestDagProcessorHeartbeat(v string)`

SetLatestDagProcessorHeartbeat sets LatestDagProcessorHeartbeat field to given value.

### HasLatestDagProcessorHeartbeat

`func (o *DagProcessorStatus) HasLatestDagProcessorHeartbeat() bool`

HasLatestDagProcessorHeartbeat returns a boolean if a field has been set.

### SetLatestDagProcessorHeartbeatNil

`func (o *DagProcessorStatus) SetLatestDagProcessorHeartbeatNil(b bool)`

 SetLatestDagProcessorHeartbeatNil sets the value for LatestDagProcessorHeartbeat to be an explicit nil

### UnsetLatestDagProcessorHeartbeat
`func (o *DagProcessorStatus) UnsetLatestDagProcessorHeartbeat()`

UnsetLatestDagProcessorHeartbeat ensures that no value is present for LatestDagProcessorHeartbeat, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


