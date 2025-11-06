# Picklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Type** | Pointer to **string** | The type of allowed values in the picklist. If the type &#x60;time&#x60; is chosen, it must be an RFC3339 timestamp string. | 
**Values** | Pointer to **[]string** | The list of allowed values provided by this picklist. | 
**ModifiedBy** | Pointer to **int64** | ID of the user who last updated this effect if available. | [optional] 
**CreatedBy** | Pointer to **int64** | ID of the user who created this effect. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | [optional] 
**Imported** | Pointer to **bool** | Imported flag shows that a picklist is imported by a CSV file or not | [optional] 

## Methods

### NewPicklist

`func NewPicklist(id int64, created time.Time, type_ string, values []string, createdBy int64, ) *Picklist`

NewPicklist instantiates a new Picklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPicklistWithDefaults

`func NewPicklistWithDefaults() *Picklist`

NewPicklistWithDefaults instantiates a new Picklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Picklist) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Picklist) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Picklist) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Picklist) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Picklist) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Picklist) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetType

`func (o *Picklist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Picklist) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Picklist) SetType(v string)`

SetType sets Type field to given value.


### GetValues

`func (o *Picklist) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Picklist) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Picklist) SetValues(v []string)`

SetValues sets Values field to given value.


### GetModifiedBy

`func (o *Picklist) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Picklist) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Picklist) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Picklist) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Picklist) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Picklist) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Picklist) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetAccountId

`func (o *Picklist) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Picklist) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Picklist) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Picklist) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetImported

`func (o *Picklist) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *Picklist) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *Picklist) SetImported(v bool)`

SetImported sets Imported field to given value.

### HasImported

`func (o *Picklist) HasImported() bool`

HasImported returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


