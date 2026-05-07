# IntegrationHubFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | ID of the integration hub flow. | 
**ApplicationID** | Pointer to **int64** | ID of application the flow is registered for. | [optional] 
**EventType** | Pointer to **string** | The event type we want to register a flow for. | 
**IntegrationHubFlowUrl** | Pointer to **string** | The URL of the integration hub flow that we want to trigger for the event. | 
**Config** | Pointer to [**IntegrationHubFlowConfigResponse**](IntegrationHubFlowConfigResponse.md) |  | 

## Methods

### NewIntegrationHubFlowResponse

`func NewIntegrationHubFlowResponse(id int64, eventType string, integrationHubFlowUrl string, config IntegrationHubFlowConfigResponse, ) *IntegrationHubFlowResponse`

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


### GetApplicationID

`func (o *IntegrationHubFlowResponse) GetApplicationID() int64`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *IntegrationHubFlowResponse) GetApplicationIDOk() (*int64, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *IntegrationHubFlowResponse) SetApplicationID(v int64)`

SetApplicationID sets ApplicationID field to given value.

### HasApplicationID

`func (o *IntegrationHubFlowResponse) HasApplicationID() bool`

HasApplicationID returns a boolean if a field has been set.

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


### GetIntegrationHubFlowUrl

`func (o *IntegrationHubFlowResponse) GetIntegrationHubFlowUrl() string`

GetIntegrationHubFlowUrl returns the IntegrationHubFlowUrl field if non-nil, zero value otherwise.

### GetIntegrationHubFlowUrlOk

`func (o *IntegrationHubFlowResponse) GetIntegrationHubFlowUrlOk() (*string, bool)`

GetIntegrationHubFlowUrlOk returns a tuple with the IntegrationHubFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationHubFlowUrl

`func (o *IntegrationHubFlowResponse) SetIntegrationHubFlowUrl(v string)`

SetIntegrationHubFlowUrl sets IntegrationHubFlowUrl field to given value.


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


