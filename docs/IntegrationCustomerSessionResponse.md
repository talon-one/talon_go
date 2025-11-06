# IntegrationCustomerSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerSession** | Pointer to [**CustomerSessionV2**](CustomerSessionV2.md) |  | [optional] 
**Effects** | Pointer to [**[]Effect**](Effect.md) | The returned effects.  **Note:** This endpoint returns only the effects that are valid after any rollback effects and their corresponding non-rollback effects are removed.  | [optional] 

## Methods

### NewIntegrationCustomerSessionResponse

`func NewIntegrationCustomerSessionResponse() *IntegrationCustomerSessionResponse`

NewIntegrationCustomerSessionResponse instantiates a new IntegrationCustomerSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationCustomerSessionResponseWithDefaults

`func NewIntegrationCustomerSessionResponseWithDefaults() *IntegrationCustomerSessionResponse`

NewIntegrationCustomerSessionResponseWithDefaults instantiates a new IntegrationCustomerSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerSession

`func (o *IntegrationCustomerSessionResponse) GetCustomerSession() CustomerSessionV2`

GetCustomerSession returns the CustomerSession field if non-nil, zero value otherwise.

### GetCustomerSessionOk

`func (o *IntegrationCustomerSessionResponse) GetCustomerSessionOk() (*CustomerSessionV2, bool)`

GetCustomerSessionOk returns a tuple with the CustomerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSession

`func (o *IntegrationCustomerSessionResponse) SetCustomerSession(v CustomerSessionV2)`

SetCustomerSession sets CustomerSession field to given value.

### HasCustomerSession

`func (o *IntegrationCustomerSessionResponse) HasCustomerSession() bool`

HasCustomerSession returns a boolean if a field has been set.

### GetEffects

`func (o *IntegrationCustomerSessionResponse) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *IntegrationCustomerSessionResponse) GetEffectsOk() (*[]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *IntegrationCustomerSessionResponse) SetEffects(v []Effect)`

SetEffects sets Effects field to given value.

### HasEffects

`func (o *IntegrationCustomerSessionResponse) HasEffects() bool`

HasEffects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


