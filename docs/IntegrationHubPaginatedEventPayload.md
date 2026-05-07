# IntegrationHubPaginatedEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | 
**BatchedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the batch was created. | [optional] 
**EventType** | Pointer to **string** |  | 
**Data** | Pointer to **[]map[string]interface{}** |  | 

## Methods

### NewIntegrationHubPaginatedEventPayload

`func NewIntegrationHubPaginatedEventPayload(totalResultSize int64, eventType string, data []map[string]interface{}, ) *IntegrationHubPaginatedEventPayload`

NewIntegrationHubPaginatedEventPayload instantiates a new IntegrationHubPaginatedEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubPaginatedEventPayloadWithDefaults

`func NewIntegrationHubPaginatedEventPayloadWithDefaults() *IntegrationHubPaginatedEventPayload`

NewIntegrationHubPaginatedEventPayloadWithDefaults instantiates a new IntegrationHubPaginatedEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *IntegrationHubPaginatedEventPayload) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *IntegrationHubPaginatedEventPayload) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *IntegrationHubPaginatedEventPayload) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetBatchedAt

`func (o *IntegrationHubPaginatedEventPayload) GetBatchedAt() time.Time`

GetBatchedAt returns the BatchedAt field if non-nil, zero value otherwise.

### GetBatchedAtOk

`func (o *IntegrationHubPaginatedEventPayload) GetBatchedAtOk() (*time.Time, bool)`

GetBatchedAtOk returns a tuple with the BatchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchedAt

`func (o *IntegrationHubPaginatedEventPayload) SetBatchedAt(v time.Time)`

SetBatchedAt sets BatchedAt field to given value.

### HasBatchedAt

`func (o *IntegrationHubPaginatedEventPayload) HasBatchedAt() bool`

HasBatchedAt returns a boolean if a field has been set.

### GetEventType

`func (o *IntegrationHubPaginatedEventPayload) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IntegrationHubPaginatedEventPayload) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IntegrationHubPaginatedEventPayload) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetData

`func (o *IntegrationHubPaginatedEventPayload) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IntegrationHubPaginatedEventPayload) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IntegrationHubPaginatedEventPayload) SetData(v []map[string]interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


