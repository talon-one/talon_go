# ScimResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the resource. | [optional] 
**Name** | Pointer to **string** | Name of the resource. | [optional] 
**Description** | Pointer to **string** | Human-readable description of the resource. | [optional] 

## Methods

### NewScimResource

`func NewScimResource() *ScimResource`

NewScimResource instantiates a new ScimResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimResourceWithDefaults

`func NewScimResourceWithDefaults() *ScimResource`

NewScimResourceWithDefaults instantiates a new ScimResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScimResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScimResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScimResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScimResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScimResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScimResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScimResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScimResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ScimResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScimResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScimResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScimResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


