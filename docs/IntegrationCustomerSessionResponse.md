# IntegrationCustomerSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerSession** | Pointer to [**CustomerSessionV2**](CustomerSessionV2.md) |  | [optional] 
**Effects** | Pointer to [**[]Effect**](Effect.md) | The returned effects.  **Note:** This endpoint returns only the effects that are valid after any rollback effects and their corresponding non-rollback effects are removed.  | [optional] 

## Methods

### GetCustomerSession

`func (o *IntegrationCustomerSessionResponse) GetCustomerSession() CustomerSessionV2`

GetCustomerSession returns the CustomerSession field if non-nil, zero value otherwise.

### GetCustomerSessionOk

`func (o *IntegrationCustomerSessionResponse) GetCustomerSessionOk() (CustomerSessionV2, bool)`

GetCustomerSessionOk returns a tuple with the CustomerSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCustomerSession

`func (o *IntegrationCustomerSessionResponse) HasCustomerSession() bool`

HasCustomerSession returns a boolean if a field has been set.

### SetCustomerSession

`func (o *IntegrationCustomerSessionResponse) SetCustomerSession(v CustomerSessionV2)`

SetCustomerSession gets a reference to the given CustomerSessionV2 and assigns it to the CustomerSession field.

### GetEffects

`func (o *IntegrationCustomerSessionResponse) GetEffects() []Effect`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *IntegrationCustomerSessionResponse) GetEffectsOk() ([]Effect, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEffects

`func (o *IntegrationCustomerSessionResponse) HasEffects() bool`

HasEffects returns a boolean if a field has been set.

### SetEffects

`func (o *IntegrationCustomerSessionResponse) SetEffects(v []Effect)`

SetEffects gets a reference to the given []Effect and assigns it to the Effects field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


