# OutgoingIntegrationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**AccountId** | Pointer to **int32** | The ID of the account to which this configuration belongs. | 
**TypeId** | Pointer to **int32** | The outgoing integration type ID. | 
**Policy** | Pointer to [**map[string]interface{}**](.md) | The outgoing integration policy specific to each integration type. | 

## Methods

### GetId

`func (o *OutgoingIntegrationConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutgoingIntegrationConfiguration) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *OutgoingIntegrationConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *OutgoingIntegrationConfiguration) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetAccountId

`func (o *OutgoingIntegrationConfiguration) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *OutgoingIntegrationConfiguration) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *OutgoingIntegrationConfiguration) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *OutgoingIntegrationConfiguration) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetTypeId

`func (o *OutgoingIntegrationConfiguration) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *OutgoingIntegrationConfiguration) GetTypeIdOk() (int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTypeId

`func (o *OutgoingIntegrationConfiguration) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeId

`func (o *OutgoingIntegrationConfiguration) SetTypeId(v int32)`

SetTypeId gets a reference to the given int32 and assigns it to the TypeId field.

### GetPolicy

`func (o *OutgoingIntegrationConfiguration) GetPolicy() map[string]interface{}`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *OutgoingIntegrationConfiguration) GetPolicyOk() (map[string]interface{}, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicy

`func (o *OutgoingIntegrationConfiguration) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicy

`func (o *OutgoingIntegrationConfiguration) SetPolicy(v map[string]interface{})`

SetPolicy gets a reference to the given map[string]interface{} and assigns it to the Policy field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


