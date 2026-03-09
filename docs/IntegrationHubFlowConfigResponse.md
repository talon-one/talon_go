# IntegrationHubFlowConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerCount** | Pointer to **int64** | Number of IntegrationHub workers to run in parallel for this flow (maximum 500). | [optional] 
**MaxEventsPerMessage** | Pointer to **int64** | Maximum number of events to send in a single message to IntegrationHub. | [optional] 
**MaxRetries** | Pointer to **int64** | Maximum number of retries for a IntegrationHub event before it is ignored. | [optional] 

## Methods

### NewIntegrationHubFlowConfigResponse

`func NewIntegrationHubFlowConfigResponse() *IntegrationHubFlowConfigResponse`

NewIntegrationHubFlowConfigResponse instantiates a new IntegrationHubFlowConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubFlowConfigResponseWithDefaults

`func NewIntegrationHubFlowConfigResponseWithDefaults() *IntegrationHubFlowConfigResponse`

NewIntegrationHubFlowConfigResponseWithDefaults instantiates a new IntegrationHubFlowConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerCount

`func (o *IntegrationHubFlowConfigResponse) GetWorkerCount() int64`

GetWorkerCount returns the WorkerCount field if non-nil, zero value otherwise.

### GetWorkerCountOk

`func (o *IntegrationHubFlowConfigResponse) GetWorkerCountOk() (*int64, bool)`

GetWorkerCountOk returns a tuple with the WorkerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerCount

`func (o *IntegrationHubFlowConfigResponse) SetWorkerCount(v int64)`

SetWorkerCount sets WorkerCount field to given value.

### HasWorkerCount

`func (o *IntegrationHubFlowConfigResponse) HasWorkerCount() bool`

HasWorkerCount returns a boolean if a field has been set.

### GetMaxEventsPerMessage

`func (o *IntegrationHubFlowConfigResponse) GetMaxEventsPerMessage() int64`

GetMaxEventsPerMessage returns the MaxEventsPerMessage field if non-nil, zero value otherwise.

### GetMaxEventsPerMessageOk

`func (o *IntegrationHubFlowConfigResponse) GetMaxEventsPerMessageOk() (*int64, bool)`

GetMaxEventsPerMessageOk returns a tuple with the MaxEventsPerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventsPerMessage

`func (o *IntegrationHubFlowConfigResponse) SetMaxEventsPerMessage(v int64)`

SetMaxEventsPerMessage sets MaxEventsPerMessage field to given value.

### HasMaxEventsPerMessage

`func (o *IntegrationHubFlowConfigResponse) HasMaxEventsPerMessage() bool`

HasMaxEventsPerMessage returns a boolean if a field has been set.

### GetMaxRetries

`func (o *IntegrationHubFlowConfigResponse) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *IntegrationHubFlowConfigResponse) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *IntegrationHubFlowConfigResponse) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *IntegrationHubFlowConfigResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


