# AttributeBlockReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of the attribute. Reverts to &#x60;0&#x60; when the attribute is deleted or does not exist. | 
**Entity** | Pointer to **string** | The entity type that owns the attribute. Reverts to an empty string when the attribute is deleted or does not exist. | 
**Name** | Pointer to **string** | The attribute name as used in API requests. | 
**Title** | Pointer to **string** | The human-readable name of the attribute. | [optional] [readonly] 
**Type** | Pointer to **string** | The data type of the attribute. | [optional] [readonly] 

## Methods

### NewAttributeBlockReference

`func NewAttributeBlockReference(id int64, entity string, name string, ) *AttributeBlockReference`

NewAttributeBlockReference instantiates a new AttributeBlockReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeBlockReferenceWithDefaults

`func NewAttributeBlockReferenceWithDefaults() *AttributeBlockReference`

NewAttributeBlockReferenceWithDefaults instantiates a new AttributeBlockReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttributeBlockReference) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeBlockReference) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeBlockReference) SetId(v int64)`

SetId sets Id field to given value.


### GetEntity

`func (o *AttributeBlockReference) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *AttributeBlockReference) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *AttributeBlockReference) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetName

`func (o *AttributeBlockReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttributeBlockReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttributeBlockReference) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AttributeBlockReference) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AttributeBlockReference) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AttributeBlockReference) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AttributeBlockReference) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *AttributeBlockReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttributeBlockReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttributeBlockReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AttributeBlockReference) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


