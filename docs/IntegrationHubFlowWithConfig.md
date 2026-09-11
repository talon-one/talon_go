# IntegrationHubFlowWithConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationID** | Pointer to **int64** | ID of the application the flow is registered for. | [optional] 
**LoyaltyProgramID** | Pointer to **int64** | ID of the loyalty program the flow is registered for. | [optional] 
**EventType** | Pointer to [**IntegrationHubEventType**](IntegrationHubEventType.md) |  | 
**IntegrationHubFlowUrl** | Pointer to **string** | The URL of the integration hub flow that we want to trigger for the event. | 
**Config** | Pointer to [**IntegrationHubFlowConfig**](IntegrationHubFlowConfig.md) |  | 

## Methods

### NewIntegrationHubFlowWithConfig

`func NewIntegrationHubFlowWithConfig(eventType IntegrationHubEventType, integrationHubFlowUrl string, config IntegrationHubFlowConfig, ) *IntegrationHubFlowWithConfig`

NewIntegrationHubFlowWithConfig instantiates a new IntegrationHubFlowWithConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubFlowWithConfigWithDefaults

`func NewIntegrationHubFlowWithConfigWithDefaults() *IntegrationHubFlowWithConfig`

NewIntegrationHubFlowWithConfigWithDefaults instantiates a new IntegrationHubFlowWithConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationID

`func (o *IntegrationHubFlowWithConfig) GetApplicationID() int64`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *IntegrationHubFlowWithConfig) GetApplicationIDOk() (*int64, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *IntegrationHubFlowWithConfig) SetApplicationID(v int64)`

SetApplicationID sets ApplicationID field to given value.

### HasApplicationID

`func (o *IntegrationHubFlowWithConfig) HasApplicationID() bool`

HasApplicationID returns a boolean if a field has been set.

### GetLoyaltyProgramID

`func (o *IntegrationHubFlowWithConfig) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *IntegrationHubFlowWithConfig) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *IntegrationHubFlowWithConfig) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.

### HasLoyaltyProgramID

`func (o *IntegrationHubFlowWithConfig) HasLoyaltyProgramID() bool`

HasLoyaltyProgramID returns a boolean if a field has been set.

### GetEventType

`func (o *IntegrationHubFlowWithConfig) GetEventType() IntegrationHubEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IntegrationHubFlowWithConfig) GetEventTypeOk() (*IntegrationHubEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IntegrationHubFlowWithConfig) SetEventType(v IntegrationHubEventType)`

SetEventType sets EventType field to given value.


### GetIntegrationHubFlowUrl

`func (o *IntegrationHubFlowWithConfig) GetIntegrationHubFlowUrl() string`

GetIntegrationHubFlowUrl returns the IntegrationHubFlowUrl field if non-nil, zero value otherwise.

### GetIntegrationHubFlowUrlOk

`func (o *IntegrationHubFlowWithConfig) GetIntegrationHubFlowUrlOk() (*string, bool)`

GetIntegrationHubFlowUrlOk returns a tuple with the IntegrationHubFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationHubFlowUrl

`func (o *IntegrationHubFlowWithConfig) SetIntegrationHubFlowUrl(v string)`

SetIntegrationHubFlowUrl sets IntegrationHubFlowUrl field to given value.


### GetConfig

`func (o *IntegrationHubFlowWithConfig) GetConfig() IntegrationHubFlowConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationHubFlowWithConfig) GetConfigOk() (*IntegrationHubFlowConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationHubFlowWithConfig) SetConfig(v IntegrationHubFlowConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


