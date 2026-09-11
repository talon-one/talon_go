# UpdateAttributeValueBlockAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of the attribute. Reverts to &#x60;0&#x60; when the attribute is deleted or does not exist. | 
**Entity** | Pointer to **string** | The entity type that owns the attribute. Reverts to an empty string when the attribute is deleted or does not exist. | 
**Name** | Pointer to **string** | The attribute name as used in API requests. | 
**Title** | Pointer to **string** | The human-readable name of the attribute. | 
**Type** | Pointer to **string** | The data type of the attribute. | 

## Methods

### NewUpdateAttributeValueBlockAttribute

`func NewUpdateAttributeValueBlockAttribute(id int64, entity string, name string, title string, type_ string, ) *UpdateAttributeValueBlockAttribute`

NewUpdateAttributeValueBlockAttribute instantiates a new UpdateAttributeValueBlockAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAttributeValueBlockAttributeWithDefaults

`func NewUpdateAttributeValueBlockAttributeWithDefaults() *UpdateAttributeValueBlockAttribute`

NewUpdateAttributeValueBlockAttributeWithDefaults instantiates a new UpdateAttributeValueBlockAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAttributeValueBlockAttribute) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAttributeValueBlockAttribute) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAttributeValueBlockAttribute) SetId(v int64)`

SetId sets Id field to given value.


### GetEntity

`func (o *UpdateAttributeValueBlockAttribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *UpdateAttributeValueBlockAttribute) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *UpdateAttributeValueBlockAttribute) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetName

`func (o *UpdateAttributeValueBlockAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAttributeValueBlockAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAttributeValueBlockAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *UpdateAttributeValueBlockAttribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateAttributeValueBlockAttribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateAttributeValueBlockAttribute) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *UpdateAttributeValueBlockAttribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAttributeValueBlockAttribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAttributeValueBlockAttribute) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


