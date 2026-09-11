# IntegrationHubFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID of the integration hub flow. | 
**IntegrationName** | Pointer to **string** | Name of the integration. | [optional] 
**InstanceName** | Pointer to **string** | Name of the integration instance. | [optional] 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the flow was created. | 
**DisabledUntil** | Pointer to [**NullableTime**](time.Time.md) | Timestamp until which the flow is disabled. Null when the flow is active. | [optional] 
**ApplicationId** | Pointer to **int64** | ID of the application the flow is registered for. | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | ID of the loyalty program the flow is registered for. | [optional] 
**EventType** | Pointer to **string** | The event type we want to register a flow for. | 
**Config** | Pointer to [**IntegrationHubFlowConfigResponse**](IntegrationHubFlowConfigResponse.md) |  | 

## Methods

### NewIntegrationHubFlowResponse

`func NewIntegrationHubFlowResponse(id int64, createdAt time.Time, eventType string, config IntegrationHubFlowConfigResponse, ) *IntegrationHubFlowResponse`

NewIntegrationHubFlowResponse instantiates a new IntegrationHubFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubFlowResponseWithDefaults

`func NewIntegrationHubFlowResponseWithDefaults() *IntegrationHubFlowResponse`

NewIntegrationHubFlowResponseWithDefaults instantiates a new IntegrationHubFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationHubFlowResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationHubFlowResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationHubFlowResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetIntegrationName

`func (o *IntegrationHubFlowResponse) GetIntegrationName() string`

GetIntegrationName returns the IntegrationName field if non-nil, zero value otherwise.

### GetIntegrationNameOk

`func (o *IntegrationHubFlowResponse) GetIntegrationNameOk() (*string, bool)`

GetIntegrationNameOk returns a tuple with the IntegrationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationName

`func (o *IntegrationHubFlowResponse) SetIntegrationName(v string)`

SetIntegrationName sets IntegrationName field to given value.

### HasIntegrationName

`func (o *IntegrationHubFlowResponse) HasIntegrationName() bool`

HasIntegrationName returns a boolean if a field has been set.

### GetInstanceName

`func (o *IntegrationHubFlowResponse) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *IntegrationHubFlowResponse) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *IntegrationHubFlowResponse) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *IntegrationHubFlowResponse) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IntegrationHubFlowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationHubFlowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationHubFlowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisabledUntil

`func (o *IntegrationHubFlowResponse) GetDisabledUntil() time.Time`

GetDisabledUntil returns the DisabledUntil field if non-nil, zero value otherwise.

### GetDisabledUntilOk

`func (o *IntegrationHubFlowResponse) GetDisabledUntilOk() (*time.Time, bool)`

GetDisabledUntilOk returns a tuple with the DisabledUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledUntil

`func (o *IntegrationHubFlowResponse) SetDisabledUntil(v time.Time)`

SetDisabledUntil sets DisabledUntil field to given value.

### HasDisabledUntil

`func (o *IntegrationHubFlowResponse) HasDisabledUntil() bool`

HasDisabledUntil returns a boolean if a field has been set.

### SetDisabledUntilNil

`func (o *IntegrationHubFlowResponse) SetDisabledUntilNil(b bool)`

 SetDisabledUntilNil sets the value for DisabledUntil to be an explicit nil

### UnsetDisabledUntil
`func (o *IntegrationHubFlowResponse) UnsetDisabledUntil()`

UnsetDisabledUntil ensures that no value is present for DisabledUntil, not even an explicit nil
### GetApplicationId

`func (o *IntegrationHubFlowResponse) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *IntegrationHubFlowResponse) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *IntegrationHubFlowResponse) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *IntegrationHubFlowResponse) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *IntegrationHubFlowResponse) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *IntegrationHubFlowResponse) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *IntegrationHubFlowResponse) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *IntegrationHubFlowResponse) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetEventType

`func (o *IntegrationHubFlowResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IntegrationHubFlowResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IntegrationHubFlowResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetConfig

`func (o *IntegrationHubFlowResponse) GetConfig() IntegrationHubFlowConfigResponse`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationHubFlowResponse) GetConfigOk() (*IntegrationHubFlowConfigResponse, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationHubFlowResponse) SetConfig(v IntegrationHubFlowConfigResponse)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


