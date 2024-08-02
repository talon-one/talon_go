# Picklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CreatedBy** | Pointer to **int32** | ID of the user who created this effect. | 
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Imported** | Pointer to **bool** | Imported flag shows that a picklist is imported by a CSV file or not | [optional] 
**ModifiedBy** | Pointer to **int32** | ID of the user who last updated this effect if available. | [optional] 
**Type** | Pointer to **string** | The type of allowed values in the picklist. If the type &#x60;time&#x60; is chosen, it must be an RFC3339 timestamp string. | 
**Values** | Pointer to **[]string** | The list of allowed values provided by this picklist. | 

## Methods

### GetAccountId

`func (o *Picklist) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Picklist) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Picklist) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Picklist) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetCreated

`func (o *Picklist) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Picklist) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Picklist) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Picklist) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCreatedBy

`func (o *Picklist) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Picklist) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *Picklist) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *Picklist) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetId

`func (o *Picklist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Picklist) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Picklist) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Picklist) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetImported

`func (o *Picklist) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *Picklist) GetImportedOk() (bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImported

`func (o *Picklist) HasImported() bool`

HasImported returns a boolean if a field has been set.

### SetImported

`func (o *Picklist) SetImported(v bool)`

SetImported gets a reference to the given bool and assigns it to the Imported field.

### GetModifiedBy

`func (o *Picklist) GetModifiedBy() int32`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Picklist) GetModifiedByOk() (int32, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModifiedBy

`func (o *Picklist) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### SetModifiedBy

`func (o *Picklist) SetModifiedBy(v int32)`

SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.

### GetType

`func (o *Picklist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Picklist) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *Picklist) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *Picklist) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetValues

`func (o *Picklist) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Picklist) GetValuesOk() ([]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *Picklist) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *Picklist) SetValues(v []string)`

SetValues gets a reference to the given []string and assigns it to the Values field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


