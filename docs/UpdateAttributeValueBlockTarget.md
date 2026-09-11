# UpdateAttributeValueBlockTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Identifies the target scope of the attribute update. | 
**Name** | Pointer to **string** | Identifies the name of the target when its type is set to &#x60;selector&#x60; or &#x60;globalFilter&#x60;. | [optional] 

## Methods

### NewUpdateAttributeValueBlockTarget

`func NewUpdateAttributeValueBlockTarget(type_ string, ) *UpdateAttributeValueBlockTarget`

NewUpdateAttributeValueBlockTarget instantiates a new UpdateAttributeValueBlockTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAttributeValueBlockTargetWithDefaults

`func NewUpdateAttributeValueBlockTargetWithDefaults() *UpdateAttributeValueBlockTarget`

NewUpdateAttributeValueBlockTargetWithDefaults instantiates a new UpdateAttributeValueBlockTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateAttributeValueBlockTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAttributeValueBlockTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAttributeValueBlockTarget) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *UpdateAttributeValueBlockTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAttributeValueBlockTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAttributeValueBlockTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAttributeValueBlockTarget) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


