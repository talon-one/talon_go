# IntegrationHubFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationID** | Pointer to **int64** | ID of the application the flow is registered for. | [optional] 
**LoyaltyProgramID** | Pointer to **int64** | ID of the loyalty program the flow is registered for. | [optional] 
**EventType** | Pointer to [**IntegrationHubEventType**](IntegrationHubEventType.md) |  | 
**IntegrationHubFlowUrl** | Pointer to **string** | The URL of the integration hub flow that we want to trigger for the event. | 

## Methods

### NewIntegrationHubFlow

`func NewIntegrationHubFlow(eventType IntegrationHubEventType, integrationHubFlowUrl string, ) *IntegrationHubFlow`

NewIntegrationHubFlow instantiates a new IntegrationHubFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubFlowWithDefaults

`func NewIntegrationHubFlowWithDefaults() *IntegrationHubFlow`

NewIntegrationHubFlowWithDefaults instantiates a new IntegrationHubFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationID

`func (o *IntegrationHubFlow) GetApplicationID() int64`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *IntegrationHubFlow) GetApplicationIDOk() (*int64, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *IntegrationHubFlow) SetApplicationID(v int64)`

SetApplicationID sets ApplicationID field to given value.

### HasApplicationID

`func (o *IntegrationHubFlow) HasApplicationID() bool`

HasApplicationID returns a boolean if a field has been set.

### GetLoyaltyProgramID

`func (o *IntegrationHubFlow) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *IntegrationHubFlow) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *IntegrationHubFlow) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.

### HasLoyaltyProgramID

`func (o *IntegrationHubFlow) HasLoyaltyProgramID() bool`

HasLoyaltyProgramID returns a boolean if a field has been set.

### GetEventType

`func (o *IntegrationHubFlow) GetEventType() IntegrationHubEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IntegrationHubFlow) GetEventTypeOk() (*IntegrationHubEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IntegrationHubFlow) SetEventType(v IntegrationHubEventType)`

SetEventType sets EventType field to given value.


### GetIntegrationHubFlowUrl

`func (o *IntegrationHubFlow) GetIntegrationHubFlowUrl() string`

GetIntegrationHubFlowUrl returns the IntegrationHubFlowUrl field if non-nil, zero value otherwise.

### GetIntegrationHubFlowUrlOk

`func (o *IntegrationHubFlow) GetIntegrationHubFlowUrlOk() (*string, bool)`

GetIntegrationHubFlowUrlOk returns a tuple with the IntegrationHubFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationHubFlowUrl

`func (o *IntegrationHubFlow) SetIntegrationHubFlowUrl(v string)`

SetIntegrationHubFlowUrl sets IntegrationHubFlowUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


