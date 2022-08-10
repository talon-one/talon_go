# MultipleAudiences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Audiences** | Pointer to [**[]MultipleAudiencesItem**](MultipleAudiencesItem.md) |  | 

## Methods

### GetAccountId

`func (o *MultipleAudiences) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MultipleAudiences) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *MultipleAudiences) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *MultipleAudiences) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetAudiences

`func (o *MultipleAudiences) GetAudiences() []MultipleAudiencesItem`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *MultipleAudiences) GetAudiencesOk() ([]MultipleAudiencesItem, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAudiences

`func (o *MultipleAudiences) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### SetAudiences

`func (o *MultipleAudiences) SetAudiences(v []MultipleAudiencesItem)`

SetAudiences gets a reference to the given []MultipleAudiencesItem and assigns it to the Audiences field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


