# IntegrationHubEventRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID of the event record. | 
**FlowId** | Pointer to **int64** | ID of the integration hub flow. | 
**IntegrationName** | Pointer to **string** | Name of the integration. | [optional] 
**InstanceName** | Pointer to **string** | Name of the integration instance. | [optional] 
**EventType** | Pointer to [**IntegrationHubEventType**](IntegrationHubEventType.md) |  | 
**PublishedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the event was published. | 
**ProcessedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the event was processed. | [optional] 
**DeliveredAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the event was delivered. | [optional] 
**ScheduledTo** | Pointer to [**time.Time**](time.Time.md) | Timestamp after which the event is scheduled to be processed. | 
**Retry** | Pointer to **int64** | Number of delivery retries attempted. | 
**Payload** | Pointer to **string** | The event payload as a formatted JSON string. | 

## Methods

### NewIntegrationHubEventRecord

`func NewIntegrationHubEventRecord(id int64, flowId int64, eventType IntegrationHubEventType, publishedAt time.Time, scheduledTo time.Time, retry int64, payload string, ) *IntegrationHubEventRecord`

NewIntegrationHubEventRecord instantiates a new IntegrationHubEventRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubEventRecordWithDefaults

`func NewIntegrationHubEventRecordWithDefaults() *IntegrationHubEventRecord`

NewIntegrationHubEventRecordWithDefaults instantiates a new IntegrationHubEventRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationHubEventRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationHubEventRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationHubEventRecord) SetId(v int64)`

SetId sets Id field to given value.


### GetFlowId

`func (o *IntegrationHubEventRecord) GetFlowId() int64`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *IntegrationHubEventRecord) GetFlowIdOk() (*int64, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *IntegrationHubEventRecord) SetFlowId(v int64)`

SetFlowId sets FlowId field to given value.


### GetIntegrationName

`func (o *IntegrationHubEventRecord) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *IntegrationHubEventRecord) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *IntegrationHubEventRecord) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.

### HasIntegrationName

`func (o *IntegrationHubEventRecord) HasIntegrationName() bool`

HasIntegrationName returns a boolean if a field has been set.

### GetInstanceName

`func (o *IntegrationHubEventRecord) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *IntegrationHubEventRecord) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *IntegrationHubEventRecord) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *IntegrationHubEventRecord) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetEventType

`func (o *IntegrationHubEventRecord) GetEventType() IntegrationHubEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IntegrationHubEventRecord) GetEventTypeOk() (*IntegrationHubEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IntegrationHubEventRecord) SetEventType(v IntegrationHubEventType)`

SetEventType sets EventType field to given value.


### GetPublishedAt

`func (o *IntegrationHubEventRecord) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *IntegrationHubEventRecord) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *IntegrationHubEventRecord) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetProcessedAt

`func (o *IntegrationHubEventRecord) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *IntegrationHubEventRecord) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *IntegrationHubEventRecord) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *IntegrationHubEventRecord) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *IntegrationHubEventRecord) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *IntegrationHubEventRecord) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *IntegrationHubEventRecord) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *IntegrationHubEventRecord) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### GetScheduledTo

`func (o *IntegrationHubEventRecord) GetScheduledTo() time.Time`

GetScheduledTo returns the ScheduledTo field if non-nil, zero value otherwise.

### GetScheduledToOk

`func (o *IntegrationHubEventRecord) GetScheduledToOk() (*time.Time, bool)`

GetScheduledToOk returns a tuple with the ScheduledTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTo

`func (o *IntegrationHubEventRecord) SetScheduledTo(v time.Time)`

SetScheduledTo sets ScheduledTo field to given value.


### GetRetry

`func (o *IntegrationHubEventRecord) GetRetry() int64`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *IntegrationHubEventRecord) GetRetryOk() (*int64, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *IntegrationHubEventRecord) SetRetry(v int64)`

SetRetry sets Retry field to given value.


### GetPayload

`func (o *IntegrationHubEventRecord) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *IntegrationHubEventRecord) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *IntegrationHubEventRecord) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


