# IntegrationHubFlowConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | 
**WorkerCount** | Pointer to **int64** | Number of IntegrationHub workers to run in parallel for this flow (maximum 500). | [optional] [default to 10]
**MaxEventsPerMessage** | Pointer to **int64** | Maximum number of events to send in a single message to IntegrationHub. | [optional] [default to 1000]
**MaxRetries** | Pointer to **int64** | Maximum number of retries for a IntegrationHub event before it is ignored. | [optional] [default to 10]

## Methods

### NewIntegrationHubFlowConfig

`func NewIntegrationHubFlowConfig(apiKey string, ) *IntegrationHubFlowConfig`

NewIntegrationHubFlowConfig instantiates a new IntegrationHubFlowConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubFlowConfigWithDefaults

`func NewIntegrationHubFlowConfigWithDefaults() *IntegrationHubFlowConfig`

NewIntegrationHubFlowConfigWithDefaults instantiates a new IntegrationHubFlowConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *IntegrationHubFlowConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *IntegrationHubFlowConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *IntegrationHubFlowConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetWorkerCount

`func (o *IntegrationHubFlowConfig) GetWorkerCount() int64`

GetWorkerCount returns the WorkerCount field if non-nil, zero value otherwise.

### GetWorkerCountOk

`func (o *IntegrationHubFlowConfig) GetWorkerCountOk() (*int64, bool)`

GetWorkerCountOk returns a tuple with the WorkerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerCount

`func (o *IntegrationHubFlowConfig) SetWorkerCount(v int64)`

SetWorkerCount sets WorkerCount field to given value.

### HasWorkerCount

`func (o *IntegrationHubFlowConfig) HasWorkerCount() bool`

HasWorkerCount returns a boolean if a field has been set.

### GetMaxEventsPerMessage

`func (o *IntegrationHubFlowConfig) GetMaxEventsPerMessage() int64`

GetMaxEventsPerMessage returns the MaxEventsPerMessage field if non-nil, zero value otherwise.

### GetMaxEventsPerMessageOk

`func (o *IntegrationHubFlowConfig) GetMaxEventsPerMessageOk() (*int64, bool)`

GetMaxEventsPerMessageOk returns a tuple with the MaxEventsPerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventsPerMessage

`func (o *IntegrationHubFlowConfig) SetMaxEventsPerMessage(v int64)`

SetMaxEventsPerMessage sets MaxEventsPerMessage field to given value.

### HasMaxEventsPerMessage

`func (o *IntegrationHubFlowConfig) HasMaxEventsPerMessage() bool`

HasMaxEventsPerMessage returns a boolean if a field has been set.

### GetMaxRetries

`func (o *IntegrationHubFlowConfig) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *IntegrationHubFlowConfig) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *IntegrationHubFlowConfig) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *IntegrationHubFlowConfig) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


